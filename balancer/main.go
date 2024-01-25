package main

import (
	"log"
	"net/http"
)

const INFO_MESSAGE string = "INFO]: FOR NOW WE ARE GOING TO ASSUME 2 SERVERS ARE ALWAYS ONLINE NO MORE NO LESS"

func main() {
	log.Println(INFO_MESSAGE)

	balancer := MakeBalancer()

	balancer.addServerURL("http://localhost:3000")
	balancer.addServerURL("http://localhost:3001")

	http.HandleFunc("/", balancer.handleRequest)

	http.ListenAndServe(":8000", nil)
}

/*func forwardRequest(w http.ResponseWriter, r *http.Request) {
	url, err := url.Parse("http://localhost:3000") //creating serve url
	if err != nil {
		log.Fatal(err)
	}

	//making our proxy
	reverseProxy := httputil.NewSingleHostReverseProxy(url)
	reverseProxy.ServeHTTP(w, r)
}*/
