package bit

import (
	"html/template"
	"log"
	"net/http"
)

type admin struct {
	Data map[string]interface{}
}

func NewAdmin() *admin {
	return &admin{}
}

func (s *admin) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(
		"template/bit/admin.html",
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	t.Execute(w, s)
}
