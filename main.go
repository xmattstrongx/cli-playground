package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	playground := cli.NewApp()
	playground.Name = " cli-playground"
	playground.Usage = "playing around with urfave cli"

	var lang string

	playground.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "language, l",
			Value:       "english",
			Usage:       "language for the greeting",
			Destination: &lang,
		},
		cli.StringFlag{
			Name:  "config, c",
			Usage: "Load configuration from `FILE`",
		},
	}

	playground.Action = func(c *cli.Context) error {
		name := "dude"
		if c.NArg() > 0 {
			name = c.Args().Get(0)
		}
		if lang == "spanish" {
			fmt.Println("Hola", name)
		} else {
			fmt.Println("Hello", name)
		}
		return nil
	}

	playground.Run(os.Args)

}