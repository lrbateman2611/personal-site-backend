package api

import (
	"encoding/json"
	"net/http"

	"example/personal-site-backend/data"

	"github.com/gin-gonic/gin"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/blogs", getBlogs)
		api.GET("/blogs/:id", getBlogById)
		api.POST("/blog", postBlog)
	
		api.GET("/comments", getComments)
		api.GET("/comments/:id", getCommentsById)
		api.POST("/comment", postComment)
	}
	
	router.ServeHTTP(w, r)
}

func getBlogs(c *gin.Context) {
	blogs := data.GetBlogs()

	// Convert comments to JSON
	blogsJSON, err := json.Marshal(blogs)
	if err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusOK, string(blogsJSON))
}

func getBlogById(c *gin.Context) {
	id := c.Param("id")
	
	blog := data.GetBlogById(id)

	// Convert comments to JSON
	blogJSON, err := json.Marshal(blog)
	if err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusOK, string(blogJSON))
}

func postBlog(c *gin.Context) {
	var newBlog data.Blog

	if err := c.BindJSON(&newBlog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}

	response := data.AddBlog(newBlog)
	if response == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add blog"})
        return
	}

	c.IndentedJSON(http.StatusCreated, response)
}

func getComments(c *gin.Context) {
	comments := data.GetComments()

	// Convert comments to JSON
	commentsJSON, err := json.Marshal(comments)
	if err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusOK, string(commentsJSON))
}

func getCommentsById(c *gin.Context) {
	id := c.Param("id")
	
	comments := data.GetCommentsById(id)

	// Convert comments to JSON
	commentsJSON, err := json.Marshal(comments)
	if err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusOK, string(commentsJSON))
}

func postComment(c *gin.Context) {
	var newComment data.Comment

	if err := c.BindJSON(&newComment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}

	response := data.AddComment(newComment)
	if response == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add comment"})
        return
	}

	c.IndentedJSON(http.StatusCreated, response)
}