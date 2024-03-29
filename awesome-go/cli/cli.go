package main

import (
	"os"
	"os/exec"

	"github.com/smallnest/rpcx/log"
	"github.com/urfave/cli/v2"
)

// command line apps
/*
Usage:

	go mod <command> [arguments]

The commands are:

	download    download modules to local cache
	edit        edit go.mod from tools or scripts
	graph       print module requirement graph
	init        initialize new module in current directory
	tidy        add missing and remove unused modules
	vendor      make vendored copy of dependencies
	verify      verify dependencies have expected content
	why         explain why packages or modules are needed


    -svg             Outputs a graph in SVG format
    -top             Outputs top entries in text form
    -traces          Outputs all profile samples in text form
    -tree            Outputs a text rendering of call graph
    -web             Visualize graph through web browser

*/
// go build

func RunCmd(cmd string, args []string) {
	execCmd := exec.Cmd{
		Path:   cmd,
		Args:   append([]string{cmd}, args...),
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	/*	if filepath.Base(cmd) == cmd {
		lp, err := exec.LookPath(cmd)
		if err == nil {
			execCmd.Path = lp
		}
	}*/
	err := execCmd.Run()
	if err != nil {
		log.Error(err)
	}
}

var pprof Prof

func main() {
	app := cli.NewApp()
	app.Name = "gomod"
	app.Usage = "go mod <command> <global options> [arguments]"
	app.Commands = []*cli.Command{
		{
			Name:    "download",
			Usage:   "download modules to local cache",
			Aliases: []string{"d"},
			Action: func(context *cli.Context) error {
				RunCmd("/usr/local/go/bin/go", []string{"help", "mod", "download"})
				return nil
			},
		},
		{
			Name:    "edit",
			Usage:   "edit go.mod from tools or scripts",
			Aliases: []string{"e"},
			Action: func(context *cli.Context) error {
				RunCmd("/usr/local/go/bin/go", []string{"mod", "help", "edit"})
				return nil
			},
		},
		{
			Name:    "graph",
			Usage:   "print module requirement graph",
			Aliases: []string{"g"},
			Action: func(context *cli.Context) error {
				RunCmd("/usr/local/go/bin/go", []string{"mod", "graph"})
				return nil
			},
		},
		{
			Name:    "init",
			Usage:   "initialize new module in current directory",
			Aliases: []string{"i"},
			Action: func(context *cli.Context) error {
				RunCmd("/usr/local/go/bin/go", []string{"mod", "help", "init"})
				return nil
			},
		},
		{
			Name:    "tidy",
			Usage:   "add missing and remove unused modules",
			Aliases: []string{"t"},
			Action: func(context *cli.Context) error {
				RunCmd("/usr/local/go/bin/go", []string{"help", "mod", "tidy"})
				return nil
			},
		},
		{
			Name:    "why",
			Usage:   "explain why packages or modules are needed",
			Aliases: []string{"w"},
			Action: func(context *cli.Context) error {
				log.Info("action why")
				return nil
			},
		},
	}
	app.Flags = []cli.Flag{
		&cli.IntFlag{
			Name:        "port",
			Destination: &pprof.Port,
			Aliases:     []string{"p"},
			Value:       8888,
		},
		&cli.StringFlag{
			Name:        "svg",
			Destination: &pprof.Svg,
			Aliases:     []string{"s"},
			Value:       "svg",
		},
		&cli.StringFlag{
			Name:        "top",
			Value:       "top",
			Destination: &pprof.Top,
			Aliases:     []string{"t"},
		},
		&cli.StringFlag{
			Name:        "traces",
			Value:       "traces",
			Destination: &pprof.Traces,
		},
		&cli.StringFlag{
			Name:        "tree",
			Value:       "tree",
			Destination: &pprof.Tree,
		},
		&cli.StringFlag{
			Name:        "web",
			Value:       "web",
			Destination: &pprof.web,
			Aliases:     []string{"w"},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
	log.Info(pprof)
}
