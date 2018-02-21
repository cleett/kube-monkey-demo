package main

import (
	"net/http"
	"os"

	"log"
)

func main() {
	g := os.Getenv("GREETING")
	if len(g) == 0 {
		g = "Default greeting here..."
	}
	log.Printf("greeting: %s\n", g)

	log.Println("starting server...")

	log.Fatal(http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s\n", r.Method, r.URL.Path)
		w.Write([]byte(g))
	})))
}
