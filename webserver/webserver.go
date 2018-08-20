package webserver

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/nsip/nias3/xml2triples"
	"github.com/radovskyb/watcher"
	"github.com/tidwall/gjson"
	"github.com/twinj/uuid"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

// root element in an XML string
var xmlobject = regexp.MustCompile(`(?s:^\s*<([^> ]+))`)

const DROPBOX = "./in"

// write actions in REST
const (
	POST = iota
	PUT
	PATCH
	DELETE
)

// Watch a dropbox directory for new files, and post their contents to REST
func directoryWatcher(dropbox string) {
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

	if err = w.AddRecursive(dropbox); err != nil {
		log.Fatalln(err)
	}
	go func() {
		if err := w.Start(time.Millisecond * 100); err != nil {
			log.Fatalln(err)
		}
	}()
}

// Post the contents of an XML file to REST
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
	ret, err1, err := sendReaderToDataStore(file, true, POST)
	if err != nil {
		return err
	}
	log.Println(createResponse(ret))
	if err1 != nil {
		return err1
	}
	log.Printf("Read in file %s into filestore\n", filename)
	return nil
}

// Post the contents of a JSON file to REST
func SendJsonToDataStore(filename string, context string) error {
	fi, err := os.Lstat(filename)
	if err != nil {
		return err
	}
	if fi.Mode().IsDir() {
		return fmt.Errorf("%s is a directory", filename)
	}
	if !strings.HasSuffix(filename, ".json") {
		return fmt.Errorf("%s is not an JSON file", filename)
	}
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("Cannot read in file %s\n", filename)
		return err
	}
	ret, err1, err := sendReaderToDataStoreJSON(file, context)
	if err != nil {
		return err
	}
	log.Println(createResponse(ret))
	if err1 != nil {
		return err1
	}
	log.Printf("Read in file %s into filestore\n", filename)
	return nil
}

// url: e.g. http://hits.nsip.edu.au/SIF3InfraREST/hits/requests/SchoolInfos?navigationPage=1&navigationPageSize=5&access_token=ZmZhODMzNjEtMGExOC00NDk5LTgyNjMtYjMwNjI4MGRjZDRlOmYxYzA1NjNhOWIzZTQyMGJiMDdkYTJkOTBkYjQ3OWVm&authenticationMethod=Basic
// do an HTTP GET on a URL, and post the XML respone to REST
func SIFGetToDataStore(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, err1, err := sendReaderToDataStore(resp.Body, true, POST)
	if err != nil {
		return err
	}
	if err1 != nil {
		return err1
	}
	return nil
}

// Asynchronously post a slice of triples to REST
func simultSendTriples(alltriples []*xml2triples.Triple, context string, done chan<- struct{}) {
	ch := make(chan struct{})
	tmp := make([]*xml2triples.Triple, len(alltriples))
	copy(tmp, alltriples)
	go xml2triples.SendTriplesAsync(tmp, context, ch)
	<-ch
	done <- struct{}{} // to sync all triples batches sent
}

const BATCHSIZE = 1000

// Asynchronously post a channel of triples to REST, in batches of BATCHSIZE triples
func sendTriplesAsync(triplech <-chan *xml2triples.Triple, context string, done chan<- struct{}) {
	ch := make(chan struct{}) // channel to sync all the triples being sent through REST
	batchcount := 0

	triples := make([]*xml2triples.Triple, 0)
	i := 0
	for t := range triplech {
		i++
		triplesSent++
		triples = append(triples, t)
		if i == BATCHSIZE {
			batchcount++
			go simultSendTriples(triples, context, ch)
			triples = make([]*xml2triples.Triple, 0)
			i = 0
		}

	}
	batchcount++
	go simultSendTriples(triples, context, ch)
	for i := 0; i < batchcount; i++ {
		<-ch // synchronise all batches
	}
	done <- struct{}{}
}

var triplesSent int

