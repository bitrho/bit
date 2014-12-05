package index

import (
	"fmt"
	"log"
	"net/http"
)

type index struct {
	count int
}

func New() *index {
	return &index{}
}

func (i *index) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// Log status
	log.Printf("index: %v\n", i.count)

	fmt.Fprintln(w, "hello index")
}
