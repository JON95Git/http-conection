package main

import (
	"log"
	"net/http"
)

func homePageExample(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("URI custom Response"))
}

func main() {

	// Define the URI
	http.HandleFunc("/home", homePageExample)

	// Listen in port 5000
	log.Fatal(http.ListenAndServe(":5000", nil))
}
