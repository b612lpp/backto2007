package posts

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// PostStr - struct describes post
type PostStr struct {
	QID   primitive.ObjectID `bson:"_id"`
	ID    int                `json:"id"`
	Title string             `json:"title"`
	Text  string             `json:"text"`
}

//NewPostStr container for new or modified post
type NewPostStr struct {
	NewTitle string `bson:"title"`
	NewText  string `bson:"text"`
}

//GetPostsList returns list of all posts
func GetPostsList() (tmpPostHeader []PostStr) {

	qq := OpenConnection()
	rowData, _ := qq.DB.Find(qq.Cntx, bson.D{})
	_ = rowData.All(qq.Cntx, &tmpPostHeader)
	return tmpPostHeader
}

//GetPostByParam returns exact post
func GetPostByParam(getID string) (PostByIDStruct PostStr) {

	qq := OpenConnection()
	id, _ := primitive.ObjectIDFromHex(getID)
	filter := bson.M{"_id": id}
	_ = (qq.DB.FindOne(qq.Cntx, filter)).Decode(&PostByIDStruct)
	return PostByIDStruct
}

//InsertNewPost does what it must
func InsertNewPost(Title, Text string) {
	qq := OpenConnection()

	var newPost = NewPostStr{Title, Text}
	result, err := qq.DB.InsertOne(qq.Cntx, newPost)
	fmt.Println(result, err)

}

//UpdatetNewPost updates existing post
func UpdatetNewPost(ID string, Title, Text string) {
	qq := OpenConnection()
	id, _ := primitive.ObjectIDFromHex(ID)
	updPost := bson.M{"$set": bson.M{"title": Title, "text": Text}}
	qq.DB.FindOneAndUpdate(qq.Cntx, bson.M{"_id": id}, updPost)

}
