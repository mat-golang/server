package main

import (
	"fmt"
	"net/http"

	h "ws-chat.example/http"
)

func main() {
	fmt.Println("Listening to :8080")

	// http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	template.Must(template.ParseFiles("pages/index.html")).Execute(w, nil)
	// })

	http.ListenAndServe(":8080", h.NewService())
}
