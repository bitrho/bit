package bit

import (
	"fmt"
	"net/http"
)

type search struct {
	Data map[string]interface{}
}

func NewSearch() *search {
	return &search{}
}

func (s *search) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "comming soon")
}
