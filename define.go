package main

import (
	"fmt"
	"log"
	"os"

	"github.com/foxyblue/define/request"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "define"
	app.Usage = "Define a word in the command line!"
	app.Action = func(c *cli.Context) error {
		word := c.Args().Get(0)
		definition := request.Define(word)
		fmt.Println(definition)

		correction := request.Correct(word)
		fmt.Println(correction)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
