package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Hello!\n"))
	})

	log.Fatal(http.ListenAndServe("0.0.0.0:80", nil))
}
