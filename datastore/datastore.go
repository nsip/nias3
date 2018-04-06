package datastore

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/syndtr/goleveldb/leveldb/util"
	"log"
	"strconv"
	"strings"
	"text/scanner"
)

var db *leveldb.DB
var dbOpen bool = false

// if allow_empty is false, abort if database is empty: enforces
// requirement to run ingest first
func GetDB() *leveldb.DB {
	if !dbOpen {
		log.Println("DB not initialised. Opening...")
		openDB()
	}
	return db
}

//
// Check if the databse contains any entries
//
func DatabaseIsEmpty() bool {

	iter := GetDB().NewIterator(nil, nil)
	iter.Next()
	if len(iter.Key()) == 0 {
		return true
	}
	return false
}

//
// run compaction on db, typically post-ingest
//
func CompactDatastore() {
	err := GetDB().CompactRange(util.Range{Limit: nil, Start: nil})
	if err != nil {
		log.Println("Error compacting datastore: ", err)
	}
}

//
// open the kv store, this must be called before any access is attempted
//
func openDB() {

	workingDir := "kvs"

	o := &opt.Options{
		Filter: filter.NewBloomFilter(10),
		// BlockCacheCapacity: 128 * 1024 * 1024,
		// NoSync: true,
		// OpenFilesCacheCapacity: 1024,
		CompactionTableSize: (4 * opt.MiB),
	}
	var dbErr error
	db, dbErr = leveldb.OpenFile(workingDir, o)
	if dbErr != nil {
		log.Fatalln("DB Create error: ", dbErr)
	}

	dbOpen = true
}

func DataDump() {
	iter := GetDB().NewIterator(nil, nil)
	for iter.Next() {
		log.Printf("RETRIEVED KEY: %+v VALUE: %+v\n", string(iter.Key()), string(iter.Value()))
	}
	iter.Release()
	err := iter.Error()
	if err != nil {
		log.Println("Iterator error: ", err)
	}

}

//
// Given a key-prefix, returns the reference ids that
// can be used in a Get operation to retreive the
// desired object
//
func GetIdentifiers(keyPrefix string) []string {

	db = GetDB()
	objIDs := make([]string, 0)

	searchKey := []byte(keyPrefix)
	// log.Printf("search_key: %s\n\n", searchKey)
	iter := db.NewIterator(util.BytesPrefix(searchKey), nil)
	for iter.Next() {
		id := fmt.Sprintf("%s", iter.Value())
		objIDs = append(objIDs, id)
		// break
	}
	iter.Release()
	err := iter.Error()
	if err != nil {
		log.Println("Iterator error: ", err)
	}

	return objIDs
}

type Triple struct {
	S string
	P string
	O string
}

func ParseTriples(list []string) []Triple {
	ret := make([]Triple, 0)
	for _, t := range list {
		var s scanner.Scanner
		s.Init(strings.NewReader(t))
		o := Triple{}
		for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
			if s.TokenText() == "s" {
				tok = s.Scan() // colon
				tok = s.Scan()
				o.S, _ = strconv.Unquote(s.TokenText())
			}
			if s.TokenText() == "p" {
				tok = s.Scan() // colon
				tok = s.Scan()
				o.P, _ = strconv.Unquote(s.TokenText())
			}
			if s.TokenText() == "o" {
				tok = s.Scan() // colon
				tok = s.Scan()
				o.O, _ = strconv.Unquote(s.TokenText())
			}
		}
		ret = append(ret, o)
	}
	return ret
}

// first permutation is always SPO
func PermuteTriple(triple Triple) []string {
	ret := make([]string, 0)
	ret = append(ret, fmt.Sprintf("s:%s p:%s o:%s", strconv.Quote(triple.S), strconv.Quote(triple.P), strconv.Quote(triple.O)))
	ret = append(ret, fmt.Sprintf("s:%s o:%s p:%s", strconv.Quote(triple.S), strconv.Quote(triple.O), strconv.Quote(triple.P)))
	ret = append(ret, fmt.Sprintf("p:%s s:%s o:%s", strconv.Quote(triple.P), strconv.Quote(triple.S), strconv.Quote(triple.O)))
	ret = append(ret, fmt.Sprintf("p:%s o:%s s:%s", strconv.Quote(triple.P), strconv.Quote(triple.O), strconv.Quote(triple.S)))
	ret = append(ret, fmt.Sprintf("o:%s p:%s s:%s", strconv.Quote(triple.O), strconv.Quote(triple.P), strconv.Quote(triple.S)))
	ret = append(ret, fmt.Sprintf("o:%s s:%s p:%s", strconv.Quote(triple.O), strconv.Quote(triple.S), strconv.Quote(triple.P)))
	return ret
}

func PermuteTripleKeys(list []string) []string {
	ret := make([]string, 0)
	triples := ParseTriples(list)
	for _, triple := range triples {
		keys := PermuteTriple(triple)
		for _, key := range keys {
			ret = append(ret, key)
		}
	}
	return ret
}
