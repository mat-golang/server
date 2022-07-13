package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	fmt.Println("Listening to :8080")

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		template.Must(template.ParseFiles("pages/index.html")).Execute(w, nil)
	})

	http.ListenAndServe(":8080", nil)
}
