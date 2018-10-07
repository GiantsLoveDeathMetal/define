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
		definition, err := request.Define(word)
		if err != nil {
			correction := request.ReadCorrection("mock_correction")
			fmt.Println(correction.Suggestion)
			definition, _ := request.ReadDefinition("mock_definition")
			fmt.Println(definition.Subsenses)
		}
		for _, s := range definition.Subsenses {
			fmt.Println(s)
		}

		// correction := request.Correct(word)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
