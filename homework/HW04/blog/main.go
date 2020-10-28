package main

import (
	"backto2007/homework/HW04/blog/posts"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

var connectionString string = "root:password@tcp(127.0.0.1:3306)/blog"

func main() {

	router := http.NewServeMux()
	router.HandleFunc("/", mainPage)
	router.HandleFunc("/post", showPost)
	router.HandleFunc("/addpost", addPost)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", router))

}

func mainPage(servResp http.ResponseWriter, reqHand *http.Request) {
	posts.OpenConnection(connectionString)
	Zz := posts.GetPostsList()

	tmpl := template.Must(template.New("first").ParseFiles("static/main.html"))

	_ = tmpl.ExecuteTemplate(servResp, "PostTmpl", Zz)

}

func showPost(servResp http.ResponseWriter, reqHand *http.Request) {
	tmpl := template.Must(template.New("first").ParseFiles("static/post.html"))
	posts.OpenConnection(connectionString)
	getIDStr := reqHand.URL.Query().Get("id")
	getIDInt, _ := strconv.Atoi(getIDStr)
	Qq := posts.GetPostByParam(getIDInt)
	_ = tmpl.ExecuteTemplate(servResp, "ShowPost", Qq)
}

func addPost(servResp http.ResponseWriter, reqHand *http.Request) {
	posts.OpenConnection(connectionString)
	editCheck := reqHand.FormValue("Edit")
	newPostTitle := reqHand.FormValue("Title")
	newPostText := reqHand.FormValue("Text")
	curentID, _ := strconv.Atoi(reqHand.FormValue("id"))
	if editCheck == "true" {
		posts.UpdatetNewPost(curentID, newPostTitle, newPostText)
		http.Redirect(servResp, reqHand, "/", 301)
	} else {
		NewID := posts.GetLastID()
		posts.InsertNewPost(NewID, newPostTitle, newPostText)
		http.Redirect(servResp, reqHand, "/", 301)
	}
}
