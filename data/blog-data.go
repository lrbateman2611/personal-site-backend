package data

import (
	"context"
	"os"
	"time"

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

var client *mongo.Client

func init() {
	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGO_CONNECTION_STRING")))
	if err != nil {
		panic(err)
	}
}

func GetBlogs() []Blog {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// database and colletion code goes here
	db := client.Database("personal-site-blog")
	coll := db.Collection("blogs")

	// find code goes here
	cursor, err := coll.Find(ctx, bson.D{})
	if err != nil {
		panic(err)
	}

	var blogs []Blog

	if err = cursor.All(ctx, &blogs); err != nil {
		panic(err)
	}

	return blogs
}

func GetBlogById(blogId string) Blog {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	// database and colletion code goes here
	db := client.Database("personal-site-blog")
	coll := db.Collection("blogs")
	
	// convert id string to objectID
	id, err := primitive.ObjectIDFromHex(blogId)
	if err != nil {
		panic(err)
	}

	// find code goes here
	result := coll.FindOne(ctx, bson.D{{Key: "_id", Value: id}})
	
	var blog Blog
	
	if err := result.Decode(&blog); err != nil {
		panic(err)
	}
	
	return blog
}

func AddBlog(blog Blog) *mongo.InsertOneResult {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// database and colletion code goes here
	db := client.Database("personal-site-blog")
	coll := db.Collection("blogs")

	newId, err := coll.InsertOne(ctx, blog)
	if err != nil {
		panic(err)
	}
	return newId
}

func GetComments() []Comment {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// database and colletion code goes here
	db := client.Database("personal-site-blog")
	coll := db.Collection("comments")

	// find code goes here
	cursor, err := coll.Find(ctx, bson.D{})
	if err != nil {
		panic(err)
	}

	var comments []Comment

	if err = cursor.All(ctx, &comments); err != nil {
		panic(err)
	}

	return comments
}

func GetCommentsById(blogId string) []Comment {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// database and colletion code goes here
	db := client.Database("personal-site-blog")
	coll := db.Collection("comments")

	// convert id string to objectID
	id, err := primitive.ObjectIDFromHex(blogId)
	if err != nil {
		panic(err)
	}

	// find code goes here
	cursor, err := coll.Find(ctx, bson.D{{Key: "blog", Value: id}})
	if err != nil {
		panic(err)
	}

	var comments []Comment

	if err = cursor.All(ctx, &comments); err != nil {
		panic(err)
	}

	return comments
}

func AddComment(comment Comment) *mongo.InsertOneResult {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// database and colletion code goes here
	db := client.Database("personal-site-blog")
	coll := db.Collection("comments")

	newId, err := coll.InsertOne(ctx, comment)
	if err != nil {
		panic(err)
	}
	return newId
}
