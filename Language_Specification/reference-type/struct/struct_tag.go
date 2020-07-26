package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

/*
1. only exported Fields can be json output
2. use tag name when has struct tag name
*/
// struct with no tag
type User struct {
	Name      string
	Password  string
	Phone     int64
	CreatedAt time.Time
	Ignoring  int64
}

// struct with json tag
type User1 struct {
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	Phone     int64     `json:"phone, omitempty"`
	CreatedAt time.Time `json:"created_at"`
	Ignoring  int64     `json:"-"`
}

// Get Tags
func PrintTags(s interface{}) {
	t := reflect.TypeOf(s).Elem()
	v := reflect.ValueOf(s).Elem()
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		vf := v.Field(i)
		fmt.Printf("Sturct PrintTags index: %v, field: %v, value: %v, type: %v, tag: %v \n", i, f.Name, vf.Interface(), vf.Type().Name(), f.Tag.Get("json"))
	}
}

// Depopulate to map
func Depopulate(s interface{}, values map[string]string) {
	st := reflect.TypeOf(s).Elem()
	v := reflect.ValueOf(s).Elem()

	for i := 0; i < st.NumField(); i++ {
		f := st.Field(i)
		key := f.Tag.Get("json")
		values[key] = v.Field(i).String()
	}
}
func main() {
	user := &User{
		Name:      "n",
		Password:  "p",
		Phone:     11,
		CreatedAt: time.Time{},
	}
	user1 := &User1{
		Name:      "n",
		Password:  "p",
		Phone:     11,
		CreatedAt: time.Time{},
	}
	// json.Marshal Tags
	out, _ := json.Marshal(user)
	out1, _ := json.Marshal(user1)
	fmt.Println("user marshal: ", string(out))
	fmt.Println("user1 with json tag marshal: ", string(out1))

	// json.Unmarshal Tags, 如name小写，对应json tag name
	userUm := `
{
  	"name":      "n",
	"password":  "p",
	"phone":     11
}
`
	user2 := &User1{}
	_ = json.Unmarshal([]byte(userUm), user2)
	fmt.Printf("json.Unmarshal Tags: %#v \n", user2)
	// Print json Tags
	PrintTags(user1)
	userMap := make(map[string]string)
	Depopulate(user1, userMap)
	fmt.Printf("Struct Populate: %v \n", userMap)
}
