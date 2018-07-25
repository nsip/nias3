package webserver

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/nsip/nias3/xml2triples"
	"github.com/radovskyb/watcher"
	"github.com/twinj/uuid"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	//"sync"
	"time"
)

var xmlobject = regexp.MustCompile(`(?s:^\s*<([^> ]+))`)

func directoryWatcher() {
	var err error
	w := watcher.New()
	w.SetMaxEvents(1)
	// Only notify Create and Write events; ignore Move and Rename
	w.FilterOps(watcher.Create, watcher.Write)
	go func() {
		select {
		case event := <-w.Event:
			log.Printf("Adding XML document %s\n", event.Path)
			SendXmlToDataStore(event.Path)
		case err := <-w.Error:
			log.Fatalln(err)
		case <-w.Closed:
			return
		}
	}()

	if err = w.AddRecursive("./in"); err != nil {
		log.Fatalln(err)
	}
	go func() {
		if err := w.Start(time.Millisecond * 100); err != nil {
			log.Fatalln(err)
		}
	}()
}

func SendXmlToDataStore(filename string) error {
	fi, err := os.Lstat(filename)
	if err != nil {
		return err
	}
	if fi.Mode().IsDir() {
		return fmt.Errorf("%s is a directory", filename)
	}
	if !strings.HasSuffix(filename, ".xml") {
		return fmt.Errorf("%s is not an XML file", filename)
	}
	file, err := os.Open(filename)
	if err != nil {
		log.Printf("Cannot read in file %s\n", filename)
		return err
	}
	err = sendReaderToDataStore(file)
	if err != nil {
		return err
	}
	log.Printf("Read in file %s into filestore\n", filename)
	return nil
}

func SIFGetToDataStore(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = sendReaderToDataStore(resp.Body)
	if err != nil {
		return err
	}
	return nil
}

var firsttag = regexp.MustCompile(`^\s*<[^>]+>`)
var lasttag = regexp.MustCompile(`<[^>]+>\s*$`)

// url: e.g. http://hits.nsip.edu.au/SIF3InfraREST/hits/requests/SchoolInfos?navigationPage=1&navigationPageSize=5&access_token=ZmZhODMzNjEtMGExOC00NDk5LTgyNjMtYjMwNjI4MGRjZDRlOmYxYzA1NjNhOWIzZTQyMGJiMDdkYTJkOTBkYjQ3OWVm&authenticationMethod=Basic
func SIFGetManyToDataStore(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body = firsttag.ReplaceAll(
		lasttag.ReplaceAll(body, []byte("")), []byte(""))
	err = sendReaderToDataStore(bytes.NewReader(body))
	if err != nil {
		return err
	}
	return nil
}

func simultSendTriples(alltriples []*xml2triples.Triple, done chan<- struct{}) {
	ch := make(chan struct{})
	tmp := make([]*xml2triples.Triple, len(alltriples))
	copy(tmp, alltriples)
	go xml2triples.SendTriplesAsync(tmp, ch)
	<-ch
	done <- struct{}{} // to sync all triples batches sent
}

func sendTriplesAsync(triplech <-chan *xml2triples.Triple, done chan<- struct{}) {
	ch := make(chan struct{}) // channel to sync all the triples being sent through REST
	batchcount := 0

	triples := make([]*xml2triples.Triple, 0)
	i := 0
	for t := range triplech {
		i++
		triples = append(triples, t)
		if i == 1000 {
			batchcount++
			go simultSendTriples(triples, ch)
			triples = make([]*xml2triples.Triple, 0)
			i = 0
		}

	}
	batchcount++
	go simultSendTriples(triples, ch)
	for i := 0; i < batchcount; i++ {
		<-ch
	}
	done <- struct{}{}
}

