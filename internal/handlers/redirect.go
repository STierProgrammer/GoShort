package handlers

import (
	"net/http"

	"github.com/stierprogrammer/goshort/internal/storage"
)

func RedirectURL(w http.ResponseWriter, r *http.Request) {
	shortCode := r.URL.Path[1:];

	originalURL, found := storage.Get(shortCode);

	if !found {
		http.NotFound(w, r);
		return;
	}

	http.Redirect(w, r, originalURL, http.StatusFound);
}