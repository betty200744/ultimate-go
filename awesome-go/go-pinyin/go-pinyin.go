package go_pinyin

import (
	"database/sql"
)

// https://github.com/mozillazg/go-pinyin

type Person struct {
	name string
}
type DBConnection interface {
	Exec(string, ...interface{}) (sql.Result, error)
}

func (p *Person) Name() string {
	return p.name
}
func (p *Person) SetName(name string) {
	p.name = name
}
func NewPerson(name string) *Person {
	return &Person{name: name}
}
