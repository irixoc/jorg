// Package jorg is used indent json files for readability
package jorg

import (
	"bytes"
	"encoding/json"
	"os"
)

// IndentJSON reads in a byteslice of a json object and prints it to file
func IndentJSON(byteSlice []byte) error {

	var buffer bytes.Buffer
	json.Indent(&buffer, byteSlice, "", "    ")

	file, err := os.Create("indented.json")
	if err != nil {
		return err
	}

	_, err = file.Write(buffer.Bytes())
	if err != nil {
		file.Close()
		return err
	}

	if err := file.Close(); err != nil {
		return err
	}

	return nil
}
