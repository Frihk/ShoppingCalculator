package main

import (
	"fmt"

	"ShoppingCalculator/int/src"
)

func main() {
	inputs, products := src.Input()
	result := src.Calc(inputs)
	// if err != nil{
	// 	fmt.Println("Error in calculation logic")
	// }
	fmt.Println(result)
	err := src.Jupdate(products)
	if err != nil{
		fmt.Println("Error updating the Json file")
	}
	
}
