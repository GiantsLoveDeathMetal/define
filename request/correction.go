package request

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func Correct(w string) string {
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
	return string(body)
}
