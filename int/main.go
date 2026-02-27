package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	// "int/src"
	// "Shopping-calculator/int/src"
)

type Input struct {
	ItemName string
	NumberOfItems int 
	PriceOfItem float

}
  
	 
func main() {
	var input Input
	// ItemName.item := ""
	// quantity := ""
	// price := ""

	scanner := bufio.NewScanner(os.Stdin)
	// if scanner.Scan() {
	fmt.Print("name of the item: ")
	scanner.Scan()
	input.ItemName = scanner.Text()

	fmt.Print("Quantiy of the item: ")
	scanner.Scan()
	input.NumberOfItems = strconv.Itoa(scanner.Text())

	fmt.Print("Price of the item: ")
	scanner.Scan()
	input.PriceOfItem = strconv.ParseFloat(scanner.Text())
	// }
	fmt.Print(input)
}