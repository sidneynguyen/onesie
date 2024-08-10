package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	var port = ":6969"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Yo")
	})

	fmt.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(":6969", nil))
}
