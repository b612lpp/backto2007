package posts

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//DBconn contains connection structure
type DBconn struct {
	ClientOptions *options.ClientOptions
	Client        *mongo.Client
	DB            *mongo.Collection
	Cntx          context.Context
}

//OpenConnection gets connection string and does connect
func OpenConnection() (zz DBconn) {

	clientOptionsBlog := options.Client().ApplyURI("mongodb://localhost:27017")
	Cntx := context.TODO()
	clientBlog, _ := mongo.Connect(Cntx, clientOptionsBlog)
	DBBlog := clientBlog.Database("blog").Collection("blog")

	zz = DBconn{clientOptionsBlog, clientBlog, DBBlog, Cntx}
	//err = clientBlog.Ping(Cntx, nil)

	return zz //, err
}
