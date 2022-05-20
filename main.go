package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type word struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Meaning string `json:"meaning"`
	Details string `json:"details"`
	Status  string `json:"status"`
}

type wordHandler struct {
	words map[string]word
}

func (wh *wordHandler) handleRequest(c *gin.Context) {
	switch c.Request.Method {
	case "GET":
		wh.get(c)
	case "POST":
		wh.post(c)
	case "PUT":
		wh.put(c)
	case "DELETE":
		wh.delete(c)
	default:
		c.String(http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func (wh *wordHandler) get(c *gin.Context) {
	c.String(http.StatusOK, "Hello from get")
}

func (wh *wordHandler) post(c *gin.Context) {
	c.String(http.StatusOK, "Hello from post")
}

func (wh *wordHandler) put(c *gin.Context) {
	c.String(http.StatusOK, "Hello from put")
}

func (wh *wordHandler) delete(c *gin.Context) {
	c.String(http.StatusOK, "Hello from delete")
}

func main() {
	app := gin.Default()

	wh := wordHandler{words: map[string]word{}}

	app.Any("/", wh.handleRequest)
	app.Run()
}
