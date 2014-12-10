package bit

import (
	"html/template"
	"log"
	"net/http"
	//"runtime/debug"
	//"github.com/bitrho/bit/model/user"
	//"github.com/gorilla/sessions"
)

type index struct {
	Data      map[string]interface{}
	templates []string
}

var TEMPLATE_PREFIX = "template/bit/"

func NewIndex() *index {
	templates := []string{
		"index.html",
		"header.html",
		"footer.html",
		"about.html",
		"blog.html",
		"conferences.html",
		"credit.html",
		"education.html",
		"experience.html",
		"github.html",
		"info.html",
		"languages.html",
		"latest.html",
		"music.html",
		"page_header.html",
		"project.html",
		"skill.html",
		"testimonials.html",
	}
	return &index{
		templates: templates,
	}
}

func (i *index) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles(
		"template/bit/index.html",
		"template/bit/header.html",
		"template/bit/footer.html",
		"template/bit/about.html",
		"template/bit/blog.html",
		"template/bit/conferences.html",
		"template/bit/credit.html",
		"template/bit/education.html",
		"template/bit/experience.html",
		"template/bit/github.html",
		"template/bit/info.html",

		"template/bit/languages.html",
		"template/bit/latest.html",
		"template/bit/music.html",
		"template/bit/page_header.html",
		"template/bit/project.html",
		"template/bit/skill.html",
		"template/bit/testimonials.html",
	)
	if err != nil {
		log.Fatal(err.Error())
	}

	t.Execute(w, i)
}
