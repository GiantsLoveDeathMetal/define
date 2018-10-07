package request

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Correction struct {
	Suggestion string `json:"suggestion"`
}

func Correct(w string) Correction {
	var URL string = "https://montanaflynn-spellcheck.p.mashape.com/check/?text=%v"

	client := &http.Client{}
	spell_key := os.Getenv("SPELL_CHECK_KEY")
	url := fmt.Sprintf(URL, w)

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("X-Mashape-Key", spell_key)
	req.Header.Set("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return parseCorrection(body)
}

func parseCorrection(d []byte) Correction {
	var def Correction
	err := json.Unmarshal(d, &def)
	if err != nil {
		panic(err)
	}
	return def
}
