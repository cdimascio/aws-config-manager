package main

import (
	"github.com/cdimascio/aws-config-manager/cmd"
	"github.com/cdimascio/aws-config-manager/cmd/color"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/urfave/cli/v2"
)

var re = regexp.MustCompile(`(.*)\[command options\](.*)`)

func main() {
	template := color.ColorBlue+strings.Replace(cli.CommandHelpTemplate, "[arguments...]", "<setting>", -1)
	templateNoOpts := color.ColorBlue+strings.Replace(re.ReplaceAllString(cli.CommandHelpTemplate, "$1$2"), "[arguments...]", "<setting>", -1)
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:               "use",
				Usage:              "sets the current setting",
				CustomHelpTemplate: templateNoOpts,
				Action: func(c *cli.Context) error {
					cmd.Initialize()

					if c.Args().Len() < 1 {
						return cli.ShowCommandHelp(c, "use")
					}

					return cmd.Use(c.Args())
				},
			},
			{
				Name:               "current",
				Usage:              "shows the current setting",
				Aliases:            []string{"cur"},
				CustomHelpTemplate: templateNoOpts,
				Action: func(c *cli.Context) error {
					cmd.Initialize()
					return cmd.Current()
				},
			},
			{
				Name:               "list",
				Aliases:            []string{"ls"},
				Usage:              "list all settings",
				CustomHelpTemplate: templateNoOpts,
				Action: func(c *cli.Context) error {
					cmd.Initialize()
					return cmd.List()
				},
			},
			{
				Name:               "create",
				Usage:              "creates a new empty setting.",
				CustomHelpTemplate: templateNoOpts,
				Action: func(c *cli.Context) error {
					cmd.Initialize()

					if c.Args().Len() < 1 {
						return cli.ShowCommandHelp(c, "create")
					}

					return cmd.Create(c.Args())
				},
			},
			{
				Name:               "edit",
				Usage:              "edits credentials or config file for the current <setting> or the specified <setting>. ",
				CustomHelpTemplate: template,
				Action: func(c *cli.Context) error {
					cmd.Initialize()
					return cmd.Edit(c.Args(), c.String("type"))
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "type",
						Aliases:  []string{"t"},
						Required: true,
						Usage:    "the type, value must be 'cred[entials]' or 'conf[ig]'",
					},
				},
			},
			{
				Name:               "remove",
				Usage:              "removes a setting",
				Aliases:            []string{"rm"},
				CustomHelpTemplate: templateNoOpts,
				Action: func(c *cli.Context) error {
					cmd.Initialize()

					if c.Args().Len() < 1 {
						return cli.ShowCommandHelp(c, "remove")
					}

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
