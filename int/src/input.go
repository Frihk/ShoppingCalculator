package src

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	// "ShoppingCalculator/helper"
	"ShoppingCalculator/helper"
)


func Input() ([]helper.Input, []helper.ProductStorage){
	var input helper.Input
	var store helper.ProductStorage
	k := []helper.Input{}
	x := []helper.ProductStorage{}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		// if scanner.Scan() {
		fmt.Print("name of the item: ")
		scanner.Scan()
		input.ItemName = scanner.Text()
		if input.ItemName == "" {
			break
		}
		for {
			fmt.Print("Quantiy of the item: ")
			scanner.Scan()
			quantity :=scanner.Text()
			num, err := strconv.Atoi(quantity)
			if err != nil || num <= 0{
				fmt.Println("Invalid Quantity")
				continue
			}
			input.NumberOfItems = num
			break
		}
		for {
			fmt.Print("Price of the item: ")
			scanner.Scan()
			price := scanner.Text()
			cash, err := strconv.ParseFloat(price, 32)
				if err != nil || cash <= 0{
					fmt.Println("Invalid Price")
					continue
				}
			input.PriceOfItem = cash
			break
		}
		val := float64(input.NumberOfItems) * input.PriceOfItem
		input.Cost =  float64(val)//strconv.FormatFloat(val, 'f', 2, 32) 
		// }
		k = append(k, input)
		store = helper.ProductStorage{
			Name: input.ItemName,
			Price: input.PriceOfItem,
			Freq: 0,
		}
		x = append(x, store)
	}
		// int := Input(input.ItemName, input.NumberOfItems, input.PriceOfItem, input.Cost)
		return k, x

}

