package main

type SkuEntity struct {
	Id     string   `json:"id"`
	TagIds []string `json:"tag_ids"`
}

func main() {
	skuEntity := new(SkuEntity)
	skuEntity.Id = "1"
	skuEntity.TagIds = []string{"a"}
}
