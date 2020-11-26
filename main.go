package main

import (
	"github.com/cdimascio/aws-config-manager/cmd"
	"github.com/cdimascio/aws-config-manager/template"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"sort"
)

func main() {
	app := &cli.App{

		Usage: "AWS Config and Credentials Manager",
		CustomAppHelpTemplate: template.AppTemplate,
		Description: "Manages many .aws/credentials and .aws/config files as settings",
		Commands: []*cli.Command{
			{
				Name:               "use",
				Usage:              "sets the current setting",
				CustomHelpTemplate: template.CmdTemplateNoOpts,
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
				CustomHelpTemplate: template.CmdTemplateNoOpts,
				Action: func(c *cli.Context) error {
					cmd.Initialize()
					return cmd.Current()
				},
			},
			{
				Name:               "list",
				Aliases:            []string{"ls"},
				Usage:              "list all settings",
				CustomHelpTemplate: template.CmdTemplateNoOpts,
				Action: func(c *cli.Context) error {
					cmd.Initialize()
					return cmd.List()
				},
			},
			{
				Name:               "create",
				Usage:              "creates a new empty setting.",
				CustomHelpTemplate: template.CmdTemplateNoOpts,
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
				Usage:              "edits a credentials or config file.",
				CustomHelpTemplate: template.CmdTemplate,
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
				CustomHelpTemplate: template.CmdTemplateNoOpts,
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
