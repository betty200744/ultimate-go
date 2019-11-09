package main

import (
	"fmt"
	"github.com/jinzhu/copier"
	"whgo/product/proto/source/tag"
)

type SkuEntity struct {
	Id     string   `json:"id"`
	TagIds []string `json:"tag_ids"`
}

type Sku struct {
	Id   string                 `json:"id"`
	Tags []*whale_tag_proto.Tag `json:"tag"`
}

func main() {
	skuEntity := new(SkuEntity)
	skuEntity.Id = "1"
	skuEntity.TagIds = []string{"a"}
	sku := new(Sku)
	copier.Copy(sku, skuEntity)
	fmt.Println(sku)
	tags := []*whale_tag_proto.Tag{{Id: "1", Name: "a"}}
	sku.Tags = tags
	fmt.Println(sku)
}
