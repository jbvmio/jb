package jb

import "log"

// H here
func H(desc string, err error) {
	if err != nil {
		log.Fatalf("%v: %v\n", desc, err)
	}
}