// does not ignore any suggested RefIds for objects: the refids in a static file will be cross-referenced
func sendReaderToDataStore(r io.Reader) error {
	// https://stackoverflow.com/a/40526247
	var buffer bytes.Buffer
	ch := make(chan struct{})
	errch := make(chan error)
	triplech := make(chan *xml2triples.Triple) // triples gathered from XML, to be posted to REST
	batchcount := 0
	done := make(chan struct{})
	var sendtripleerror error
	sendtripleerror = nil

	go sendTriplesAsync(triplech, done) // send out triples gathered in main goroutine to REST, in batches
	// gather any errors from errch
	go func() {
		for e := range errch {
			sendtripleerror = e
		}
	}()

	bufferOffset := int64(0)
	tee := io.TeeReader(r, &buffer)
	decoder := xml.NewDecoder(tee)
	for {
		tokenStartOffset := decoder.InputOffset()
		el, err := decoder.Token()
		if err != nil && err != io.EOF {
			log.Println(err)
			return err
		}
		if el == nil {
			break
		}
		switch el.(type) {
		case xml.StartElement:
			err := decoder.Skip()
			if err != nil {
				log.Println(err)
				return err
			}

			// Clear the buffer up to the beginning of this element.
			dif := tokenStartOffset - bufferOffset
			buffer.Next(int(dif))
			bufferOffset += dif

			// Move the offset to the end of this element
			// (so the "clearing" after this switch statement still works).
			tokenStartOffset = decoder.InputOffset()

			// Now output everything in between!
			bodyBytes := buffer.Next(int(tokenStartOffset - bufferOffset))
			b := make([]byte, len(bodyBytes))
			copy(b, string(bodyBytes))

			batchcount++
			go xml2triples.StoreXMLasDBtriplesAsync(b, "SIF", triplech, errch, ch)

			// And now the buffer's up to date.
			bufferOffset = tokenStartOffset
		}

		// Anything content left before this token's start is useless,
		// so clear it out to save memory.
		dif := tokenStartOffset - bufferOffset
		buffer.Next(int(dif))
		bufferOffset += dif

		if err == io.EOF {
			break
		}
		if sendtripleerror != nil {
			break
		}
	}
	for i := 0; i < batchcount; i++ {
		<-ch
	}
	close(triplech)
	close(errch)
	if sendtripleerror != nil {
		log.Println(sendtripleerror)
		return sendtripleerror
	}
	<-done
	return nil
}

func full_object_replace(c echo.Context) bool {
	h := c.Request().Header
	full := false
	repl, ok := h["Replacement"]
	if ok {
		for _, h1 := range repl {
			if h1 == "FULL" {
				full = true
			}
		}
	}
	return full
}

func headerJSON(c echo.Context) bool {
	h := c.Request().Header
	if accept, ok := h["Accept"]; ok {
		for _, a := range accept {
			if a == "application/json" {
				return true
			}
		}
	}
	return false
}

func mustUseAdvisory(c echo.Context) bool {
	h := c.Request().Header
	_, ok := h["Mustuseadvisory"]
	return ok
}

func ids2XMLJSON(ids []string, c echo.Context, buffer bytes.Buffer) (bytes.Buffer, error) {
	json := headerJSON(c)
	if json {
		buffer.Write([]byte("["))
	}
	for i, refid := range ids {
		obj, err := xml2triples.DbTriples2XML(refid, "SIF", true)
		if err != nil {
			log.Println(err)
			return buffer, err
		}
		if json {
			if i > 0 {
				buffer.Write([]byte(","))
			}
			x1, err := xml2triples.XMLtoJSON(obj)
			if err != nil {
				log.Println(err)
				return buffer, err
			}
			obj = x1
		}
		buffer.Write(obj)
	}
	if json {
		buffer.Write([]byte("]"))
	}
	if json {
		c.Response().Header().Set("Content-Type", "application/json")
	} else {
		c.Response().Header().Set("Content-Type", "application/xml")
	}
	return buffer, nil
}

// currently uses the simplistic mechanism of determining whether this is a CREATE ONE or
// CREATE MANY request based on whether the object name is suffixed by an -s
func requestMany(objectname string) bool {
	return strings.HasSuffix(objectname, "s")
}

