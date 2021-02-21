package handlers

import (
	"log"
	"net/http"
)

//GoodBye struct handler
type GoodBye struct {
	logger *log.Logger
}

// NewGoodBye constructor
func NewGoodBye(l *log.Logger) *GoodBye {
	return &GoodBye{l}
}

func (h *GoodBye) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	h.logger.Printf("GoodBye World!!")
}
