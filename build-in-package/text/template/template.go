package template

import (
	"os"
	"text/template"
)

func TextTemplate() error {
	const templateText = `Input: {{printf "%q" .}}`
	tmpl, _ := template.New("test").Parse(templateText)
	err := tmpl.Execute(os.Stdout, "the go programming language")
	return err
}
