package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("New request")
		fmt.Fprintf(w, "\nPath: %v \n", r.URL.Path)
		if r.URL.RawQuery != "" {
			fmt.Fprintf(w, "Parameters: \n")
			for _, q := range strings.Split(r.URL.RawQuery, "&") {
				fmt.Fprintf(w, "\t %v \n", q)
			}
		}
		if r.Header != nil {
			fmt.Fprintf(w, "Headers: \n")
			for k, v := range r.Header {
				fmt.Fprintf(w, "\t %v: %v \n", k, v)
			}
		}
	})
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
