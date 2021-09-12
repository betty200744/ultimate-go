package main

import (
	"os"
	"path/filepath"

	"github.com/gobuffalo/packr"
	"github.com/smallnest/rpcx/log"
	"github.com/urfave/cli/v2"
)

type Project struct {
	Name string
	Port int
	Auth string
	Path string
}

func NewProject(ctx *cli.Context) (err error) {
	if proj.Path != "" {
		if proj.Path, err = filepath.Abs(proj.Path); err != nil {
			return err
		}
	}
	CreateProject()
	return nil
}

func CreateProject() {
	box := packr.NewBox("./templates")
	err := os.MkdirAll(proj.Path, 0755)
	if err != nil {
		return
	}
	for _, s := range box.List() {
		// todo
		log.Info(s)
	}
	log.Info("create proj success")
}
