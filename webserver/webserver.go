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
	//"path/filepath"
	"regexp"
	"strings"
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
			sendXmlToDataStore(event.Path)
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

func sendXmlToDataStore(filename string) error {
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
	// https://stackoverflow.com/a/40526247
	var buffer bytes.Buffer
	bufferOffset := int64(0)

	tee := io.TeeReader(file, &buffer)
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
			guid := uuid.NewV4().String()
			if err := xml2triples.StoreXMLasDBtriples(bodyBytes, guid); err != nil {
				log.Println(err)
				return err
			}

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
	}
	log.Printf("Read in file %s into filestore\n", filename)
	return nil
}

func full_object_replace(c echo.Context) bool {
	// h := c.Request().(*standard.Header).Header
	h := c.Request().Header
	full := false
	repl, ok := h["replacement"]
	if ok {
		for _, h1 := range repl {
			if h1 == "FULL" {
				full = true
			}
		}
	}
	return full
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
		object := strings.TrimSuffix(c.Param("object"), "s")
		var bodyBytes []byte
		if c.Request().Body != nil {
			if bodyBytes, err = ioutil.ReadAll(c.Request().Body); err != nil {
				return err
			}
			objname := xmlobject.FindSubmatch(bodyBytes)
			if objname == nil {
				return fmt.Errorf("No XML root element")
			}
			if string(objname[1]) != object {
				return fmt.Errorf("%s does not match expected XML root element %s", objname[1], object)
			}
			guid := uuid.NewV4().String()
			if err := xml2triples.StoreXMLasDBtriples(bodyBytes, guid); err != nil {
				return err
			}
			x, err := xml2triples.DbTriples2XML(guid)
			if err != nil {
				return err
			}
			c.Response().Header().Set("Content-Type", "application/xml")
			c.String(http.StatusOK, string(x))
		}
		return nil
	})

	e.PUT("/sifxml/:object/:refid", func(c echo.Context) error {
		refid := c.Param("refid")
		var bodyBytes []byte
		object := strings.TrimSuffix(c.Param("object"), "s")
		if c.Request().Body != nil {
			if bodyBytes, err = ioutil.ReadAll(c.Request().Body); err != nil {
				return err
			}
			objname := xmlobject.FindSubmatch(bodyBytes)
			if objname == nil {
				return fmt.Errorf("No XML root element")
			}
			if string(objname[1]) != object {
				return fmt.Errorf("%s does not match expected XML root element %s", objname[1], object)
			}
			var err error
			full := full_object_replace(c)
			if full {
				if err = xml2triples.UpdateFullXMLasDBtriples(bodyBytes, refid); err != nil {
					return err
				}
			} else {
				if err = xml2triples.UpdatePartialXMLasDBtriples(bodyBytes, refid); err != nil {
					return err
				}
			}
			x, err := xml2triples.DbTriples2XML(refid)
			if err != nil {
				return err
			}
			c.Response().Header().Set("Content-Type", "application/xml")
			c.String(http.StatusOK, string(x))
		}
		return nil
	})

	e.GET("/sifxml/:object", func(c echo.Context) error {
		object := strings.TrimSuffix(c.Param("object"), "s")
		objIDs, err := xml2triples.GetAllXMLByObject(object)
		if err != nil {
			return err
		}
		//log.Printf("%+v\n", objIDs)
		pr, pw := io.Pipe()
		go func() {
			for _, refid := range objIDs {
				x, err := xml2triples.DbTriples2XML(refid)
				if err != nil {
					pw.CloseWithError(err)
				}
				_, err = pw.Write(x)
				if err != nil {
					pw.CloseWithError(err)
				}
			}
			pw.Close()
		}()
		c.Response().Header().Set("Content-Type", "application/xml")
		c.Stream(http.StatusOK, "application/xml", pr)
		return err
	})

	e.DELETE("/sifxml/:object/:refid", func(c echo.Context) error {
		//object := c.Param("object")
		refid := c.Param("refid")
		err := xml2triples.DeleteTriplesForRefId(refid)
		if err != nil {
			return err
		}
		c.Response().Header().Set("Content-Type", "application/xml")
		c.String(http.StatusOK, string(refid))
		return nil
	})

	e.GET("/sifxml/:object/:refid", func(c echo.Context) error {
		//object := c.Param("object")
		refid := c.Param("refid")
		x, err := xml2triples.DbTriples2XML(refid)
		if err != nil {
			return err
		}
		c.Response().Header().Set("Content-Type", "application/xml")
		c.String(http.StatusOK, string(x))
		return nil
	})

	fmt.Println("http://localhost:1492/sifxml is the endpoint for SIF CR[UD] queries")
	e.Logger.Fatal(e.Start(":1492"))
}
