package bit

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/knieriem/markdown"
	"html/template"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"
)

var (
	BlogPath = os.Getenv("GOPATH") + "/src/github.com/bitrho/bit/blog/"
)

type blog struct {
	Data  map[string]interface{}
	Posts []*Post
}

func NewBlog() *blog {
	return &blog{Data: make(map[string]interface{})}
}

func (b *blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	date := vars["date"]
	name := title + "-" + date + ".md"
	str, err := b.render(name)
	if err != nil {
		fmt.Fprintln(w, string(err.Error()))
	}
	b.Data["content"] = str
	b.Data["title"] = title
	b.Data["date"] = date
	t, err := template.ParseFiles(
		"template/bit/post.html",
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	t.Execute(w, b.Data)
}

func MarkdownToHtml(content string) (str string) {
	defer func() {
		e := recover()
		if e != nil {
			str = content
			log.Println("Render Markdown ERR:", e)
		}
	}()
	mdParser := markdown.NewParser(&markdown.Extensions{Smart: true})
	buf := bytes.NewBuffer(nil)
	mdParser.Markdown(bytes.NewBufferString(content), markdown.ToHTML(buf))
	str = buf.String()
	return
}

func (b *blog) render(name string) (str template.HTML, err error) {
	filePath := BlogPath + name
	if !IsFile(filePath) {
		err = errors.New("链接出错了")
		return
	}
	content := ReadFile(filePath)
	str = template.HTML(MarkdownToHtml(content))
	return
}

func (b *blog) getLatest() {
	names := ScanDir(BlogPath)
	var posts []*Post
	for _, value := range names {
		valueArr := strings.Split(value, "-")
		title := strings.TrimSpace(valueArr[0])
		date := strings.TrimSpace(strings.Trim(valueArr[1], ".md"))
		posts = append(posts, &Post{
			Title: title,
			Date:  date,
			Path:  value,
		})
	}
	sort.Sort(By(posts))
	if len(posts) > 10 {
		b.Posts = posts[:10]
	} else {
		b.Posts = posts
	}
	for _, one := range b.Posts {
		one.Url = "/blog/" + one.Date + "/" + one.Title + "/"
		content, _ := b.render(one.Path)
		one.Content = content
		the_time, _ := time.Parse("20060102", one.Date)
		one.Date = the_time.Format("Jan 02")
		one.Title = strings.Replace(one.Title, "_", " ", -1)
	}
}

type Post struct {
	Title    string
	Date     string
	Content  template.HTML
	Path     string
	Url      string
	Brief    template.HTML
	SubTitle string
}

type By []*Post

func (a By) Len() int           { return len(a) }
func (a By) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a By) Less(i, j int) bool { return a[i].Date > a[j].Date }
