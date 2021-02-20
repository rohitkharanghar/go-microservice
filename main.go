package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Printf("Hello World!!")
		data, _ := ioutil.ReadAll(r.Body)
		log.Printf("Data %s", data)

	})

	http.HandleFunc("/goodBye", func(rw http.ResponseWriter, r *http.Request) {
		log.Printf("GoodBye World!!")
	})

	http.ListenAndServe(":9090", nil)
}
