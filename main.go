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

var categories []Category
var defaultCategory = Category{Name: "General", Items: make([]Item, 0)}



func (c Category) AddItem(i Item){

}

func main() {
	setUpDefaultInfo()
	fmt.Println("==== INVENTORY APP ====")
	fmt.Printf("%+v \n", categories)
}

func setUpDefaultInfo () {
	item_1, item_2, item_3 := Item{"Chewing Gum", 1.99}, Item{"Soda", .99}, Item{"Chips", 2.89}
	defaultCategory.Items = append(defaultCategory.Items, item_1, item_2, item_3)
	categories = append(categories, defaultCategory)
}
// Create user navigation
	// Categories
	// Items
	// Add Category
	// Delete Category
	// Add Item To Category
	// Remove Item From Category
	// Back