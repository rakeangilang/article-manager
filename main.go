package main

import (
	"fmt"

	"github.com/rakeangilang/article-manager/sheesh"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	//initializeRoutes()
	// router.GET("/", func(c *gin.Context) {
	// 	c.HTML(
	// 		http.StatusOK,
	// 		"index.html",
	// 		gin.H{
	// 			"title": "Home Page",
	// 		},
	// 	)
	// })

	fmt.Println(sheesh.Ya)

	router.Run()
}