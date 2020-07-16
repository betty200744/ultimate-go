package user

//go:generate mockgen -source=user.go -destination=./mock/mock.go
type Index interface {
	Get(key string) interface{}
	GetTwo(key1, key2 string) (v1, v2 interface{})
	Put(key string, value interface{})
}
