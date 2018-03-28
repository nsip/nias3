package webserver

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/nsip/nias3/xml2triples"
	"github.com/twinj/uuid"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

var xmlobject = regexp.MustCompile(`(?s:^\s*<([^> ]+))`)

func Webserver() {
	e := echo.New()
	e.Use(middleware.Gzip())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	uuid.Init()
	var err error

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
