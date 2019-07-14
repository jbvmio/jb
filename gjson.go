package jb

import (
	"github.com/tidwall/gjson"
	"github.com/tidwall/pretty"
)

// GJsonKeepRaw .
func GJsonKeepRaw(json []byte, path string) []byte {
	var raw []byte
	r := gjson.GetBytes(json, path)
	if r.Index > 0 {
		raw = json[r.Index : r.Index+len(r.Raw)]
	} else {
		raw = []byte(r.Raw)
	}
	return raw
}

// JSONPretty .
func JSONPretty(json []byte) []byte {
	return pretty.Pretty(json)
}
