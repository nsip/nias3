package main

import (
	"github.com/nsip/nias3/mxj"
)

func main() {
	refid, _ := mxj.storeXMLasDBtriples(staff_personal)
	x, _ := mxj.dbTriples2XML(refid)
	log.Printf("Map2SIFXML\n%+v\n", string(x))
}
