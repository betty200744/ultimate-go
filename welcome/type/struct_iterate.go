package main

import (
	"fmt"
	"reflect"
)

type SearchSkuRequest struct {
	Keywords       string   `protobuf:"bytes,1,opt,name=Keywords,proto3" json:"Keywords,omitempty"`
	TagIds         []string `protobuf:"bytes,2,rep,name=TagIds,proto3" json:"TagIds,omitempty"`
	Limit          int64    `protobuf:"varint,3,opt,name=Limit,proto3" json:"Limit,omitempty"`
	Offset         int64    `protobuf:"varint,4,opt,name=Offset,proto3" json:"Offset,omitempty"`
	CompanyId      string   `protobuf:"bytes,5,opt,name=CompanyId,proto3" json:"CompanyId,omitempty"`
	Barcode        string   `protobuf:"bytes,6,opt,name=Barcode,proto3" json:"Barcode,omitempty"`
	ClassId        int64    `protobuf:"varint,7,opt,name=ClassId,proto3" json:"ClassId,omitempty"`
	Id             int64    `protobuf:"varint,8,opt,name=Id,proto3" json:"Id,omitempty"`
	Name           string   `protobuf:"bytes,9,opt,name=Name,proto3" json:"Name,omitempty"`
	SubBrandId     int64    `protobuf:"varint,11,opt,name=SubBrandId,proto3" json:"SubBrandId,omitempty"`
	MainCategory   int32    `protobuf:"varint,12,opt,name=MainCategory,proto3" json:"MainCategory,omitempty"`
	MiddleCategory int32    `protobuf:"varint,13,opt,name=MiddleCategory,proto3" json:"MiddleCategory,omitempty"`
	MinorCategory  int32    `protobuf:"varint,14,opt,name=MinorCategory,proto3" json:"MinorCategory,omitempty"`
	Description    string   `protobuf:"bytes,17,opt,name=Description,proto3" json:"Description,omitempty"`
	AboutBrand     string   `protobuf:"bytes,18,opt,name=AboutBrand,proto3" json:"AboutBrand,omitempty"`
}

/*func main() {
	req := SearchSkuRequest{
		Keywords:"k",
		CompanyId:"c",
		Barcode:"b",
		Name:"n",
		Description:"d",
		AboutBrand:"a",
		ClassId:int64(1),
		Id:int64(1),
		SubBrandId:int64(1),
		MainCategory:int32(1),
		MiddleCategory:int32(1),
		MinorCategory:int32(1),
	}

	v := reflect.ValueOf(req)
	nvalues := make([]interface{}, v.NumField())
	fmt.Println(v.Field(0))
	fmt.Println(v.Field(0).String())
	for i := 0; i < v.NumField() ; i++ {
		nvalues[i] = v.Field(i).Interface()
	}
	fmt.Println(nvalues)


	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("name: %+v, value: %+v (%T)\n",
			v.Type().Field(i).Name, // Name attribute gives us the struct's key
			v.Field(i).Interface()) // Interface() provides memory address of the value
	}

}

func main() {
	f := "foo"
	b := "bar"

	x := struct {
		Foo *string
		Bar *string
	}{&f, &b}

	v := reflect.ValueOf(x)

	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("name: %+v, value: %+v (%T)\n",
			v.Type().Field(i).Name, // Name attribute gives us the struct's key
			v.Field(i).Elem(), 	// Elem() dereferences the pointer value
			v.Field(i).Interface()) // Interface() provides memory address of the value
	}


}
*/

type Range struct {
	From int64
	To   int64
}

func main() {
	type T struct {
		A      int32
		B      int64
		C      string
		TagIds []string
		Price  Range
	}
	t := T{int32(32), int64(64), "skidoo", []string{"t1", "t2"}, Range{From: int64(2), To: int64(3)}}
	s := reflect.ValueOf(&t).Elem()
	fmt.Println(s)
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Println("this is value: ", reflect.Value(f))
		fmt.Printf("index : %d, name: %s , type:  %s , value:  %v\n", i, s.Type().Field(i).Name, f.Type(), f.Interface())
	}

}
