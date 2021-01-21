package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func postarticleHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "caught")
	fmt.Println("I display req!!!!")
	body, _ := ioutil.ReadAll(req.Body)
	fmt.Println(string(body))

}

func main() {
	fmt.Println("Hello!")

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
		log.Printf("defaulting to port %s", port)
	}

	http.HandleFunc("/postarticle", postarticleHandler)
	log.Printf("listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}

}
