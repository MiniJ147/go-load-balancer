package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	fmt.Println("Hello From GO!")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Recieved connection")

		forwardRequest(w, r)
	})

	http.ListenAndServe(":8000", nil)
}

func forwardRequest(w http.ResponseWriter, r *http.Request) {
	url, err := url.Parse("http://localhost:3000") //creating serve url
	if err != nil {
		log.Fatal(err)
	}

	//making our proxy
	reverseProxy := httputil.NewSingleHostReverseProxy(url)
	reverseProxy.ServeHTTP(w, r)
}
