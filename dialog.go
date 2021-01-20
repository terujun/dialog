package main

import (
	"fmt"
	"net/http"
)

func postarticleHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func main() {
	fmt.Println("Hello!")
	http.HandleFunc("/postarticle", postarticleHandler)
	http.ListenAndServe(":80", nil)
}
