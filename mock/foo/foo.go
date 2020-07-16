package foo

//go:generate mockgen -source=foo.go -destination=mock/mock.go
type Foo interface {
	Bar(x int) int
}
