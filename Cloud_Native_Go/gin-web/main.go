package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	eng := gin.Default()

	eng.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	eng.GET("/api/books", func(c *gin.Context) {
		c.JSON(http.StatusOK, AllBooks())
	})

	eng.POST("/api/books", func(c *gin.Context) {
		var book Book
		if c.BindJSON(&book) == nil {
			isbn, created := CreateBook(book)
			if created {
				c.Header("Location", "/api/books/"+isbn)
				c.Status(http.StatusCreated)
			} else {
				c.Status(http.StatusConflict)
			}
		}
	})

	eng.GET("/api/books/:isbn", func(c *gin.Context) {
		isbn := c.Params.ByName("isbn")
		book, found := GetBook(isbn)
		if found {
			c.JSON(http.StatusOK, book)
		} else {
			c.AbortWithStatus(http.StatusNotFound)
		}
	})

	eng.PUT("/api/books/:isbn", func(c *gin.Context) {
		isbn := c.Params.ByName("isbn")

		var book Book
		if c.BindJSON(&book) == nil {
			exists := UpdateBook(isbn, book)
			if exists {
				c.Status(http.StatusOK)
			} else {
				c.Status(http.StatusNotFound)
			}
		}
	})

	eng.DELETE("/api/books/:isbn", func(c *gin.Context) {
		isbn := c.Params.ByName("isbn")
		DeleteBook(isbn)
		c.Status(http.StatusOK)
	})

	eng.Run(port())
}

func port() string {
	// Good practice make port param
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8091"
	}
	return ":" + port
}
