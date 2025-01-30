package handlers

import (
	"net/http"

	"github.com/stierprogrammer/goshort/internal/storage"
);

func RedirectURL(w http.ResponseWriter, r *http.Request) {
	// Excluding the dash symbol
	shortCode := r.URL.Path[1:];
	
	// Gets the original URL
	originalURL, found := storage.Get(shortCode);

	// Checks if it is found
	if !found {
		http.NotFound(w, r);
		return;
	}

	// If it is found it redirects to the original URL
	http.Redirect(w, r, originalURL, http.StatusFound);
}