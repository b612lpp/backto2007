package main

import (
	"backto2007/homework/HW06/blog/posts"
	"log"
	"net/http"
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

	AllPosts := posts.GetPostsList()
	tmpl := template.Must(template.New("first").ParseFiles("static/main.html"))
	_ = tmpl.ExecuteTemplate(servResp, "PostTmpl", AllPosts)

}

//Shows single post by specified id
func showPost(servResp http.ResponseWriter, reqHand *http.Request) {
	tmpl := template.Must(template.New("first").ParseFiles("static/singlpost.html"))
	getIDStr := reqHand.URL.Query().Get("id")
	postRequestedByID := posts.GetPostByParam(getIDStr)
	_ = tmpl.ExecuteTemplate(servResp, "ShowPost", postRequestedByID)
}

//Handler gets request with parameters. "editCheck = true" means existing post with specified ID is updating. If there is no parameter handler creates new post
func addPost(servResp http.ResponseWriter, reqHand *http.Request) {

	editCheck := reqHand.FormValue("Edit")
	newPostTitle := reqHand.FormValue("Title")
	newPostText := reqHand.FormValue("Text")
	curentID := reqHand.FormValue("id")
	if editCheck == "true" {
		posts.UpdatetNewPost(curentID, newPostTitle, newPostText)
		http.Redirect(servResp, reqHand, "/", 301)
	} else {

		posts.InsertNewPost(newPostTitle, newPostText)
		http.Redirect(servResp, reqHand, "/", 301)
	}
}
