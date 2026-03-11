package storage

import (
	// "encoding/json"
	"errors"
	"fmt"
	"os"
)

// func storage checks if a .json file already exists if it does then it just exist returning a bool value tha is sent to the db func, if it does not exist then it creates one

func Storage() {
	filepath := "/home/fian/frihk/ShoppingCalculator/storage"
	filename := filepath + "/shopinglogs.json"
	if !checker(filepath) {
		fmt.Println("path does not exist")
		return
	}
	if !checker (filename) {
		file, err := os.Create(filename)
		if err != nil {
			fmt.Println("There was an error Creating the file")
			return
		}
		defer file.Close()

		// encoder := json.NewEncoder(file)
		// if err := encoder.Encode(data); err != nil {
		// 	fmt.Println("Error encoding JSON:", err)
		// }
	} else {
		fmt.Println("it exists")
	}
}


func checker(path string) bool{
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}