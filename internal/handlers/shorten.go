package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/stierprogrammer/goshort/internal/services"
)

type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}

func ShortenURL(w http.ResponseWriter, r *http.Request) {
	var req ShortenRequest;

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest);
		return;
	}

	shortURL := services.Shorten(req.URL);

	res := ShortenResponse{ShortURL: shortURL};

	w.Header().Set("Content-Type", "application/json");

	json.NewEncoder(w).Encode(res);
}