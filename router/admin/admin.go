package admin

import (
	"fmt"
	"log"
	"net/http"
)

type admin struct {
	count int
}

func New() *admin {
	return &admin{}
}

func (a *admin) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// Log status
	log.Printf("admin: %v\n", a.count)
	fmt.Fprintln(w, "hello admin")

}
