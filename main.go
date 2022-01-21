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
	var input int
	setUpDefaultInfo()
	fmt.Println("==== INVENTORY APP ====")
	printMainNavigation()
	fmt.Scanf("%d", &input)
}

func setUpDefaultInfo () {
	item_1, item_2, item_3 := Item{"Chewing Gum", 1.99}, Item{"Soda", .99}, Item{"Chips", 2.89}
	defaultCategory.Items = append(defaultCategory.Items, item_1, item_2, item_3)
	categories = append(categories, defaultCategory)
}

func printMainNavigation() {
	var input int
	fmt.Println("==== Navigation ====")
	fmt.Println("1) Add a category \n 2) View categories \n 3) Add an Item \n 4) View an Item")
	n, _ := fmt.Scanf("%d", &input)
	if n == 0{
		return 
	}
	routeMainNavigation(input)
	
}

func routeMainNavigation(selection int) {
	switch selection {
	case 1:
		
		addNewCategory()
	case 2:
		fmt.Printf("%+v \n", categories)
	case 3:
		addNewItem()
	case 4: 
		viewItem()
	default:
		fmt.Printf("Invalid selection")
	}
}

func addNewCategory(){
	fmt.Printf("Name:")
}

func addNewItem(){
	fmt.Printf("Name:")
}

func viewItem(){
	fmt.Printf("Item Number:")
}
// Create user navigation
	// Categories
	// Items
	// Add Category
	// Delete Category
	// Add Item To Category
	// Remove Item From Category
	// Back