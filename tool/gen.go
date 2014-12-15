package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/bitrho/bit/router/bit"
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	"html/template"
	"log"
	"strings"
)

func main() {
	fmt.Println("gen posts start!")
	gen()
	fmt.Println("gen posts success!")
}

func gen() {
	names := bit.ScanDir(bit.BlogPath)

	//写入mongodb
	session, err := mgo.Dial("localhost")
	if err != nil {
		return
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	session.EnsureSafe(&mgo.Safe{W: 1})
	c := session.DB("bit").C("post")
	c.DropCollection()

	for _, value := range names {
		valueArr := strings.Split(value, "-")
		title := strings.TrimSpace(valueArr[0])
		date := strings.TrimSpace(strings.Trim(valueArr[1], ".md"))
		rawContent := bit.ReadFile(bit.BlogPath + "/" + value)
		content := bit.MarkdownToHtml(rawContent)
		log.Println(content)
		contentReader := strings.NewReader(content)
		doc, err := goquery.NewDocumentFromReader(contentReader)
		if err != nil {
			log.Println(string(err.Error()))
		}
		log.Println(doc)
		err = c.Insert(&bit.Post{
			Title:    title,
			Date:     date,
			Path:     value,
			Content:  template.HTML(content),
			Url:      "/blog/" + date + "/" + title + "/",
			Brief:    template.HTML(""),
			SubTitle: "",
		})
		if err != nil {
			log.Println(string(err.Error()))
		}
	}

}
