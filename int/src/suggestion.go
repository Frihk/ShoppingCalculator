package src

import (
	"encoding/json"
	"os"

	"ShoppingCalculator/helper"
)

const (
	minSuggestionLimit = 5
	maxSuggestionLimit = 10
)

func normalizeSuggestionLimit(limit int) int {
	if limit < minSuggestionLimit {
		return minSuggestionLimit
	}
	if limit > maxSuggestionLimit {
		return maxSuggestionLimit
	}
	return limit
}

func Suggestions(limit int) ([]helper.ProductStorage, error) {
	limit = normalizeSuggestionLimit(limit)

	file, err := os.Open(logPath())
	if err != nil {
		if os.IsNotExist(err) {
			return []helper.ProductStorage{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var products []helper.ProductStorage
	if err := json.NewDecoder(file).Decode(&products); err != nil {
		return nil, err
	}

	if len(products) > limit {
		return products[:limit], nil
	}
	return products, nil
}
