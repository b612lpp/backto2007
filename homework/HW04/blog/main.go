package main

import (
	"backto2007/homework/HW04/tmp/blog/posts"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

func main() {

	router := http.NewServeMux()
	router.HandleFunc("/", mainPage)
	router.HandleFunc("/post", showPost)
	router.HandleFunc("/addpost", addPost)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", router))

}

func mainPage(servResp http.ResponseWriter, reqHand *http.Request) {
	posts.OpenConnection("root:password@tcp(127.0.0.1:3306)/blog")
	Zz := posts.GetPostsList()

	tmpl := template.Must(template.New("first").ParseFiles("static/main.html"))

	_ = tmpl.ExecuteTemplate(servResp, "PostTmpl", Zz)

}

func showPost(servResp http.ResponseWriter, reqHand *http.Request) {
	tmpl := template.Must(template.New("first").ParseFiles("static/post.html"))
	posts.OpenConnection("root:password@tcp(127.0.0.1:3306)/blog")
	getIDStr := reqHand.URL.Query().Get("id")
	getIDInt, _ := strconv.Atoi(getIDStr)
	Qq := posts.GetPostByParam(getIDInt)
	_ = tmpl.ExecuteTemplate(servResp, "ShowPost", Qq)
}

func addPost(servResp http.ResponseWriter, reqHand *http.Request) {
	posts.OpenConnection("root:password@tcp(127.0.0.1:3306)/blog")
	newPostTitle := reqHand.FormValue("Title")
	newPostText := reqHand.FormValue("Text")
	NewID := posts.GetLastID()

	posts.InsertNewPost(NewID, newPostTitle, newPostText)
	http.Redirect(servResp, reqHand, "/", 301)
}
