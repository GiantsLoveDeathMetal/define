package request

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Definition struct {
	Results []Result `json:"results"`
}

type Result struct {
	LexicalEntries []LexicalEntry `json:"lexicalEntries"`
}

type LexicalEntry struct {
	Entries []Entry `json:"entries"`
}

type Entry struct {
	Senses []Sense `json:"senses"`
}

type Sense struct {
	Subsenses []Subsense `json:"subsenses"`
}

type Subsense struct {
	Shortdefinitions []string `json:"short_definitions"`
	Domains          []string `json:"domains"`
}

var LANGUAGE string = "en"

func Define(w string) (*Sense, error) {
	var URL string = "https://od-api.oxforddictionaries.com:443/api/v1/entries/%v/%v"

	client := &http.Client{}
	dic_key := os.Getenv("DICTIONARY_APP_KEY")
	dic_id := os.Getenv("DICTIONARY_APP_ID")

	url := fmt.Sprintf(URL, LANGUAGE, w)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("app_key", dic_key)
	req.Header.Set("app_id", dic_id)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return parseDefinition(body)
}

func parseDefinition(d []byte) (*Sense, error) {
	var def Definition
	err := json.Unmarshal(d, &def)
	if err != nil {
		return nil, err
	}
	return &def.Results[0].LexicalEntries[0].Entries[0].Senses[0], nil
}
