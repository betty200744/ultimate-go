package main

import (
	"testing"
)

func TestFormatDemo(t *testing.T) {
	FormatDemo()
}

func TestFormatGMT(t *testing.T) {
	FormatGMT()
}

func TestFormatBy(t *testing.T) {
	FormatBy("Mon, 02 Jan 2006 15:04:05 GMT")
}

func TestFormatStandard(t *testing.T) {
	FormatStandard()
}
