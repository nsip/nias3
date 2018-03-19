package sifxml

import (
	"regexp"
)

var re = regexp.MustCompile(`^(\s*\{"[^"]+"\s*:)`)
var re2 = regexp.MustCompile(`\}\s*$`)
