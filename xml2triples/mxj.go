package xml2triples

import (
	"bytes"
	"encoding/json"
	"fmt"
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

var client = &http.Client{}
var conf = config.LoadConfig()
var baseUrl = fmt.Sprintf("http://localhost:%d", conf.N3EngineWebport)

func changeJSONTags(j string) string {
	return strings.Replace(j, `"#text":`, `"Value":`, -1)
}

func Map2SIFXML(m mxj.Map, stripempty bool) ([]byte, error) {
	root, err := m.Root()
	if err != nil {
		return nil, err
	}
	//log.Println(root)
	m02 := m[root].(map[string]interface{})
	// log.Printf("m02\n%+v\n", m02)
	j, err := json.Marshal(m02)
	if err != nil {
		return nil, err
	}
	// log.Println(string(j))
	ret, err := Root2SIF(root, []byte(changeJSONTags(string(j))))
	if err != nil {
		return nil, err
	}
	if stripempty {
		ret = stripEmptyTags(ret)
	}
	return ret, nil
}

func send_triple(triple Triple) {
	json, err := json.Marshal(triple)
	if err != nil {
		panic(err)
	}
	//log.Println(string(json))
	req, err := http.NewRequest("POST", baseUrl+"/tuple", bytes.NewBuffer(json))
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

func Send_triples(triples []*Triple) {
	json, err := json.Marshal(triples)
	if err != nil {
		panic(err)
	}
	//log.Println(string(json))
	req, err := http.NewRequest("POST", baseUrl+"/tuples", bytes.NewBuffer(json))
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

// check if key prefix is on Hexastore
// TODO restrict query to a context
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
	//log.Printf("hasKey status: %d\n", resp.StatusCode)
	return resp.StatusCode == 200
}

// retrieve tuples from Hexastore matching a key prefix (involving a subset of s: o: p:; the c: prefix
// will be added here)
// TODO restrict query to a context
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

/* NSW DIG hard coded queries */
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

// TODO restrict query to a context
func DeleteTriplesForRefId(refid string, context string) error {
	//log.Println("DeleteTriplesForRefId")
	triples := getTuples(fmt.Sprintf("s:%s ", strconv.Quote(fmt.Sprintf("%v", refid))), context)
	//log.Printf("%+v\n", triples)
	for i := range triples {
		triples[i].Object = ""
	}
	Send_triples(triples)
	return nil
}

func UpdateFullXMLasDBtriples(s []byte, refid string, context string) error {
	m, err := mxj.NewMapXml(s)
	if err != nil {
		return err
	}
	err = DeleteTriplesForRefId(refid, context)
	triples := make([]*Triple, 0)
	for _, n := range m.LeafNodes() {
		triples = append(triples, &Triple{Subject: fmt.Sprintf("%v", refid), Predicate: n.Path, Object: fmt.Sprintf("%v", n.Value), Context: context})
	}
	Send_triples(triples)
	return nil
}

// does not delete anything, including extra list entries: will not shrink list of 2 to list of 1
func UpdatePartialXMLasDBtriples(s []byte, refid string, context string) error {
	m, err := mxj.NewMapXml(s)
	if err != nil {
		return err
	}
	all_triples := make([]*Triple, 0)
	for _, n := range m.LeafNodes() {
		triples := getTuples(fmt.Sprintf("s:%s p:%s", strconv.Quote(fmt.Sprintf("%v", refid)), strconv.Quote(fmt.Sprintf("%v", n.Path))), context)
		for i := range triples {
			triples[i].Object = ""
		}
		all_triples = append(all_triples, triples...)
	}
	Send_triples(all_triples)
	all_triples = make([]*Triple, 0)
	for _, n := range m.LeafNodes() {
		all_triples = append(all_triples, &Triple{Subject: fmt.Sprintf("%v", refid), Predicate: n.Path, Object: fmt.Sprintf("%v", n.Value), Context: context})
	}
	Send_triples(all_triples)
	return nil
}

type Triple struct {
	Subject   string
	Object    string
	Predicate string
	Context   string
}

// nominated refid overrides any refid in the object
func StoreXMLasDBtriples(s []byte, mustUseAdvisory bool, context string) (string, []*Triple, error) {
	m, err := mxj.NewMapXml(s)
	if err != nil {
		return "", nil, err
	}
	//log.Printf("mustUseAdvisory %v\n", mustUseAdvisory)
	refid, err := m.ValueForPath("*.-RefId")
	if err != nil {
		refid = strings.ToUpper(uuid.NewV4().String())
	} else {
		if mustUseAdvisory && hasKey(fmt.Sprintf("s:%s p:", strconv.Quote(fmt.Sprintf("%v", refid))), context) {
			return "", nil, fmt.Errorf("RefID %v already in use\n", refid.(string))
		} else {
			refid = strings.ToUpper(uuid.NewV4().String())
		}
	}
	m.SetValueForPath(refid, "*.-RefId")
	triples := make([]*Triple, 0)
	for _, n := range m.LeafNodes() {
		triples = append(triples, &Triple{Subject: fmt.Sprintf("%v", refid), Predicate: n.Path, Object: fmt.Sprintf("%v", n.Value), Context: context})
	}
	//Send_triples(triples)
	return refid.(string), triples, nil
}

var mxj2sjsonPathRe1 = regexp.MustCompile(`\[(\d+)\]`)
var mxj2sjsonPathRe2 = regexp.MustCompile(`\.#text$`)

func mxj2sjsonPath(p string) string {
	return mxj2sjsonPathRe1.ReplaceAllString(
		mxj2sjsonPathRe2.ReplaceAllString(p, ".Value"), ".$1")
}

// no flow control yet
func GetAllXMLByObject(object string, context string) ([]string, error) {
	triples := getTuples(fmt.Sprintf("p:%s s:", strconv.Quote(fmt.Sprintf("%s.-RefId", object))), context)
	//log.Println("%+v\n", triples)
	objIDs := make([]string, 0)
	for _, t := range triples {
		//objIDs = append(objIDs, t.S)
		objIDs = append(objIDs, t.Subject)
	}
	return objIDs, nil
}

func DbTriples2XML(refid string, context string, stripempty bool) ([]byte, error) {
	triples := getTuples(fmt.Sprintf("s:%s p:", strconv.Quote(fmt.Sprintf("%v", refid))), context)

	json := ""
	var err error
	for _, t := range triples {
		//log.Printf("%s %s %s\n", t.S, t.P, t.O)
		//log.Printf("%s %s %s\n", t.S, mxj2sjsonPath(t.P), t.O)
		json, err = sjson.Set(json, mxj2sjsonPath(t.Predicate), t.Object)
		if err != nil {
			return nil, err
		}
	}
	//log.Printf("%+v\n", json)
	mm, err := mxj.NewMapJson([]byte(json))
	if err != nil {
		return nil, err
	}
	// log.Printf("%+v\n", mm)
	return Map2SIFXML(mm, stripempty)
}

// Brute force stripping of empty tags and attributes from XML string.
// Works for SIF because it does not have mixed tags and text.
var emptytag1 = regexp.MustCompile(`(?s:\s*<[^>/]+></[^>]+>\s*)`)
var emptytag2 = regexp.MustCompile(`(?s:\s*<[^>/]+/>\s*)`)
var emptytag3 = regexp.MustCompile(`(?s:\s+[^>='" ]+=(''|""))`)

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

// convert to Goessner notation
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
