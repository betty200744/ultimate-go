package template

import (
	"fmt"
	"testing"
)

func TestHtmlTemplate(t *testing.T) {
	err := HtmlTemplate()
	fmt.Println(err)
}
