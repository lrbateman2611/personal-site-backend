package data

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Blog struct {
	ID       	primitive.ObjectID 	`bson:"_id,omitempty" json:"_id,omitempty"`
	Author	 	string							`bson:"author" json:"author"`
	Content 	string							`bson:"content" json:"content"`
}

type Comment struct {
	ID				primitive.ObjectID	`bson:"_id,omitempty" json:"_id,omitempty"`
	Blog			primitive.ObjectID	`bson:"blog" json:"blog"`
	Author	 	string							`bson:"author" json:"author"`
	Content 	string							`bson:"content" json:"content"`
}

func GetBlogs() []Blog {
	client := connect()
	defer disconnect(client)

	// database and colletion code goes here
	db := client.Database("personal-site-blog")
	coll := db.Collection("blogs")

	// find code goes here
	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}

	var blogs []Blog

	if err = cursor.All(context.TODO(), &blogs); err != nil {
		panic(err)
	}

	return blogs
}

func GetBlogById(blogId string) Blog {
	client := connect()
	defer disconnect(client)
	
	// database and colletion code goes here
	db := client.Database("personal-site-blog")
	coll := db.Collection("blogs")
	
	// convert id string to objectID
	id, err := primitive.ObjectIDFromHex(blogId)
	if err != nil {
		panic(err)
	}

	// find code goes here
	result := coll.FindOne(context.TODO(), bson.D{{Key: "_id", Value: id}})
	
	var blog Blog
	
	if err := result.Decode(&blog); err != nil {
		panic(err)
	}
	
	return blog
}

//TODO: get unused ID from the db and use it here
func AddBlog(blog Blog) *mongo.InsertOneResult {
	client := connect()
	defer disconnect(client)

	// database and colletion code goes here
	db := client.Database("personal-site-blog")
	coll := db.Collection("blogs")

	newId, err := coll.InsertOne(context.TODO(), blog)
	if err != nil {
		panic(err)
	}
	return newId
}

func GetComments() []Comment {
	client := connect()
	defer disconnect(client)

	// database and colletion code goes here
	db := client.Database("personal-site-blog")
	coll := db.Collection("comments")

	// find code goes here
	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}

	var comments []Comment

	if err = cursor.All(context.TODO(), &comments); err != nil {
		panic(err)
	}

	return comments
}

func GetCommentsById(blogId string) []Comment {
	client := connect()
	defer disconnect(client)

	// database and colletion code goes here
	db := client.Database("personal-site-blog")
	coll := db.Collection("comments")

	// convert id string to objectID
	id, err := primitive.ObjectIDFromHex(blogId)
	if err != nil {
		panic(err)
	}

	// find code goes here
	cursor, err := coll.Find(context.TODO(), bson.D{{Key: "blog", Value: id}})
	if err != nil {
		panic(err)
	}

	var comments []Comment

	if err = cursor.All(context.TODO(), &comments); err != nil {
		panic(err)
	}

	return comments
}

func AddComment(comment Comment) *mongo.InsertOneResult {
	client := connect()
	defer disconnect(client)

	// database and colletion code goes here
	db := client.Database("personal-site-blog")
	coll := db.Collection("comments")

	newId, err := coll.InsertOne(context.TODO(), comment)
	if err != nil {
		panic(err)
	}
	return newId
}

func connect() *mongo.Client {
	connectionString := os.Getenv("MONGO_CONNECTION_STRING")
	if connectionString == "" {
		panic("MONGO_CONNECTION_STRING must be set")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connectionString))
	if err != nil {
		panic(err)
	}
	return client
}

func disconnect(client *mongo.Client) {
	err := client.Disconnect(context.TODO()); 
	if err != nil {
		panic(err)
	}
}