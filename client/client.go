package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		state := r.URL.Query().Get("state")
		if code == "" {
			http.Error(w, "No code in the URL", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "Authorization Code: %s\nState: %s", code, state)
	})

	log.Println("Client is running at 9094 port.")
	log.Fatal(http.ListenAndServe(":9094", nil))
}
