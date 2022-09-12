package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "hsm",
		Usage: "Hardware Server Manager",
		Commands: []*cli.Command{
			{
				Name:    "wake",
				Aliases: []string{"w"},
				Usage:   "wake [target]",
				Action:  wake,
			},
			{
				Name:    "run",
				Aliases: []string{"s"},
				Usage:   "run [target] [command]",
				Action:  run,
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "file",
				Value:   "serverlist.json",
				Aliases: []string{"f"},
				Usage:   "Specify server list config file",
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
