package xml2triples

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/clbanning/mxj"
	"github.com/nsip/nias3/config"
	"github.com/tidwall/sjson"
	"github.com/twinj/uuid"
)

var conf = config.LoadConfig()
var baseUrl string
var client = initClient()

// channel to throttle simultaneous http posts
var limitch chan struct{}

// max number of simultaneous http posts allowed
const HTTPTHREADS = 100

// https://forfuncsake.github.io/post/2017/08/trust-extra-ca-cert-in-go-app/
// Create an HTTP client to pass on messages to Nias3Engine
// Permit self-signed certificate out of Nias3Engine for HTTPS
func initClient() *http.Client {
	limitch = make(chan struct{}, HTTPTHREADS)
	localCertFile := conf.N3EngineTLSCert
	if len(localCertFile) == 0 {
		log.Println("Using HTTP")
		baseUrl = fmt.Sprintf("http://localhost:%d", conf.N3EngineWebport)
		return &http.Client{}
	}
	rootCAs, _ := x509.SystemCertPool()
	if rootCAs == nil {
		rootCAs = x509.NewCertPool()
	}
	certs, err := ioutil.ReadFile(localCertFile)
	if err != nil {
		log.Fatalf("Failed to append %q to RootCAs: %v", localCertFile, err)
	}
	if ok := rootCAs.AppendCertsFromPEM(certs); !ok {
		log.Println("No certs appended, using system certs only")
	}

	// Trust the augmented cert pool in our client
	config := &tls.Config{
		InsecureSkipVerify: true,
		RootCAs:            rootCAs,
	}
	tr := &http.Transport{TLSClientConfig: config}
	baseUrl = fmt.Sprintf("https://localhost:%d", conf.N3EngineWebport)
	log.Println("Using HTTP/2")
	return &http.Client{Transport: tr}
}

// Change MXJ JSON tags to the JSON XML tags used in our internal SIF XML representation
func changeJSONTags(j string) string {
	return strings.Replace(j, `"#text":`, `"Value":`, -1)
}

// Convert an MXJ map, as retrieved from tuples, into a byte encoding of a SIF XML object.
// If stripempty, make an effort to remove empty tags in the XML.
func Map2SIFXML(m mxj.Map, stripempty bool) ([]byte, error) {
	root, err := m.Root()
	if err != nil {
		return nil, err
	}
	// The root element is the name of the object; we process beneath it
	m02 := m[root].(map[string]interface{})
	j, err := json.Marshal(m02)
	if err != nil {
		return nil, err
	}
	ret, err := Root2SIF(root, []byte(changeJSONTags(string(j))))
	if err != nil {
		return nil, err
	}
	if stripempty {
		ret = stripEmptyTags(ret)
	}
	ret = append(ret, []byte("\n")...)
	return ret, nil
}

type Triple struct {
	Subject   string
	Object    string
	Predicate string
	Context   string
}

// Send a slice of triples synchronously to Nias3Engine
// Always add limitch <- struct{}{} before calling
func send_triples(triples []*Triple, context string) {
	if len(triples) == 0 {
		<-limitch
		return
	}
	json, err := json.Marshal(triples)
	if err != nil {
		panic(err)
	}
	//log.Println(string(json))
	<-limitch
	req, err := http.NewRequest("POST", baseUrl+"/tuples/"+context, bytes.NewBuffer(json))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}

// Send a slice of triples asynchronously to Nias3Engine
func SendTriplesAsync(triples []*Triple, context string, done chan<- struct{}) {
	limitch <- struct{}{}
	send_triples(triples, context)
	done <- struct{}{}
}

