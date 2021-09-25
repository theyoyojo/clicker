package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Game struct {
	Clock   int
	Credits int
}

func test(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "test")
}

func getgame(g Game) (handler func(c *gin.Context)) {

	return func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, g)
	}
}

func ginrouter() *gin.Engine {

	var g Game

	g.Clock = 0
	g.Credits = 666

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/game", getgame(g))

	return r
}

func main() {
	// http.HandleFunc("/", test)

	r := ginrouter()

	// g Game

	fmt.Printf("yeet\n")
	// log.Fatal(http.ListenAndServe(":3000", nil))

	r.Run(":3000")

}
