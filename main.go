package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body := `"bob":{"name":"Robert","age":25}`
		time.Sleep(time.Second * 3)
		fmt.Println("Hit the endpoint")
		fmt.Fprint(w, body)
	})

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		body := `"bob":{"name":"Robert","age":25}`
		fmt.Println("Hit the api endpoint")
		fmt.Fprint(w, body)
	})

	fmt.Println("Serving on localhost:8081")
	http.ListenAndServe(":8081", nil)
}
