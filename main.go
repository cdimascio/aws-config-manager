package main

import (
	"github.com/cdimascio/aws-config-manager/cmd"
	"log"
	"os"
	"sort"

	"github.com/urfave/cli/v2"
)

const ConfigDir="$HOME/.aws_cred_man"

func main() {
	app := &cli.App{
		//Flags: []cli.Flag{
		//	&cli.StringFlag{
		//		Name:    "lang",
		//		Aliases: []string{"l"},
		//		Value:   "english",
		//		Usage:   "Language for the greeting",
		//	},
		//	&cli.StringFlag{
		//		Name:    "config",
		//		Aliases: []string{"c"},
		//		Usage:   "Load configuration from `FILE`",
		//	},
		//},
		Commands: []*cli.Command{
			{
				Name:  "use",
				Usage: "use <setting>",
				Action: func(c *cli.Context) error {
					cmd.Initialize()
					return cmd.Use(c.Args())
				},
			},
			{
				Name:  "current",
				Usage: "current",
				Aliases: []string{"cur"},
				Action: func(c *cli.Context) error {
					cmd.Initialize()
					return cmd.Current()
				},
			},
			{
				Name:    "list",
				Aliases: []string{"ls"},
				Usage:   "add a task to the list",
				Action: func(c *cli.Context) error {
					cmd.Initialize()
					return cmd.List()
				},
			},
			{
				Name:  "create",
				Usage: "create",
				Action: func(c *cli.Context) error {
					cmd.Initialize()
					return cmd.Create(c.Args())
				},
			},
			{
				Name:  "edit",
				Usage: "edit",
				Action: func(c *cli.Context) error {
					cmd.Initialize()
					return cmd.Edit(c.Args())
				},
			},
			{
				Name:  "remove",
				Usage: "remove",
				Aliases: []string{"rm"},
				Action: func(c *cli.Context) error {
					cmd.Initialize()
					return cmd.Remove(c.Args())
				},
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
