package main

import (
	"fmt"
	"log"
	"net/http"
)

func http_logger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	fmt.Println("Serving HTTP on http://127.0.0.1:8000")
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.ListenAndServe(":8000", http_logger(http.DefaultServeMux))
}
