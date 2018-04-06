package xml2triples

import (
	"encoding/json"
	"fmt"
	"github.com/clbanning/mxj"
	"github.com/nsip/nias3/datastore"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/tidwall/sjson"
	"github.com/twinj/uuid"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func changeJSONTags(j string) string {
	return strings.Replace(j, `"#text":`, `"Value":`, -1)
}

func Map2SIFXML(m mxj.Map) ([]byte, error) {
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
	return Root2SIF(root, []byte(changeJSONTags(string(j))))
}

func put_triple(batch *leveldb.Batch, triple datastore.Triple) {
	log.Printf("s:%s p:%s o:%s", strconv.Quote(triple.S), strconv.Quote(triple.P), strconv.Quote(triple.O))
	keys := datastore.PermuteTriple(triple)
	for _, key := range keys {
		batch.Put([]byte(key), []byte(keys[0]))
	}
}

func DeleteTriplesForRefId(refid string) error {
	db := datastore.GetDB()
	batch := new(leveldb.Batch)
	triple_strings := datastore.GetIdentifiers(fmt.Sprintf("s:%s ", strconv.Quote(fmt.Sprintf("%v", refid))))
	all_keys := datastore.PermuteTripleKeys(triple_strings)
	for _, key := range all_keys {
		batch.Delete([]byte(key))
	}
	batcherr := db.Write(batch, nil)
	if batcherr != nil {
		return batcherr
	}
	batch.Reset()
	return nil
}

func UpdateFullXMLasDBtriples(s []byte, refid string) error {
	m, err := mxj.NewMapXml(s)
	if err != nil {
		return err
	}
	db := datastore.GetDB()
	batch := new(leveldb.Batch)
	/*
		refid, err := m.ValueForPath("*.-RefId")
		if err != nil {
			return "", err
		}
	*/
	err = DeleteTriplesForRefId(refid)
	for _, n := range m.LeafNodes() {
		put_triple(batch, datastore.Triple{S: fmt.Sprintf("%v", refid), P: n.Path, O: fmt.Sprintf("%v", n.Value)})
	}
	batcherr := db.Write(batch, nil)
	if batcherr != nil {
		return err
	}
	batch.Reset()
	return nil
}

// does not delete anything, including extra list entries: will not shrink list of 2 to list of 1
func UpdatePartialXMLasDBtriples(s []byte, refid string) error {
	m, err := mxj.NewMapXml(s)
	if err != nil {
		return err
	}
	db := datastore.GetDB()
	batch := new(leveldb.Batch)
	/*
		refid, err := m.ValueForPath("*.-RefId")
		if err != nil {
			return "", err
		}
	*/
	for _, n := range m.LeafNodes() {
		triple_strings := datastore.GetIdentifiers(fmt.Sprintf("s:%s p:%s", strconv.Quote(fmt.Sprintf("%v", refid)), strconv.Quote(fmt.Sprintf("%v", n.Path))))
		all_keys := datastore.PermuteTripleKeys(triple_strings)
		for _, key := range all_keys {
			batch.Delete([]byte(key))
		}
	}
	batcherr := db.Write(batch, nil)
	if batcherr != nil {
		return err
	}
	batch.Reset()
	for _, n := range m.LeafNodes() {
		put_triple(batch, datastore.Triple{S: fmt.Sprintf("%v", refid), P: n.Path, O: fmt.Sprintf("%v", n.Value)})
	}
	batcherr = db.Write(batch, nil)
	if batcherr != nil {
		return err
	}
	batch.Reset()
	return nil
}

// nominated refid overrides any refid in the object
func StoreXMLasDBtriples(s []byte, mustUseAdvisory bool) (string, error) {
	db := datastore.GetDB()
	batch := new(leveldb.Batch)
	m, err := mxj.NewMapXml(s)
	if err != nil {
		return "", err
	}
	log.Printf("mustUseAdvisory %v\n", mustUseAdvisory)
	refid, err := m.ValueForPath("*.-RefId")
	if err != nil {
		refid = strings.ToUpper(uuid.NewV4().String())
	} else {
		triple_strings := datastore.GetIdentifiers(fmt.Sprintf("s:%s p:", strconv.Quote(fmt.Sprintf("%v", refid))))
		if len(triple_strings) > 0 {
			if mustUseAdvisory {
				return "", fmt.Errorf("RefID %v already in use\n", refid.(string))
			} else {
				refid = strings.ToUpper(uuid.NewV4().String())
			}
		}
	}
	m.SetValueForPath(refid, "*.-RefId")
	for _, n := range m.LeafNodes() {
		put_triple(batch, datastore.Triple{S: fmt.Sprintf("%v", refid), P: n.Path, O: fmt.Sprintf("%v", n.Value)})
	}
	batcherr := db.Write(batch, nil)
	if batcherr != nil {
		return "", err
	}
	batch.Reset()
	return refid.(string), nil
}

var mxj2sjsonPathRe1 = regexp.MustCompile(`\[(\d+)\]`)
var mxj2sjsonPathRe2 = regexp.MustCompile(`\.#text$`)

func mxj2sjsonPath(p string) string {
	return mxj2sjsonPathRe1.ReplaceAllString(
		mxj2sjsonPathRe2.ReplaceAllString(p, ".Value"), ".$1")
}

// no flow control yet
func GetAllXMLByObject(object string) ([]string, error) {
	triple_strings := datastore.GetIdentifiers(fmt.Sprintf("p:%s s:", strconv.Quote(fmt.Sprintf("%s.-RefId", object))))
	triples := datastore.ParseTriples(triple_strings)
	objIDs := make([]string, 0)
	for _, t := range triples {
		objIDs = append(objIDs, t.S)
	}
	return objIDs, nil
}

func DbTriples2XML(refid string) ([]byte, error) {
	triple_strings := datastore.GetIdentifiers(fmt.Sprintf("s:%s p:", strconv.Quote(fmt.Sprintf("%v", refid))))
	triples := datastore.ParseTriples(triple_strings)
	json := ""
	var err error
	for _, t := range triples {
		//log.Printf("%s %s %s\n", t.S, t.P, t.O)
		//log.Printf("%s %s %s\n", t.S, mxj2sjsonPath(t.P), t.O)
		json, err = sjson.Set(json, mxj2sjsonPath(t.P), t.O)
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
	return Map2SIFXML(mm)
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
