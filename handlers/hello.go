package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
)

// Hello handler for handling hello type requests
type Hello struct {
	logger *log.Logger
}

// NewHello constructor
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	h.logger.Printf("Hello World!!")
	data, _ := ioutil.ReadAll(req.Body)
	h.logger.Printf("Data %s", data)
}