// Send a file of XML to REST, and return the response for each object in the XML.
// The file is assumed to be well-formed XML, containing either a single or multiple objects. If it contains
// multiple objects, these are to be contained in an element with a plural name (e.g. StudentPersonals),
// or else the conventional wrapper <sif>
// If mustUseAdvisory = true, does not ignore any suggested RefIds for objects:
// the refids in a static file will be cross-referenced.
// Verb is the REST operation to be done with the file: POST, PUT, PATCH, DELETE
// Returns the slice of SIF response messages for each object in the payload,
// error for failing to send any triples (resource collision), and generic error
func sendReaderToDataStore(r io.Reader, mustUseAdvisory bool, verb int) ([]xml2triples.SifResponse, error, error) {
	// https://stackoverflow.com/a/40526247
	var buffer bytes.Buffer
	ch := make(chan struct{})                     // channel of batches of tuples being sent
	refidch := make(chan xml2triples.SifResponse) // SIF response for each object sent
	triplech := make(chan *xml2triples.Triple)    // triples gathered from XML, to be posted to REST
	batchcount := 0
	done := make(chan struct{})     // sent all SIF objects pending
	ret_done := make(chan struct{}) // received all SIF responses pending
	var sendtripleerror error
	sendtripleerror = nil
	var genericerror error
	genericerror = nil
	ret := make([]xml2triples.SifResponse, 0) // SIF response for each object sent
	triplesSent = 0
	objs := 0

	// send out triples gathered in main goroutine to REST, in batches
	go sendTriplesAsync(triplech, "SIF", done)
	// read SIF response for each object sent
	go func() {
		for e := range refidch {
			ret = append(ret, e)
			if e.Error != nil {
				log.Println(sendtripleerror)
				sendtripleerror = e.Error
			}
		}
		ret_done <- struct{}{}
	}()

	bufferOffset := int64(0)
	tee := io.TeeReader(r, &buffer)
	decoder := xml.NewDecoder(tee)
	for {
		tokenStartOffset := decoder.InputOffset()
		el, err := decoder.Token()
		if err != nil && err != io.EOF {
			log.Println(err)
			genericerror = err
			break
		}
		if el == nil {
			break
		}
		switch se := el.(type) {
		case xml.StartElement:
			// skip multiple-object container tags, such as StudentPersonals
			if requestMany(se.Name.Local) {
				continue
			}
			objs++
			err := decoder.Skip()
			if err != nil {
				log.Println(err)
				genericerror = err
				break
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
			switch verb {
			case POST:
				go xml2triples.StoreXMLasDBtriplesAsync(b, mustUseAdvisory, "SIF", triplech, refidch, ch)
			case PUT:
				go xml2triples.UpdateFullXMLasDBtriplesAsync(b, "SIF", triplech, refidch, ch)
			case PATCH:
				go xml2triples.UpdatePartialXMLasDBtriplesAsync(b, "SIF", triplech, refidch, ch)
			case DELETE:
				go xml2triples.DeleteTriplesForRefIdAsync(b, "SIF", triplech, refidch, ch)
			default:
				go xml2triples.StoreXMLasDBtriplesAsync(b, mustUseAdvisory, "SIF", triplech, refidch, ch)
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
		/*
			if sendtripleerror != nil {
				break
			}
		*/
	}
	for i := 0; i < batchcount; i++ {
		<-ch // sync all tuples batches sent
	}
	close(triplech)
	close(refidch)
	<-ret_done
	<-done
	log.Printf("sent %d triples in %d SIF XML objects\n", triplesSent, objs)
	return ret, sendtripleerror, genericerror
}

// Send a file of JSON to REST
// The file is assumed to be a well-formed JSON array.
// Currently does only POST.
func sendReaderToDataStoreJSON(bodyBytes []byte, context string) ([]xml2triples.SifResponse, error, error) {
	ch := make(chan struct{})                     // channel of batches of tuples being sent
	refidch := make(chan xml2triples.SifResponse) // SIF response for each object sent
	triplech := make(chan *xml2triples.Triple)    // triples gathered from XML, to be posted to REST
	batchcount := 0
	done := make(chan struct{})     // sent all SIF objects pending
	ret_done := make(chan struct{}) // received all SIF responses pending
	var sendtripleerror error
	sendtripleerror = nil
	var genericerror error
	genericerror = nil
	ret := make([]xml2triples.SifResponse, 0) // SIF response for each object sent
	triplesSent = 0
	objs := 0

	m, ok := gjson.Parse(string(bodyBytes)).Value().([]interface{})
	if !ok {
		return ret, errors.New("malformed JSON payload"), nil
	}

	// send out triples gathered in main goroutine to REST, in batches
	go sendTriplesAsync(triplech, context, done)
	// read SIF response for each object sent
	go func() {
		for e := range refidch {
			ret = append(ret, e)
			if e.Error != nil {
				log.Println(sendtripleerror)
				sendtripleerror = e.Error
			}
		}
		ret_done <- struct{}{}
	}()

	for _, elem := range m {
		objs++
		jsonBytes, err := json.Marshal(elem)
		//log.Println(string(jsonBytes))
		if err != nil {
			log.Println(err)
			genericerror = err
			break
		} else {
			batchcount++
			go xml2triples.StoreJSONasDBtriplesAsync(jsonBytes, context, triplech, refidch, ch)
		}
	}
	for i := 0; i < batchcount; i++ {
		<-ch // sync all tuples batches sent
	}
	close(triplech)
	close(refidch)
	<-ret_done
	<-done
	log.Printf("sent %d triples in %d JSON objects\n", triplesSent, objs)
	return ret, sendtripleerror, genericerror
}

// Is this update request a PUT or a PATCH?
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

// Should the response be in JSON?
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

// Is the mustUseAdvisory field on in the request?
func mustUseAdvisory(c echo.Context) bool {
	h := c.Request().Header
	_, ok := h["Mustuseadvisory"]
	return ok
}

// Is the methodOverride field in the request set to DELETE?
func deleteMany(c echo.Context) bool {
	h := c.Request().Header
	v, ok := h["Methodoverride"]
	return ok && len(v) > 0 && v[0] == "DELETE"
}

// Query REST for the XML corresponding to a set of RefIDs. If content has been requested in JSON,
// convert XML to Goessner JSON. Return XML or JSON, and sets HTTP response content type.
// If XML, wrap the objects with the tag given in wrapper.
func ids2XMLJSON(ids []string, wrapper string, c echo.Context, buffer bytes.Buffer) (bytes.Buffer, error) {
	json := headerJSON(c)
	if json {
		buffer.Write([]byte("["))
	} else { // xml
		buffer.Write([]byte("<" + wrapper + ">\n"))
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
		c.Response().Header().Set("Content-Type", "application/json")
	} else {
		buffer.Write([]byte("</" + wrapper + ">\n"))
		c.Response().Header().Set("Content-Type", "application/xml")
	}
	return buffer, nil
}

// currently uses the simplistic mechanism of determining whether this is a CREATE ONE or
// CREATE MANY request based on whether the object name is suffixed by an -s
func requestMany(objectname string) bool {
	return strings.HasSuffix(objectname, "s") || objectname == "sif"
}

// formulate SIF Create Response to POST MANY, based on the responses received
func createResponse(responses []xml2triples.SifResponse) string {
	var bld strings.Builder
	bld.WriteString("<createResponse>\n  <creates>\n")
	for _, r := range responses {
		if r.Error == nil {
			bld.WriteString(fmt.Sprintf("    <create id=\"%s\" advisoryId=\"%s\" statusCode = \"201\"/>\n", r.RefId, r.AdvisoryId))
		} else {
			bld.WriteString(fmt.Sprintf("    <create advisoryId=\"%s\" statusCode = \"409\">\n", r.AdvisoryId))
			bld.WriteString(fmt.Sprintf("      <error id=\"%s\">\n", strings.ToUpper(uuid.NewV4().String())))
			bld.WriteString(fmt.Sprintf("        <code>409</code>\n        <scope>StateConflict</scope>\n        <message>Object with that RefId already exists</message>\n      </error>\n    </create>\n"))
		}
	}
	bld.WriteString("  <creates>\n<createResponse>\n")
	return bld.String()
}

// formulate SIF Update Response to PUT MANY, based on the responses received
func updateResponse(responses []xml2triples.SifResponse) string {
	var bld strings.Builder
	bld.WriteString("<updateResponse>\n  <updates>\n")
	for _, r := range responses {
		if r.Error == nil {
			bld.WriteString(fmt.Sprintf("    <update id=\"%s\" statusCode = \"200\"/>\n", r.RefId))
		} else {
			bld.WriteString(fmt.Sprintf("    <update id=\"%s\" statusCode = \"404\">\n", r.RefId))
			bld.WriteString(fmt.Sprintf("      <error id=\"%s\">\n", strings.ToUpper(uuid.NewV4().String())))
			bld.WriteString(fmt.Sprintf("        <code>404</code>\n        <scope>Not Found</scope>\n        <message>Object with that RefId not found</message>\n      </error>\n    </update>\n"))
		}
	}
	bld.WriteString("  <updates>\n<updateResponse>\n")
	return bld.String()
}

// formulate SIF Update Response to DELETE MANY, based on the responses received
func deleteResponse(responses []xml2triples.SifResponse) string {
	var bld strings.Builder
	bld.WriteString("<deleteResponse>\n  <deletes>\n")
	for _, r := range responses {
		if r.Error == nil {
			bld.WriteString(fmt.Sprintf("    <delete id=\"%s\" statusCode = \"200\"/>\n", r.RefId))
		} else {
			bld.WriteString(fmt.Sprintf("    <delete id=\"%s\" statusCode = \"404\">\n", r.RefId))
			bld.WriteString(fmt.Sprintf("      <error id=\"%s\">\n", strings.ToUpper(uuid.NewV4().String())))
			bld.WriteString(fmt.Sprintf("        <code>404</code>\n        <scope>Not Found</scope>\n        <message>Object with that RefId not found</message>\n      </error>\n    </delete>\n"))
		}
	}
	bld.WriteString("  <deletes>\n<deleteResponse>\n")
	return bld.String()
}

// time to wait for tuple request to be processed, for single synchronous object request
const STOREWAIT = 200 * time.Millisecond

func Webserver() {
	var err error
	e := echo.New()
	e.Use(middleware.Gzip())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	uuid.Init()
	directoryWatcher(DROPBOX)

	// NSW Digital Classroom adhoc queries
	e.GET("/sifxml/kla2student", func(c echo.Context) error {
		var buffer bytes.Buffer
		kla := c.QueryParam("kla")
		yrlvl := c.QueryParam("yrlvl")
		ids := xml2triples.Kla2student(kla, yrlvl)
		buffer, err := ids2XMLJSON(ids, "sif", c, buffer)
		if err != nil {
			log.Println(err.Error())
			c.String(http.StatusBadRequest, err.Error())
			return err
		}
		c.String(http.StatusOK, buffer.String())
		return nil
	})

	// NSW Digital Classroom adhoc queries
	e.GET("/sifxml/kla2staff", func(c echo.Context) error {
		var buffer bytes.Buffer
		kla := c.QueryParam("kla")
		yrlvl := c.QueryParam("yrlvl")
		ids := xml2triples.Kla2staff(kla, yrlvl)
		buffer, err := ids2XMLJSON(ids, "sif", c, buffer)
		if err != nil {
			log.Println(err.Error())
			c.String(http.StatusBadRequest, err.Error())
			return err
		}
		c.String(http.StatusOK, buffer.String())
		return nil
	})

	// NSW Digital Classroom adhoc queries
	e.GET("/sifxml/kla2teachinggroup", func(c echo.Context) error {
		var buffer bytes.Buffer
		kla := c.QueryParam("kla")
		yrlvl := c.QueryParam("yrlvl")
		ids := xml2triples.Kla2teachinggroup(kla, yrlvl)
		buffer, err := ids2XMLJSON(ids, "sif", c, buffer)
		if err != nil {
			log.Println(err.Error())
			c.String(http.StatusBadRequest, err.Error())
			return err
		}
		c.String(http.StatusOK, buffer.String())
		return nil
	})

	// NSW Digital Classroom adhoc queries
	e.GET("/sifxml/kla2timetablesubject", func(c echo.Context) error {
		var buffer bytes.Buffer
		kla := c.QueryParam("kla")
		yrlvl := c.QueryParam("yrlvl")
		ids := xml2triples.Kla2timetablesubject(kla, yrlvl)
		buffer, err := ids2XMLJSON(ids, "sif", c, buffer)
		if err != nil {
			log.Println(err.Error())
			c.String(http.StatusBadRequest, err.Error())
			return err
		}
		c.String(http.StatusOK, buffer.String())
		return nil
	})

	// POST ONE, e.g. StaffPersonals/StaffPersonal
	e.POST("/sifxml/:objects/:refid", func(c echo.Context) error {
		var bodyBytes []byte
		if c.Request().Body != nil {
			if bodyBytes, err = ioutil.ReadAll(c.Request().Body); err != nil {
				c.String(http.StatusBadRequest, err.Error())
				return err
			}
			objects := c.Param("objects")
			// https://github.com/labstack/echo/issues/1139 : have to name this parameter "refid", to match
			// other path instances, but it's actually the object name
			object := c.Param("refid")
			if object != strings.TrimSuffix(objects, "s") {
				err = errors.New("POST ONE request path is malformed")
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
			go xml2triples.SendTriplesAsync(triples, "SIF", ch)
			<-ch
			// allow time for triples to be stored
			//time.Sleep(200 * time.Millisecond)
			time.Sleep(STOREWAIT)
			x, err := xml2triples.DbTriples2XML(guid, "SIF", true)
			if err != nil {
				c.String(http.StatusUnprocessableEntity, err.Error())
				return err
			}
			c.Response().Header().Set("Content-Type", "application/xml")
			c.String(http.StatusCreated, string(x))
		}
		return nil
	})

	// is interpreted as POST MANY, but will process single object payload
	e.POST("/sifxml/:objects", func(c echo.Context) error {
		var bodyBytes []byte
		if c.Request().Body != nil {
			if bodyBytes, err = ioutil.ReadAll(c.Request().Body); err != nil {
				c.String(http.StatusBadRequest, err.Error())
				return err
			}
			// we don't validate the object name against each element
			// we don't validate the well-formedness of the payload in advance
			ret, _, genericerr := sendReaderToDataStore(bytes.NewReader(bodyBytes), mustUseAdvisory(c), POST)
			if genericerr != nil {
				c.String(http.StatusBadRequest, genericerr.Error())
				return err
			}
			c.Response().Header().Set("Content-Type", "application/xml")
			c.String(http.StatusOK, createResponse(ret))
		}
		return nil
	})

	// POST MANY, array of xAPI JSON objects. Expect input to be JSON array
	e.POST("/xapi", func(c echo.Context) error {
		var bodyBytes []byte
		if c.Request().Body != nil {
			if bodyBytes, err = ioutil.ReadAll(c.Request().Body); err != nil {
				c.String(http.StatusBadRequest, err.Error())
				return err
			}
			_, _, genericerr := sendReaderToDataStoreJSON(bodyBytes, "xAPI")
			if genericerr != nil {
				c.String(http.StatusBadRequest, genericerr.Error())
				return err
			}
			c.Response().Header().Set("Content-Type", "application/xml")
			c.String(http.StatusOK, "OK")
		}
		return nil
	})

	// PUT ONE
	e.PUT("/sifxml/:objects/:refid", func(c echo.Context) error {
		refid := c.Param("refid")
		object := strings.TrimSuffix(c.Param("objects"), "s")
		var bodyBytes []byte
		if c.Request().Body != nil {
			if bodyBytes, err = ioutil.ReadAll(c.Request().Body); err != nil {
				c.String(http.StatusBadRequest, err.Error())
				return err
			}
			full := full_object_replace(c)
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
			// allow time for triples to be stored
			time.Sleep(200 * time.Millisecond)
			x, err := xml2triples.DbTriples2XML(refid, "SIF", true)
			if err != nil {
				c.String(http.StatusUnprocessableEntity, err.Error())
				return err
			}
			c.Response().Header().Set("Content-Type", "application/xml")
			c.String(http.StatusOK, string(x))
		}
		return nil
	})

	// PUT MANY, DELETE MANY
	e.PUT("/sifxml/:objects", func(c echo.Context) error {
		// object := c.Param("object")
		var bodyBytes []byte
		if c.Request().Body != nil {
			if bodyBytes, err = ioutil.ReadAll(c.Request().Body); err != nil {
				c.String(http.StatusBadRequest, err.Error())
				return err
			}
			full := full_object_replace(c)
			// we don't validate the object name against each element
			// we don't validate the well-formedness of the payload in advance
			verb := PATCH
			if full {
				verb = PUT
			} else if deleteMany(c) {
				verb = DELETE
				log.Println("DELETE")
			} else {
				log.Println("PATCH")
			}
			ret, _, genericerr := sendReaderToDataStore(bytes.NewReader(bodyBytes), true, verb)
			if genericerr != nil {
				c.String(http.StatusBadRequest, genericerr.Error())
				return err
			}
			c.Response().Header().Set("Content-Type", "application/xml")
			if deleteMany(c) {
				c.String(http.StatusOK, deleteResponse(ret))
			} else {
				c.String(http.StatusOK, updateResponse(ret))
			}
		}
		return nil
	})

	// GET ONE
	e.GET("/sifxml/:objects/:refid", func(c echo.Context) error {
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

	// GET MANY, as JSON or XML
	e.GET("/sifxml/:objects", func(c echo.Context) error {
		object := strings.TrimSuffix(c.Param("objects"), "s")
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

	// GET MANY xAPI aligned to SIF RefId
	e.GET("/sif2xapi/:objects", func(c echo.Context) error {
		refid := c.Param("objects")
		tuples := xml2triples.Sif2Xapi(refid)
		ret, err := xml2triples.Tuples2JSON(tuples)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return err
		}
		c.Response().Header().Set("Content-Type", "application/json")
		c.String(http.StatusOK, ret)
		return nil
	})

	// DELETE ONE
	e.DELETE("/sifxml/:objects/:refid", func(c echo.Context) error {
		// we don't check the type of object we're deleting
		refid := c.Param("refid")
		err := xml2triples.DeleteTriplesForRefId(refid, "SIF")
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return err
		}
		c.Response().Header().Set("Content-Type", "application/xml")
		c.String(http.StatusOK, string(refid))
		return nil
	})

	fmt.Println("http://localhost:1492/sifxml is the endpoint for SIF CRUD queries")
	e.Logger.Fatal(e.Start(":1492"))
}
