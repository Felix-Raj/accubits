package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func main() {
	r := gin.Default()
	r.RedirectTrailingSlash = true

	r.GET("/search", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": fmt.Sprintf("searching for %s", c.Query("query")),
		})
	})

	r.POST("/fetch", func(c *gin.Context) {
		query := c.PostForm("query")
		limit, err := strconv.Atoi(c.PostForm("limit"))
		if err != nil {
			log.Fatal("cannot parse limit, err: ", err)
		}
		GetCourses(query, limit)
		c.JSON(200, gin.H{
			"message": fmt.Sprintf("fetch for %s, with limit %d", query, limit),
		})
	})

	// r.OPTIONS("/fetch", func(c *gin.Context) {
	// 	c.Header("Access-Control-Allow-Origin", "https://hoppscotch.io") // to use with https://hoppscotch.io/
	// 	c.JSON(200, struct{}{})
	// })

	r.Run() // listen and serve on 0.0.0.0:8080
}
