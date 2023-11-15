package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

var proj Project

func main() {
	app := cli.NewApp()
	app.Name = "packr"
	app.Usage = "create project"
	app.Commands = []*cli.Command{{
		Name:   "create",
		Action: NewProject,
	}}
	app.Flags = []cli.Flag{
		cli.HelpFlag,
		&cli.StringFlag{
			Name:        "name",
			Aliases:     []string{"n"},
			Value:       "",
			Destination: &proj.Name,
		},
		&cli.IntFlag{
			Name:        "port",
			Aliases:     []string{"p"},
			Value:       50001,
			Destination: &proj.Port,
		},
		&cli.StringFlag{
			Name:        "auth",
			Aliases:     []string{"a"},
			Value:       "",
			Destination: &proj.Auth,
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