func Webserver() {
	var err error
	e := echo.New()
	e.Use(middleware.Gzip())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	uuid.Init()
	directoryWatcher()

	e.POST("/sifxml/:object", func(c echo.Context) error {
		if requestMany(c.Param("object")) {
			// TODO: implement POST MANY
			err = fmt.Errorf("Multiple object payloads not yet implemented")
			c.String(http.StatusBadRequest, err.Error())
			return err
		} else {
			object := strings.TrimSuffix(c.Param("object"), "s")
			var bodyBytes []byte
			if c.Request().Body != nil {
				if bodyBytes, err = ioutil.ReadAll(c.Request().Body); err != nil {
					c.String(http.StatusBadRequest, err.Error())
					return err
				}
				objname := xmlobject.FindSubmatch(bodyBytes)
				if objname == nil {
					err = fmt.Errorf("No XML root element")
					c.String(http.StatusBadRequest, err.Error())
					return err
				}
				if string(objname[1]) != object {
					err = fmt.Errorf("%s does not match expected XML root element %s", objname[1], object)
					c.String(http.StatusBadRequest, err.Error())
					return err
				}
				var guid string
				var triples []*xml2triples.Triple
				if guid, triples, err = xml2triples.StoreXMLasDBtriples(bodyBytes, mustUseAdvisory(c), "SIF"); err != nil {
					c.String(http.StatusUnprocessableEntity, err.Error())
					return err
				}
				ch := make(chan struct{})
				go xml2triples.SendTriplesAsync(triples, ch)
				<-ch
				x, err := xml2triples.DbTriples2XML(guid, "SIF", true)
				if err != nil {
					c.String(http.StatusUnprocessableEntity, err.Error())
					return err
				}
				c.Response().Header().Set("Content-Type", "application/xml")
				c.String(http.StatusOK, string(x))
			}
		}
		return nil
	})

	e.PUT("/sifxml/:object/:refid", func(c echo.Context) error {
		refid := c.Param("refid")
		object := c.Param("object")
		var bodyBytes []byte
		if requestMany(object) {
			// TODO: implement PUT MANY
			err = fmt.Errorf("Multiple object payloads not yet implemented")
			c.String(http.StatusBadRequest, err.Error())
			return err
		} else {
			object := strings.TrimSuffix(object, "s")
			if c.Request().Body != nil {
				if bodyBytes, err = ioutil.ReadAll(c.Request().Body); err != nil {
					c.String(http.StatusBadRequest, err.Error())
					return err
				}
				objname := xmlobject.FindSubmatch(bodyBytes)
				if objname == nil {
					err = fmt.Errorf("No XML root element")
					c.String(http.StatusBadRequest, err.Error())
					return err
				}
				if string(objname[1]) != object {
					err = fmt.Errorf("%s does not match expected XML root element %s", objname[1], object)
					c.String(http.StatusBadRequest, err.Error())
					return err
				}
				var err error
				full := full_object_replace(c)
				if full {
					if err = xml2triples.UpdateFullXMLasDBtriples(bodyBytes, refid, "SIF"); err != nil {
						c.String(http.StatusUnprocessableEntity, err.Error())
						return err
					}
				} else {
					if err = xml2triples.UpdatePartialXMLasDBtriples(bodyBytes, refid, "SIF"); err != nil {
						c.String(http.StatusUnprocessableEntity, err.Error())
						return err
					}
				}
				x, err := xml2triples.DbTriples2XML(refid, "SIF", true)
				if err != nil {
					c.String(http.StatusUnprocessableEntity, err.Error())
					return err
				}
				c.Response().Header().Set("Content-Type", "application/xml")
				c.String(http.StatusOK, string(x))
			}
		}
		return nil
	})

	e.GET("/sifxml/:object", func(c echo.Context) error {
		object := strings.TrimSuffix(c.Param("object"), "s")
		objIDs, err := xml2triples.GetAllXMLByObject(object, "SIF")
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return err
		}
		//log.Printf("GETMANY: %+v\n", objIDs)
		for _, refid := range objIDs {
			_, err := xml2triples.DbTriples2XML(refid, "SIF", false)
			if err != nil {
				log.Println(err.Error())
			}
		}
		pr, pw := io.Pipe()
		json := headerJSON(c)
		go func() {
			defer pw.Close()
			if json {
				pw.Write([]byte("["))
			}
			for i, refid := range objIDs {
				x, err := xml2triples.DbTriples2XML(refid, "SIF", false)
				if err != nil {
					log.Println(err.Error())
					c.String(http.StatusInternalServerError, err.Error())
					pw.CloseWithError(err)
					return
				} else {
					if json {
						x1, err := xml2triples.XMLtoJSON(x)
						if err != nil {
							log.Println(err.Error())
							c.String(http.StatusInternalServerError, err.Error())
							pw.CloseWithError(err)
							return
						}
						if i > 0 {
							pw.Write([]byte(","))
						}
						x = x1
					}
					io.Copy(pw, bytes.NewBuffer(x))
				}
			}
			if json {
				pw.Write([]byte("]"))
			}
		}()
		if json {
			c.Stream(http.StatusOK, "application/json", pr)
		} else {
			c.Stream(http.StatusOK, "application/xml", pr)
		}
		pr.Close()
		return err
	})

	e.DELETE("/sifxml/:object/:refid", func(c echo.Context) error {
		// we don't check the type of object we're deleting
		refid := c.Param("refid")
		if requestMany(c.Param("object")) {
			// TODO: implement DELETE MANY
			err = fmt.Errorf("Multiple object payloads not yet implemented")
			c.String(http.StatusBadRequest, err.Error())
			return err
		} else {
			err := xml2triples.DeleteTriplesForRefId(refid, "SIF")
			if err != nil {
				c.String(http.StatusBadRequest, err.Error())
				return err
			}
			c.Response().Header().Set("Content-Type", "application/xml")
			c.String(http.StatusOK, string(refid))
		}
		return nil
	})

	e.GET("/sifxml/:object/:refid", func(c echo.Context) error {
		//object := c.Param("object")
		refid := c.Param("refid")
		json := headerJSON(c)
		x, err := xml2triples.DbTriples2XML(refid, "SIF", true)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return err
		}
		if json {
			x1, err := xml2triples.XMLtoJSON(x)
			if err != nil {
				log.Println(err.Error())
				c.String(http.StatusInternalServerError, err.Error())
				return err
			}
			x = x1
			c.Response().Header().Set("Content-Type", "application/json")
		} else {
			c.Response().Header().Set("Content-Type", "application/xml")
		}
		c.String(http.StatusOK, string(x))
		return nil
	})

	// NSW Digital Classroom adhoc queries
	e.GET("/sifxml/kla2student", func(c echo.Context) error {
		var buffer bytes.Buffer
		kla := c.QueryParam("kla")
		yrlvl := c.QueryParam("yrlvl")
		ids := xml2triples.Kla2student(kla, yrlvl)
		buffer, err := ids2XMLJSON(ids, c, buffer)
		if err != nil {
			log.Println(err.Error())
			c.String(http.StatusBadRequest, err.Error())
			return err
		}
		c.String(http.StatusOK, buffer.String())
		return nil
	})

	e.GET("/sifxml/kla2staff", func(c echo.Context) error {
		var buffer bytes.Buffer
		kla := c.QueryParam("kla")
		yrlvl := c.QueryParam("yrlvl")
		ids := xml2triples.Kla2staff(kla, yrlvl)
		buffer, err := ids2XMLJSON(ids, c, buffer)
		if err != nil {
			log.Println(err.Error())
			c.String(http.StatusBadRequest, err.Error())
			return err
		}
		c.String(http.StatusOK, buffer.String())
		return nil
	})

	e.GET("/sifxml/kla2teachinggroup", func(c echo.Context) error {
		var buffer bytes.Buffer
		kla := c.QueryParam("kla")
		yrlvl := c.QueryParam("yrlvl")
		ids := xml2triples.Kla2teachinggroup(kla, yrlvl)
		buffer, err := ids2XMLJSON(ids, c, buffer)
		if err != nil {
			log.Println(err.Error())
			c.String(http.StatusBadRequest, err.Error())
			return err
		}
		c.String(http.StatusOK, buffer.String())
		return nil
	})

	e.GET("/sifxml/kla2timetablesubject", func(c echo.Context) error {
		var buffer bytes.Buffer
		kla := c.QueryParam("kla")
		yrlvl := c.QueryParam("yrlvl")
		ids := xml2triples.Kla2timetablesubject(kla, yrlvl)
		buffer, err := ids2XMLJSON(ids, c, buffer)
		if err != nil {
			log.Println(err.Error())
			c.String(http.StatusBadRequest, err.Error())
			return err
		}
		c.String(http.StatusOK, buffer.String())
		return nil
	})

	fmt.Println("http://localhost:1492/sifxml is the endpoint for SIF CRUD queries")
	e.Logger.Fatal(e.Start(":1492"))
}
