package request

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var LANGUAGE string = "en"

func Define(w string) string {
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
	return string(body)
}
