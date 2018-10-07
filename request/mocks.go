package request

import (
	"fmt"
	"io/ioutil"
)

// Read JSON into response
func ReadJson(w string) []byte {
	filename := fmt.Sprintf("request/%v.json", w)

	f, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return f
}

func ReadDefinition(f string) (*Sense, error) {
	return parseDefinition(ReadJson(f))
}

func ReadCorrection(f string) Correction {
	return parseCorrection(ReadJson(f))
}
