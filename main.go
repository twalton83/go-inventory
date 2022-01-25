package main

import (
	"fmt"
	// "math/big"
	"strings"
	"sort"
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
		viewCategories()
	case 3:
		addNewItem()
	case 4: 
		viewItem()
	default:
		fmt.Printf("Invalid selection")
	}
}

func addNewCategory(){
	var input string
	fmt.Printf("Name:\n")
	n, _ := fmt.Scanf("%s", &input)

	if n == 0 {
		return
	}
	newCategory := Category{Name: input, Items: make([]Item, 0)}
	categories = append(categories, newCategory)
	fmt.Printf("Category Added\n")
	printMainNavigation()
}

func viewCategories(){
	for _, category := range categories {
		fmt.Printf("%+v \n", category.Name)
	}

	fmt.Printf("View category? Y/N\n")	
	var input string
	n, _ := fmt.Scanf("%s", &input)
	if n == 0  || strings.ToLower(input) == "n" {
		printMainNavigation()
	}

	if input == "Y" || input == "y" {
		fmt.Printf("Enter category name:\n")	
		var category string
		n, _ := fmt.Scanf("%s", &category)
		if n == 0  {
		printMainNavigation()
		}
		viewACategory(strings.ToLower(category))
	}
}

func viewACategory(categoryName string){
	var selectedCategory Category
	for _, category := range categories {
		if strings.EqualFold(category.Name, categoryName) {
			selectedCategory = category
		}
	}
	fmt.Print(selectedCategory.Name)
	for _, item := range selectedCategory.Items {
		fmt.Print(item.Name, " - ", "$", item.Price, "\n")
	}
}
// TODO, fix bug where it skips when a space is involved
func addNewItem(){
	var itemName string
	var categoryName string
	var price float32
	fmt.Printf("Name:\n")
	fmt.Scanf("%s", &itemName)

	fmt.Printf("Price:\n")
	fmt.Scanf("%f", &price)

	newItem := Item{Name: itemName, Price: price}

	fmt.Printf("Category name:\n")
	fmt.Scanf("%s", &categoryName)

	categoryIndex := sort.Search(len(categories), func(int) bool {
		return categoryName != "" && categories[0].Name == categoryName
	})

	categories[categoryIndex].Items = append(categories[categoryIndex].Items, newItem)
	fmt.Printf("Item added!\n")
	printMainNavigation()
}

func viewItem(){
	fmt.Printf("Item Number:")
}
