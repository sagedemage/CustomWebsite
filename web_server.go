package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Serving HTTP on http://127.0.0.1:8000")
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.ListenAndServe(":8000", nil)
}
