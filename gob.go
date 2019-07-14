package jb

import (
	"bytes"
	"encoding/gob"
)

// GobData encodes the given object and returns gob encoded bytes.
func GobData(v interface{}) ([]byte, error) {
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(v)
	return buffer.Bytes(), err
}

// UnGobData decodes the given data into the given object.
func UnGobData(data []byte, v interface{}) error {
	buffer := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buffer)
	err := dec.Decode(v)
	return err
}
