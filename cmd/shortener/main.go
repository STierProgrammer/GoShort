package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/stierprogrammer/goshort/internal/handlers"
)

func main() {
	// I put this on purpose, cause I wanted to see how envs worked in GO
	port := os.Getenv("PORT");

	// Sets the default port to 8080
	if port == "" {
		port = "8080";
	}

	http.HandleFunc("/shorten", handlers.ShortenURL);
	http.HandleFunc("/", handlers.RedirectURL);

	fmt.Println("Server started on port: ", port);
	log.Fatal(http.ListenAndServe(":" + port, nil));
}