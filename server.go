package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
)

func main() {

	// generate a `Certificate` struct
	cert, _ := tls.LoadX509KeyPair("localhost.crt", "localhost.key")

	// create a custom server with `TLSConfig`
	s := &http.Server{
		Addr:    ":9000",
		Handler: nil, // use `http.DefaultServeMux`
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{cert},
		},
	}

	s2 := &http.Server{
		Addr:    ":8000",
		Handler: nil, // use `http.DefaultServeMux`
	}

	// handle `/` route
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello World!")
	})

	// run https server on port "9000"
	go func() {
		log.Fatal(s.ListenAndServeTLS("", ""))
	}()

	log.Fatal(s2.ListenAndServe())

}
