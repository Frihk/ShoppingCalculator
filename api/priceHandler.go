package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"ShoppingCalculator/int/src"
)

func SuggestionsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	limit := 10
	if rawLimit := r.URL.Query().Get("limit"); rawLimit != "" {
		parsedLimit, err := strconv.Atoi(rawLimit)
		if err != nil {
			http.Error(w, `{"error":"limit must be an integer between 5 and 10"}`, http.StatusBadRequest)
			return
		}
		if parsedLimit < 5 || parsedLimit > 10 {
			http.Error(w, `{"error":"limit must be between 5 and 10"}`, http.StatusBadRequest)
			return
		}
		limit = parsedLimit
	}

	suggestions, err := src.Suggestions(limit)
	if err != nil {
		http.Error(w, `{"error":"failed to load suggestions"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]any{"suggestions": suggestions})
}
