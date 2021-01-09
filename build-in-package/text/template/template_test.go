package template

import (
	"fmt"
	"testing"
)

func TestTextTemplate(t *testing.T) {
	t.Run("template", func(t *testing.T) {
		err := TextTemplate()
		fmt.Println(err)
	})
}
