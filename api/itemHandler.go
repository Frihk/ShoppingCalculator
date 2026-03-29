package api

import (
	"encoding/json"
	"net/http"

	"ShoppingCalculator/helper"
	"ShoppingCalculator/int/src"
)

type CartItem struct {
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

type CheckoutRequest struct {
	Items []CartItem `json:"items"`
}

type CheckoutResponse struct {
	TotalQuantity int     `json:"totalQuantity"`
	TotalCost     float64 `json:"totalCost"`
	Message       string  `json:"message"`
}

func CheckoutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req CheckoutRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"invalid request body"}`, http.StatusBadRequest)
		return
	}

	if len(req.Items) == 0 {
		http.Error(w, `{"error":"items cannot be empty"}`, http.StatusBadRequest)
		return
	}

	inputs := make([]helper.Input, 0, len(req.Items))
	products := make([]helper.ProductStorage, 0, len(req.Items))

	for _, item := range req.Items {
		if item.Name == "" || item.Quantity <= 0 || item.Price <= 0 {
			http.Error(w, `{"error":"each item must have name, quantity > 0, and price > 0"}`, http.StatusBadRequest)
			return
		}

		cost := float64(item.Quantity) * item.Price
		inputs = append(inputs, helper.Input{
			ItemName:      item.Name,
			NumberOfItems: item.Quantity,
			PriceOfItem:   item.Price,
			Cost:          cost,
		})

		products = append(products, helper.ProductStorage{
			Name:  item.Name,
			Price: item.Price,
		})
	}

	result := src.Calc(inputs)
	if err := src.Jupdate(products); err != nil {
		http.Error(w, `{"error":"failed to save shopping log"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(CheckoutResponse{
		TotalQuantity: result.TotalQuantity,
		TotalCost:     result.TotalCost,
		Message:       "checkout saved",
	})
}
