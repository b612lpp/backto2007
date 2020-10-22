package main

import (
	"html/template"
	"log"
	"net/http"
)

//BlogPageStr - main page
type BlogPageStr struct {
	Name string
	List []PostStr
}

// PostStr - struct describes post
type PostStr struct {
	ID    string
	Title string
	Text  string
}

//BlogPage contains posts
var BlogPage = BlogPageStr{
	Name: "Main Page",
	List: []PostStr{
		{"1", "title1", "text1"},
		{"2", "title2", "text2"},
		{"3", "title3", "text3"},
	},
}

func main() {

	router := http.NewServeMux()
	router.HandleFunc("/", mainPage)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", router))

}
func mainPage(servResp http.ResponseWriter, reqHand *http.Request) {

	tmpl := template.Must(template.New("first").ParseFiles("static/main.html"))

	_ = tmpl.ExecuteTemplate(servResp, "PostTmpl", BlogPage)

}
