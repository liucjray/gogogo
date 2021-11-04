package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

var (
	token string
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("views/*.html")

	r.GET("/", index)
	r.GET("/ping", ping)
	r.Any("/api/line/send_text", SendText)

	r.Run(":" + os.Getenv("PORT"))
}

func index(c *gin.Context) {
	c.HTML(200, "helloworld.html", gin.H{
		"data": "Hello Go/Gin World.",
	})
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func SendText(c *gin.Context) {

	// token = os.Getenv("LINE_NOTIFY_ACCESS_TOKEN")

	token := c.DefaultPostForm("token", "")
	text := c.DefaultPostForm("text", "")

	if token == "" || text == "" {
		fmt.Println("幹智障喔")
	}

	v := url.Values{}
	v.Set("message", text)
	Request("POST", "https://notify-api.line.me/api/notify", v, token)

}

func Request(method string, url string, data url.Values, token string) ([]byte, error) {
	client := &http.Client{}

	res, err := http.NewRequest(method, url, strings.NewReader(data.Encode()))

	if err != nil {
		return nil, err
	}

	res.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	if token != "" {
		res.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	}

	resp, err := client.Do(res)

	if err != nil {
		return nil, err
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Println(err.Error())
		}
	}()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}
