package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// Hello is a simple handler
type Hello struct {
	l *log.Logger
}

// NewHello creates a new hello handler with the given logger.
func NewHello(l *log.Logger) *Hello {
	return &Hello{l: l}
}

// ServeHTTP implements the go http.Handler interface
// https://pkg.go.dev/net/http#Handler
func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.l.Println("Handle Hello request")

	// read the body
	b, err := io.ReadAll(r.Body)
	if err != nil {
		h.l.Println("Error reading body", err)

		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}

	// write the response
	fmt.Fprintf(w, "Hello, %s\n", b)
}
