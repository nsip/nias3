package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"text/scanner"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/syndtr/goleveldb/leveldb/util"
)

var db *leveldb.DB
var dbOpen bool = false

//var ge = GobEncoder{}

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

func dataDump() {
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
func getIdentifiers(keyPrefix string) []string {

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
	s string
	p string
	o string
}

func parseTriples(list []string) []Triple {
	ret := make([]Triple, 0)
	for _, t := range list {
		var s scanner.Scanner
		s.Init(strings.NewReader(t))
		o := Triple{}
		for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
			if s.TokenText() == "s" {
				tok = s.Scan() // colon
				tok = s.Scan()
				o.s, _ = strconv.Unquote(s.TokenText())
			}
			if s.TokenText() == "p" {
				tok = s.Scan() // colon
				tok = s.Scan()
				o.p, _ = strconv.Unquote(s.TokenText())
			}
			if s.TokenText() == "o" {
				tok = s.Scan() // colon
				tok = s.Scan()
				o.o, _ = strconv.Unquote(s.TokenText())
			}
		}
		ret = append(ret, o)
	}
	return ret
}
