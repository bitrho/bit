package bit

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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

func (a *admin) get(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(
		"template/bit/admin.html",
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	t.Execute(w, a)
}

func (a *admin) post(w http.ResponseWriter, r *http.Request) {

	defer func() {
		if re := recover(); re != nil {
			log.Println("Recovered in f", re)
		}
	}()

	session, err := mgo.Dial("localhost")
	if err != nil {
		return
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("bit").C("blog")
	err = c.Insert(u)

	AjaxReturn(4004, "", w)
}

func (a *admin) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		a.post(w, r)
	} else {
		a.get(w, r)
	}
}