// Retrieve xAPI tuples aligned to SIF refid. Synchronous.
func Sif2Xapi(refid string) []*Triple {
	req, err := http.NewRequest("GET", baseUrl+"/sif2xapi/"+url.PathEscape(refid), nil)
	if err != nil {
		panic(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	ret := make([]*Triple, 0)
	json.NewDecoder(resp.Body).Decode(&ret)
	return ret
}

// Check if S(P(O)) key prefix is on Nias3Engine Hexastore. Adds context prefix to key before checking. Synchronous.
func hasKey(keyprefix string, context string) bool {
	keyprefix1 := fmt.Sprintf("c:%s %s", strconv.Quote(context), keyprefix)
	req, err := http.NewRequest("GET", baseUrl+"/HasKey/"+url.PathEscape(keyprefix1), nil)
	if err != nil {
		panic(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	return resp.StatusCode == 200
}

// Check if S(P(O)) key prefix is on Nias3Engine Hexastore. Adds context prefix to key before checking. Asynchronous.
// Always add limitch <- struct{}{} before calling
func hasKeyAsync(keyprefix string, context string, ch chan<- bool) {
	keyprefix1 := fmt.Sprintf("c:%s %s", strconv.Quote(context), keyprefix)
	req, err := http.NewRequest("GET", baseUrl+"/HasKey/"+url.PathEscape(keyprefix1), nil)
	if err != nil {
		panic(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	<-limitch
	ch <- resp.StatusCode == 200
}

// Retrieve tuples from Hexastore matching a key prefix (involving a subset of s: o: p:; the c: prefix
// will be added here). Synchronous.
func getTuples(keyprefix string, context string) []*Triple {
	keyprefix1 := fmt.Sprintf("c:%s %s", strconv.Quote(context), keyprefix)
	req, err := http.NewRequest("GET", baseUrl+"/tuple/"+url.PathEscape(keyprefix1), nil)
	if err != nil {
		panic(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	ret := make([]*Triple, 0)
	json.NewDecoder(resp.Body).Decode(&ret)
	return ret
}

// Retrieve tuples from Influx store matching a subject
func getTuplesInflux(subject string, context string) []*Triple {
	req, err := http.NewRequest("GET", baseUrl+"/influxtuple/"+url.PathEscape(subject), nil)
	if err != nil {
		panic(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	ret := make([]*Triple, 0)
	json.NewDecoder(resp.Body).Decode(&ret)
	return ret
}

// check if there are any entries in the Hexastore with a non-empty object (i.e. not deleted), for the given subject; asynchronous
func existsRefid(refid string, context string, ch chan<- bool) {
	tuples := getTuples(fmt.Sprintf("s:%s o:", strconv.Quote(refid)), context)
	<-limitch
	ret := false
	for _, t := range tuples {
		if t.Object != "" {
			ret = true
			break
		}
	}
	ch <- ret
}

/* NSW Digital Classroom hard coded query: retrieve students studying KLA (key learning area) in given year level */
func Kla2student(kla string, yrlvl string) []string {
	req, err := http.NewRequest("GET", baseUrl+"/kla2student?kla="+url.PathEscape(kla)+"&yrlvl="+url.PathEscape(yrlvl), nil)
	if err != nil {
		panic(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	ret := make([]string, 0)
	json.NewDecoder(resp.Body).Decode(&ret)
	return ret
}

/* NSW Digital Classroom hard coded query: retrieve staff teaching KLA (key learning area) in given year level */
func Kla2staff(kla string, yrlvl string) []string {
	req, err := http.NewRequest("GET", baseUrl+"/kla2staff?kla="+url.PathEscape(kla)+"&yrlvl="+url.PathEscape(yrlvl), nil)
	if err != nil {
		panic(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	ret := make([]string, 0)
	json.NewDecoder(resp.Body).Decode(&ret)
	return ret
}

/* NSW Digital Classroom hard coded query: retrieve teaching groups about KLA (key learning area) in given year level */
func Kla2teachinggroup(kla string, yrlvl string) []string {
	req, err := http.NewRequest("GET", baseUrl+"/kla2teachinggroup?kla="+url.PathEscape(kla)+"&yrlvl="+url.PathEscape(yrlvl), nil)
	if err != nil {
		panic(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	ret := make([]string, 0)
	json.NewDecoder(resp.Body).Decode(&ret)
	return ret
}

/* NSW Digital Classroom hard coded query: retrieve timetable subjects about KLA (key learning area) in given year level */
func Kla2timetablesubject(kla string, yrlvl string) []string {
	req, err := http.NewRequest("GET", baseUrl+"/kla2timetablesubject?kla="+url.PathEscape(kla)+"&yrlvl="+url.PathEscape(yrlvl), nil)
	if err != nil {
		panic(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	ret := make([]string, 0)
	json.NewDecoder(resp.Body).Decode(&ret)
	return ret
}

// Delete all triples with given refis as subject; synchronous
func DeleteTriplesForRefId(refid string, context string) error {
	triples, err := makeDeleteTriplesForRefId(refid, context)
	if err != nil {
		return err
	}
	send_triples(triples, context)
	return nil
}

type DeleteRequest struct {
	XMLName            xml.Name           `xml:"deleteRequest"`
	DeletesRequestElem DeletesRequestElem `xml:"deletes"`
}

type DeletesRequestElem struct {
	Delete []DeleteRequestElem `xml:"delete"`
}

type DeleteRequestElem struct {
	Id string `xml:"id,attr"`
}

// processes deleteRequest object
// Delete all triples with their subject being one of the refids given in the SIF deleteRequest object deleteReq; asynchronous.
// Outputs SIF responses for each refid to errch
// Outputs a list of tuples to be sent to Hexastore to outch
func DeleteTriplesForRefIdAsync(deleteReq []byte, context string, outch chan<- *Triple, errch chan<- SifResponse, ch chan<- struct{}) {
	v := DeleteRequest{}
	err := xml.Unmarshal(deleteReq, &v)
	if err != nil {
		errch <- SifResponse{RefId: "", Error: errors.New("Request Object malformed")}
		ch <- struct{}{}
		return
	}
	for _, del := range v.DeletesRequestElem.Delete {
		refid := del.Id
		triples, err := makeDeleteTriplesForRefId(refid, context)
		if err != nil {
			errch <- SifResponse{RefId: "", Error: err}
		} else if len(triples) == 0 {
			errch <- SifResponse{RefId: "", Error: errors.New("RefId not found")}
		} else {
			for _, t := range triples {
				outch <- t
			}
			errch <- SifResponse{AdvisoryId: refid, RefId: refid, Error: nil}
		}
	}
	ch <- struct{}{}
}

// Generates the triples to be sent to the Hexastore to realise deletion of the object with the given refid
func makeDeleteTriplesForRefId(refid string, context string) ([]*Triple, error) {
	triples := getTuples(fmt.Sprintf("s:%s ", strconv.Quote(fmt.Sprintf("%v", refid))), context)
	ret := make([]*Triple, 0)
	seen := make(map[string]bool)
	for _, t := range triples {
		key := fmt.Sprintf("s:%s p:%s", strconv.Quote(t.Subject), strconv.Quote(t.Predicate))
		if _, ok := seen[key]; !ok {
			t.Object = "" // tuples are deleted by setting their object to empty
			ret = append(ret, t)
			seen[key] = true
		}
	}
	limitch <- struct{}{}
	return ret, nil
}

// Update an object, with Full Replacement on the Hexastore; synchronous
// Deletes all tuples for the refid from Hexastore, then creates tuples for the replacement object,
// and sends them to Hexastore as an ordered set of tuples
func UpdateFullXMLasDBtriples(object []byte, requested_refid string, context string) error {
	m, err := mxj.NewMapXml(object)
	if err != nil {
		return err
	}
	refid, err := m.ValueForPath("*.-RefId")
	if err != nil {
		refid = requested_refid
	}
	triples, err := makeDeleteTriplesForRefId(refid.(string), context)
	for _, n := range m.LeafNodes() {
		triples = append(triples, &Triple{Subject: refid.(string), Predicate: n.Path, Object: fmt.Sprintf("%v", n.Value), Context: context})
	}
	limitch <- struct{}{}
	send_triples(triples, context)
	return nil
}

// Update an object, with Full Replacement on the Hexastore; asynchronous
// Deletes all tuples for the refid from Hexastore, then creates tuples for the replacement object,
// and sends them to Hexastore as an ordered set of tuples
// Outputs SIF responses for each refid to errch
// Outputs a list of tuples to be sent to Hexastore to outch
func UpdateFullXMLasDBtriplesAsync(object []byte, context string, outch chan<- *Triple, errch chan<- SifResponse, ch chan<- struct{}) {
	haskeych := make(chan bool)
	m, err := mxj.NewMapXml(object)
	if err != nil {
		errch <- SifResponse{RefId: "", Error: err}
		ch <- struct{}{}
		return
	}
	advisoryIdRaw, err := m.ValueForPath("*.-RefId")
	var refid string
	if err != nil {
		errch <- SifResponse{RefId: "", Error: err}
		ch <- struct{}{}
		return
	} else {
		limitch <- struct{}{}
		refid = advisoryIdRaw.(string)
		go existsRefid(refid, context, haskeych)
	}
	triples, err := makeDeleteTriplesForRefId(refid, context)
	if err != nil {
		errch <- SifResponse{RefId: "", Error: err}
		ch <- struct{}{}
		return
	}
	for _, n := range m.LeafNodes() {
		triples = append(triples, &Triple{
			Subject:   strconv.Quote(refid),
			Predicate: n.Path,
			Object:    fmt.Sprintf("%v", n.Value),
			Context:   context})
	}
	refIdClash := <-haskeych
	if !refIdClash {
		errch <- SifResponse{RefId: refid, Error: fmt.Errorf("RefID %v not found\n", refid)}
	} else {
		for _, t := range triples {
			outch <- t
		}
		errch <- SifResponse{AdvisoryId: refid, RefId: refid, Error: nil}
	}
	ch <- struct{}{}

}

// Update an object, with Partial Replacement on the Hexastore; synchronous.
// Is a brute force patch; deletes then writes to all the subject-predicate instances it finds in the object.
// Does not otherwise delete anything, including extra list entries:
// e.g. if a list with two items is patched with a list of one item, it will overwrite the first item in the list,
// and will leave alone the second.
func UpdatePartialXMLasDBtriples(object []byte, requested_refid string, context string) error {
	m, err := mxj.NewMapXml(object)
	if err != nil {
		return err
	}
	refid, err := m.ValueForPath("*.-RefId")
	if err != nil {
		refid = requested_refid
	}
	all_triples := make([]*Triple, 0)
	// delete for only the subject-predicate pairs found in the patch object
	for _, n := range m.LeafNodes() {
		triples := getTuples(
			fmt.Sprintf("s:%s p:%s",
				strconv.Quote(fmt.Sprintf("%v", refid)), strconv.Quote(fmt.Sprintf("%v", n.Path))),
			context)
		for i := range triples {
			triples[i].Object = ""
		}
		all_triples = append(all_triples, triples...)
	}
	limitch <- struct{}{}
	send_triples(all_triples, context)
	all_triples = make([]*Triple, 0)
	for _, n := range m.LeafNodes() {
		all_triples = append(all_triples, &Triple{
			Subject:   fmt.Sprintf("%v", refid),
			Predicate: n.Path,
			Object:    fmt.Sprintf("%v", n.Value),
			Context:   context})
	}
	limitch <- struct{}{}
	send_triples(all_triples, context)
	return nil
}

// Update an object, with Partial Replacement on the Hexastore; asynchronous.
// Is a brute force patch; deletes then writes to all the subject-predicate instances it finds in the object.
// Does not otherwise delete anything, including extra list entries:
// e.g. if a list with two items is patched with a list of one item, it will overwrite the first item in the list,
// and will leave alone the second.
// Outputs SIF responses for each refid to errch
// Outputs a list of tuples to be sent to Hexastore to outch
func UpdatePartialXMLasDBtriplesAsync(object []byte, context string, outch chan<- *Triple, errch chan<- SifResponse, ch chan<- struct{}) {
	haskeych := make(chan bool)
	m, err := mxj.NewMapXml(object)
	if err != nil {
		errch <- SifResponse{RefId: "", Error: err}
		ch <- struct{}{}
		return
	}
	advisoryIdRaw, err := m.ValueForPath("*.-RefId")
	var refid string
	if err != nil {
		errch <- SifResponse{RefId: "", Error: err}
		ch <- struct{}{}
		return
	} else {
		limitch <- struct{}{}
		refid = advisoryIdRaw.(string)
		go existsRefid(refid, context, haskeych)
	}
	triples := make([]*Triple, 0)
	// delete for only the subject-predicate pairs found in the patch object
	for _, n := range m.LeafNodes() {
		triples = append(triples, &Triple{
			Subject:   strconv.Quote(refid),
			Predicate: n.Path,
			Object:    fmt.Sprintf("%v", n.Value),
			Context:   context})
	}
	refIdClash := <-haskeych
	if !refIdClash {
		errch <- SifResponse{RefId: refid, Error: fmt.Errorf("RefID %v not found\n", refid)}
	} else {
		for _, t := range triples {
			outch <- t
		}
		errch <- SifResponse{AdvisoryId: refid, RefId: refid, Error: nil}
	}
	ch <- struct{}{}

}

// Store the object on the Hexastore. If mustUseAdvisory is on, use the supplied RefId if available,
// else raise error; if off, generate new RefId. Synchronous. Return the RefId used, the triples
// to be posted, and the error.
func StoreXMLasDBtriples(object []byte, mustUseAdvisory bool, context string) (string, []*Triple, error) {
	ch := make(chan bool)

	m, err := mxj.NewMapXml(object)
	if err != nil {
		return "", nil, err
	}

	refid, err := m.ValueForPath("*.-RefId")
	if err != nil {
		refid = strings.ToUpper(uuid.NewV4().String())
		mustUseAdvisory = false
		m.SetValueForPath(refid, "*.-RefId")
	} else {
		if mustUseAdvisory {
			limitch <- struct{}{}
			go existsRefid(refid.(string), context, ch)
		} else {
			refid = strings.ToUpper(uuid.NewV4().String())
			m.SetValueForPath(refid, "*.-RefId")
		}
	}
	triples := make([]*Triple, 0)
	for _, n := range m.LeafNodes() {
		triples = append(triples, &Triple{
			Subject:   fmt.Sprintf("%v", refid),
			Predicate: n.Path,
			Object:    fmt.Sprintf("%v", n.Value),
			Context:   context})
	}
	if mustUseAdvisory {
		refIdClash := <-ch
		if refIdClash {
			return "", nil, fmt.Errorf("RefID %v already in use\n", refid.(string))
		}
	}
	return refid.(string), triples, nil
}

// Response to CRUD action, to be posted back as SIF response for a payload with multiple records.
type SifResponse struct {
	AdvisoryId string // The RefId in the request
	RefId      string // The RefId of the response object, which may be different
	Error      error
}

// Store an XML object on the Hexastore. If mustUseAdvisory is on, use the supplied RefId if available,
// else raise error; if off, generate new RefId. Asynchronous.
// Outputs SIF responses for each refid to errch
// Outputs a list of tuples to be sent to Hexastore to outch
func StoreXMLasDBtriplesAsync(s []byte, mustUseAdvisory bool, context string, outch chan<- *Triple, errch chan<- SifResponse, ch chan<- struct{}) {
	haskeych := make(chan bool)
	m, err := mxj.NewMapXml(s)
	if err != nil {
		errch <- SifResponse{RefId: "", Error: err}
		ch <- struct{}{}
		return
	}
	advisoryIdRaw, err := m.ValueForPath("*.-RefId")
	var refid, advisoryId string
	if err != nil {
		refid = strings.ToUpper(uuid.NewV4().String())
		advisoryId = ""
		mustUseAdvisory = false
	} else {
		advisoryId = advisoryIdRaw.(string)
		if mustUseAdvisory {
			refid = advisoryId
			limitch <- struct{}{}
			go existsRefid(refid, context, haskeych)
		} else {
			refid = strings.ToUpper(uuid.NewV4().String())
		}
	}
	m.SetValueForPath(refid, "*.-RefId")
	triples := make([]*Triple, 0)
	for _, n := range m.LeafNodes() {
		triples = append(triples, &Triple{
			Subject:   fmt.Sprintf("%v", refid),
			Predicate: n.Path,
			Object:    fmt.Sprintf("%v", n.Value),
			Context:   context})
	}
	refIdClash := false
	if mustUseAdvisory {
		refIdClash = <-haskeych
		if refIdClash {
			errch <- SifResponse{AdvisoryId: advisoryId, RefId: "",
				Error: fmt.Errorf("RefID %v already in use\n", refid)}
		}
	}
	if !refIdClash {
		for _, t := range triples {
			outch <- t
		}
		errch <- SifResponse{AdvisoryId: advisoryId, RefId: refid, Error: nil}
	}
	ch <- struct{}{}
}

// Store a JSON object on the Hexastore. Asynchronous.
// Outputs SIF responses for each refid to errch
// Outputs a list of tuples to be sent to Hexastore to outch
// JSON object is assumed to have an "id" top-level value, which is the subject for the tuples of that object.
// If not present, one will be created. If it is present, and it has already been used as an object in that
// context, the request will be denied.
func StoreJSONasDBtriplesAsync(s []byte, context string, outch chan<- *Triple, errch chan<- SifResponse, ch chan<- struct{}) {
	haskeych := make(chan bool)
	m, err := mxj.NewMapJson(s)
	if err != nil {
		log.Println(err)
		errch <- SifResponse{RefId: "", Error: err}
		ch <- struct{}{}
		return
	}
	advisoryIdRaw, err := m.ValueForPath("id")
	var refid, advisoryId string
	checkRefId := false
	if err != nil {
		advisoryId = ""
		refid = strings.ToUpper(uuid.NewV4().String())
	} else {
		advisoryId = advisoryIdRaw.(string)
		refid = strings.ToUpper(advisoryId)
		checkRefId = true
		limitch <- struct{}{}
		go existsRefid(refid, context, haskeych)
	}
	m.SetValueForPath(refid, "id")
	triples := make([]*Triple, 0)
	for _, n := range m.LeafNodes() {
		triples = append(triples, &Triple{
			Subject:   fmt.Sprintf("%v", refid),
			Predicate: n.Path,
			Object:    fmt.Sprintf("%v", n.Value),
			Context:   context})
	}
	refIdClash := false
	if checkRefId {
		refIdClash = <-haskeych
	}
	if refIdClash {
		errch <- SifResponse{AdvisoryId: advisoryId, RefId: "",
			Error: fmt.Errorf("ID %v already in use\n", refid)}
	} else {
		for _, t := range triples {
			//log.Printf("%#v\n", t)
			outch <- t
		}
		errch <- SifResponse{AdvisoryId: advisoryId, RefId: refid, Error: nil}
	}
	ch <- struct{}{}
}

var mxj2sjsonPathRe1 = regexp.MustCompile(`\[(\d+)\]`)
var mxj2sjsonPathRe2 = regexp.MustCompile(`\.#text$`)

// Convert an MXJ path to a node to an sjson path, using dot notation instead of array, and ".Value" instead of .#text
func mxj2sjsonPath(p string) string {
	return mxj2sjsonPathRe1.ReplaceAllString(
		mxj2sjsonPathRe2.ReplaceAllString(p, ".Value"), ".$1")
}

// TODO: no flow control yet: may retrieve a massive number of RefIds
// Retrieve all RefIDs from Hexastore for the given object class. Synchronous
func GetAllXMLByObject(object string, context string) ([]string, error) {
	triples := getTuples(fmt.Sprintf("p:%s s:", strconv.Quote(fmt.Sprintf("%s.-RefId", object))), context)
	objIDs := make([]string, 0)
	for _, t := range triples {
		objIDs = append(objIDs, t.Subject)
	}
	return objIDs, nil
}

// Retrieve the XML object corresponding to the tuples stored under the given RefId. Synchronous.
// If stripempty, strip empty tags from the XML object.
func DbTriples2XML(refid string, context string, stripempty bool) ([]byte, error) {
	//triples := getTuples(fmt.Sprintf("s:%s p:", strconv.Quote(fmt.Sprintf("%v", refid))), context)
	triples := getTuplesInflux(refid, context)
	json := ""
	var err error
	for _, t := range triples {
		// Build up the XML object through sjson
		json, err = sjson.Set(json, mxj2sjsonPath(t.Predicate), t.Object)
		if err != nil {
			return nil, err
		}
	}
	mm, err := mxj.NewMapJson([]byte(json))
	if err != nil {
		return nil, err
	}
	return Map2SIFXML(mm, stripempty)
}

var emptytag1 = regexp.MustCompile(`(?s:\s*<[^>/]+></[^>]+>\s*)`)
var emptytag2 = regexp.MustCompile(`(?s:\s*<[^>/]+/>\s*)`)
var emptytag3 = regexp.MustCompile(`(?s:\s+[^>='" ]+=(''|""))`)

// Brute force stripping of empty tags and attributes from XML string.
// Works for SIF because it does not have mixed tags and text.
// Optional to call; will have performance hit
func stripEmptyTags(s []byte) []byte {
	s = emptytag1.ReplaceAll(s, []byte(""))
	s = emptytag1.ReplaceAll(s, []byte(""))
	s = emptytag1.ReplaceAll(s, []byte(""))
	s = emptytag1.ReplaceAll(s, []byte(""))
	s = emptytag1.ReplaceAll(s, []byte(""))
	s = emptytag2.ReplaceAll(s, []byte(""))
	arr := strings.SplitAfter(string(s), ">")
	for i, _ := range arr {
		if strings.HasPrefix(arr[i], "<") {
			// only works because we don't have mixed tags and text in SIF
			arr[i] = emptytag3.ReplaceAllString(arr[i], "")
		}
	}
	s = []byte(strings.Join(arr, ""))
	return s
}

var remove_empty_structs = regexp.MustCompile(`,"[^"]+":\{\}|,"[^"]+":\{"[^"]+":\{\}\}`)
var remove_empty_structs1 = regexp.MustCompile(`\{"[^"]+":\{\},`)
var attribute_rename = regexp.MustCompile(`([^\\])"-([^"\\]+)":`)
var value_rename = regexp.MustCompile(`([^\\])"Value":`)

// Convert an XML object to JSON using Goessner notation.
// Ignores namespaces, and gets rid of xs:nil
func XMLtoJSON(x []byte) ([]byte, error) {
	m, err := mxj.NewMapXml(x)
	json, err := m.Json(true)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	for _, n := range m.LeafNodes() {
		if len(n.Value.(string)) == 0 || strings.HasSuffix(n.Path, ".-xmlns") ||
			strings.HasSuffix(n.Path, ".-xsi") || strings.HasSuffix(n.Path, ".-xsd") {
			json, err = sjson.DeleteBytes(json, mxj2sjsonPath(n.Path))
			if err != nil {
				log.Println(err)
				return nil, err
			}
		} else if strings.HasSuffix(n.Path, ".-nil") {
			json, err = sjson.DeleteBytes(json, strings.TrimSuffix(mxj2sjsonPath(n.Path), ".-nil"))
			if err != nil {
				log.Println(err)
				return nil, err
			}
		}
	}
	// and some brute force
	json = remove_empty_structs.ReplaceAll(json, []byte(""))
	json = remove_empty_structs.ReplaceAll(json, []byte(""))
	json = remove_empty_structs.ReplaceAll(json, []byte(""))
	json = remove_empty_structs1.ReplaceAll(json, []byte("{"))
	json = attribute_rename.ReplaceAll(json, []byte("$1\"@$2\":"))
	json = value_rename.ReplaceAll(json, []byte("$1\"#text\":"))
	return json, nil
}

// get tuples, and organise them into JSON objects by subject
func Tuples2JSON(triples []*Triple) (string, error) {
	json := make(map[string]string)
	var err error
	for _, t := range triples {
		// Build up the JSON object through sjson
		if _, ok := json[t.Subject]; !ok {
			json[t.Subject] = ""
		}
		json[t.Subject], err = sjson.Set(json[t.Subject], mxj2sjsonPath(t.Predicate), t.Object)
		if err != nil {
			return "", err
		}
	}
	ret := "["
	for i, j := range json {
		ret += j
		if i < len(json)-1 {
			ret += ",\n"
		}
	}
	ret += "]"
	return ret, nil
}
