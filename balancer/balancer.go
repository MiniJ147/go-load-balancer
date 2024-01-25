package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Balancer struct {
	CurrentServerPtr uint16
	ProxyNumber      uint16
	Proxies          []*httputil.ReverseProxy
}

func MakeBalancer() *Balancer {
	balancer := Balancer{
		CurrentServerPtr: 0,
		ProxyNumber:      0,
		Proxies:          []*httputil.ReverseProxy{},
	}

	return &balancer
}

func (balancer *Balancer) addServerURL(reqURL string) {
	newURL, err := url.Parse(reqURL)
	if err != nil {
		log.Fatal(err)
		return
	}

	newProxy := httputil.NewSingleHostReverseProxy(newURL)

	balancer.Proxies = append(balancer.Proxies, newProxy)
	balancer.ProxyNumber += 1
}

func (bal *Balancer) handleRequest(w http.ResponseWriter, r *http.Request) {
	log.Printf("Sending request to server %v\n", (bal.CurrentServerPtr + 1))

	bal.Proxies[bal.CurrentServerPtr].ServeHTTP(w, r)

	if bal.CurrentServerPtr+1 >= bal.ProxyNumber {
		bal.CurrentServerPtr = 0
	} else {
		bal.CurrentServerPtr += 1
	}

	log.Printf("Now pointing at server %v\n\n", (bal.CurrentServerPtr + 1))
}
