package bit

import (
	"html/template"
	"log"
	"net/http"
)

type index struct {
	Data map[string]interface{}
}

func NewIndex() *index {
	return &index{}
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
	i.Data = make(map[string]interface{})
	blog := NewBlog()
	blog.getLatest()
	i.Data["blog"] = blog
	t.Execute(w, i)
}
