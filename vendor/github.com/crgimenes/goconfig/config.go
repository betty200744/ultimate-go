//Package goconfig uses a struct as input and populates the
//fields of this struct with parameters fom command
//line, environment variables and configuration file.
package goconfig

import (
	"errors"
	"fmt"
	"path"

	"path/filepath"

	"github.com/crgimenes/goconfig/goenv"
	"github.com/crgimenes/goconfig/goflags"
	"github.com/crgimenes/goconfig/validate"
)

// Tag to set main name of field
var Tag = "cfg"

// TagDefault to set default value
var TagDefault = "cfgDefault"

// TagHelper to set usage help line
var TagHelper = "cfgHelper"

// Path sets default config path
var Path string

// File name of default config file
var File string

// FileRequired config file required
var FileRequired bool

// HelpString temporarily saves help
var HelpString string

// PrefixFlag is a string that would be placed at the beginning of the generated Flag tags.
var PrefixFlag string

// PrefixEnv is a string that would be placed at the beginning of the generated Event tags.
var PrefixEnv string

// ErrFileFormatNotDefined Is the error that is returned when there is no defined configuration file format.
var ErrFileFormatNotDefined = errors.New("file format not defined")

//Usage is a function to show the help, can be replaced by your own version.
var Usage func()

// Fileformat struct holds the functions to Load the file containing the settings
type Fileformat struct {
	Extension   string
	Load        func(config interface{}) (err error)
	PrepareHelp func(config interface{}) (help string, err error)
}

// Formats is the list of registered formats.
var Formats []Fileformat

func findFileFormat(extension string) (format Fileformat, err error) {
	format = Fileformat{}
	for _, f := range Formats {
		if f.Extension == extension {
			format = f
			return
		}
	}
	err = ErrFileFormatNotDefined
	return
}

func init() {
	Usage = DefaultUsage
	Path = "./"
	File = ""
	FileRequired = false
}

// Parse configuration
func Parse(config interface{}) (err error) {
	ext := path.Ext(File)
	if ext != "" {
		var format Fileformat
		format, err = findFileFormat(ext)
		if err != nil {
			return
		}
		err = format.Load(config)
		if err != nil {
			return
		}
		HelpString, err = format.PrepareHelp(config)
		if err != nil {
			return
		}
	}

	goenv.Prefix = PrefixEnv
	goenv.Setup(Tag, TagDefault)
	err = goenv.Parse(config)
	if err != nil {
		return
	}

	goflags.Prefix = PrefixFlag
	goflags.Setup(Tag, TagDefault, TagHelper)
	goflags.Usage = Usage
	goflags.Preserve = true
	err = goflags.Parse(config)
	if err != nil {
		return
	}

	validate.Prefix = PrefixFlag
	validate.Setup(Tag, TagDefault)
	err = validate.Parse(config)

	return
}

// PrintDefaults print the default help
func PrintDefaults() {
	if File != "" {
		fmt.Printf("Config file %q:\n", filepath.Join(Path, File))
		fmt.Println(HelpString)
	}
}

// DefaultUsage is assigned for Usage function by default
func DefaultUsage() {
	fmt.Println("Usage")
	goflags.PrintDefaults()
	goenv.PrintDefaults()
	PrintDefaults()
}
