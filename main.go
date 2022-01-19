package main

import (
	"fmt"
	"math/big"
)

type Item struct{
	Name string
	Price big.Rat
}
type Category struct {
	Name string
	Items []Item
}

func (c Category) AddItem(i Item){

}

func main() {
	fmt.Println("==== INVENTORY APP ====")
}
