package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("views/*.html")

	r.GET("/", index)
	r.GET("/ping", ping)
	r.GET("/ping2", ping2)

	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func index(c *gin.Context) {
	c.HTML(200, "helloworld.html", gin.H{
		"data": "Hello Go/Gin World.",
	})
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong1",
	})
}

func ping2(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong2",
	})
}
