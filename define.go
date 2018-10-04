package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/urfave/cli"
)

var LANGUAGE string = "en"
var URL string = "https://od-api.oxforddictionaries.com:443/api/v1/entries/%v/%v"

func main() {
	client := &http.Client{}

	dic_key := os.Getenv("DICTIONARY_APP_KEY")
	dic_id := os.Getenv("DICTIONARY_APP_ID")

	app := cli.NewApp()
	app.Name = "define"
	app.Usage = "Define a word in the command line!"
	app.Action = func(c *cli.Context) error {
		url := fmt.Sprintf(URL, LANGUAGE, c.Args().Get(0))
		req, err := http.NewRequest("GET", url, nil)
		req.Header.Set("app_key", dic_key)
		req.Header.Set("app_id", dic_id)
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)

		fmt.Println(req)
		fmt.Println(string(body))
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
