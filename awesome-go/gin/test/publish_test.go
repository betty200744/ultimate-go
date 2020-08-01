package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPublish(t *testing.T) {
	res, _ := MakeRequest("GET", "/publish", `{}`)
	assert.Equal(t, 200, res.Code)
	fmt.Println(res.Body.String())
}
