package template

import (
	"html/template"
	"os"
)

type Data struct {
	Title string
	Items []string
}

func HtmlTemplate() error {
	const tpl = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>{{.Title}}</title>
	</head>
	<body>
		{{range .Items}}<div>{{ . }}</div>{{else}}<div><strong>no rows</strong></div>{{end}}
	</body>
</html>`

	data := Data{
		Title: "hello",
		Items: []string{"a", "b", "c"},
	}
	tmp, _ := template.New("html test").Parse(tpl)
	return tmp.Execute(os.Stdout, data)
}
