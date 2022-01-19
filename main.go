package main

import (
	"fmt"
	// "math/big"
)

type Item struct{
	Name string
	Price float32
}
type Category struct {
	Name string
	Items []Item
}

func (c Category) AddItem(i Item){

}




func main() {
	fmt.Println("==== INVENTORY APP ====")
	defaultItem := Item{"Chewing Gum", 1.99}
	defaultCategory := Category{Name: "General", Items: make([]Item, 0)}
	defaultCategory.Items = append(defaultCategory.Items, defaultItem)
	var categories []Category
	categories = append(categories, defaultCategory)

	fmt.Printf("%+v \n", categories)
}
