package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
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
		{"0", "title1", "text1"},
		{"1", "title2", "text2"},
		{"2", "title3", "text3"},
	},
}

func main() {

	router := http.NewServeMux()
	router.HandleFunc("/", mainPage)
	router.HandleFunc("/post", showPost)
	router.HandleFunc("/addpost", addPost)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", router))

}
func mainPage(servResp http.ResponseWriter, reqHand *http.Request) {

	tmpl := template.Must(template.New("first").ParseFiles("static/main.html"))

	_ = tmpl.ExecuteTemplate(servResp, "PostTmpl", BlogPage)

}

func showPost(servResp http.ResponseWriter, reqHand *http.Request) {
	tmpl := template.Must(template.New("first").ParseFiles("static/post.html"))
	getID := reqHand.URL.Query()["id"]
	for _, v := range BlogPage.List {
		if v.ID == getID[0] {
			i, _ := strconv.Atoi(v.ID)
			_ = tmpl.ExecuteTemplate(servResp, "ShowPost", BlogPage.List[i])
		}
	}
}

func addPost(servResp http.ResponseWriter, reqHand *http.Request) {
	getTitle := reqHand.URL.Query()["Title"]
	getText := reqHand.URL.Query()["Text"]
	getNewID := strconv.Itoa(len(BlogPage.List))
	newPostAdd := PostStr{getNewID, getTitle[0], getText[0]}
	BlogPage.List = append(BlogPage.List, newPostAdd)
}
