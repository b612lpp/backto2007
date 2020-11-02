package main

import (
	"backto2007/homework/HW04/blog/posts"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

//variable contains service type. depens on the type db connection parameters are requesting
var servicename string = "blog"

func main() {

	router := http.NewServeMux()
	router.HandleFunc("/", mainPage)
	router.HandleFunc("/post", showPost)
	router.HandleFunc("/addpost", addPost)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", router))

}

//shows all titles as URL
func mainPage(servResp http.ResponseWriter, reqHand *http.Request) {
	posts.GetConnParams(servicename).OpenConnection()
	Zz := posts.GetPostsList()

	tmpl := template.Must(template.New("first").ParseFiles("static/main.html"))

	_ = tmpl.ExecuteTemplate(servResp, "PostTmpl", Zz)

}

//Shows single post by specified id
func showPost(servResp http.ResponseWriter, reqHand *http.Request) {
	tmpl := template.Must(template.New("first").ParseFiles("static/post.html"))
	posts.GetConnParams(servicename).OpenConnection()
	getIDStr := reqHand.URL.Query().Get("id")
	getIDInt, _ := strconv.Atoi(getIDStr)
	postRequestedByID := posts.GetPostByParam(getIDInt)
	_ = tmpl.ExecuteTemplate(servResp, "ShowPost", postRequestedByID)
}

//Handler gets request with parameters. "editCheck = true" means existing post with specified ID is updating. If there is no parameter handler creates new post
func addPost(servResp http.ResponseWriter, reqHand *http.Request) {
	posts.GetConnParams(servicename).OpenConnection()
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
