package main

import (
	"fmt"
	"ShoppingCalculator/int/src"
	
)


// type Input struct {
// 	ItemName string
// 	NumberOfItems string 
// 	PriceOfItem string
// 	Cost string

// }
  
	 
func main() {
	list := src.Input()
	
	fmt.Println(list)
}