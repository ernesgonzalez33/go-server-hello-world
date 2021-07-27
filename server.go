package main

import (
	"crypto/tls"
	"io"
	"log"
	"net/http"
)

func httpHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world from HTTP!\n")
}

func httpsHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world from HTTPS!\n")
}

func main() {

	// generate a `Certificate` struct
	cert, _ := tls.LoadX509KeyPair("localhost.crt", "localhost.key")

	// create a custom server with `TLSConfig`
	s := &http.Server{
		Addr:    ":9000",
		Handler: http.HandlerFunc(httpsHandler), // use `http.DefaultServeMux`
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{cert},
		},
	}

	s2 := &http.Server{
		Addr:    ":8000",
		Handler: http.HandlerFunc(httpHandler), // use `http.DefaultServeMux`
	}

	// // handle `/` route
	// http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
	// 	fmt.Fprint(res, "Hello World!")
	// })

	// run https server on port "9000"
	go func() {
		log.Fatal(s.ListenAndServeTLS("", ""))
	}()

	log.Fatal(s2.ListenAndServe())

}
