package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello From GO!")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Recieved connection")
		w.WriteHeader(200)
	})

	http.ListenAndServe(":8000", nil)
}
