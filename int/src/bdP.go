package src

import (
	"fmt"
	"os"
)

//this script is incharge of saving, editing the data, loading the data from the json file where data is saved.
// I'll also try to make sure the data is sorted by the frequency/ how many times the consumer buys the commodity, this will help when am making the suggestions

func CreateJson() {
	// var shoping *os.File
	file, err := os.Create("shoping.json")
	if err != nil {
		fmt.Println("Error")
	}
	defer file.Close()
}