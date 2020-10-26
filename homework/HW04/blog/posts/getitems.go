package posts

import (
	"fmt"
)

// PostStr - struct describes post
type PostStr struct {
	ID    int
	Title string
	Text  string
}

//GetPostsList returns list of all posts
func GetPostsList() (PostsHeader []PostStr) {
	var tmpPostHeader PostStr

	headers, _ := db.Query("select id, title from posts;")
	defer headers.Close()
	for headers.Next() {
		err := headers.Scan(&tmpPostHeader.ID, &tmpPostHeader.Title)
		if err != nil {
			fmt.Println(err)
		}
		PostsHeader = append(PostsHeader, tmpPostHeader)
	}
	return PostsHeader
}

//GetPostByParam returns exact post
func GetPostByParam(getID int) (PostByIDStruct PostStr) {
	postByID := db.QueryRow("select * from posts where ID=?;", getID)
	err := postByID.Scan(&PostByIDStruct.ID, &PostByIDStruct.Title, &PostByIDStruct.Text)
	if err != nil {
		fmt.Println(err)
	}
	//defer postByID.Close()
	return PostByIDStruct
}

//InsertNewPost does what it must
func InsertNewPost(ID int, Title, Text string) {
	insReq, err := db.Prepare("Insert into posts values(?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	insReq.Exec(ID, Title, Text)

	defer db.Close()

}

//GetLastID does what it must. There is no autoincrement so we assign new ID by hands
func GetLastID() (NewPostID int) {
	lastPostIDReq := db.QueryRow("select max(id) from posts;")
	err := lastPostIDReq.Scan(&NewPostID)
	if err != nil {
		fmt.Println(err)
	}
	return NewPostID + 1

}
