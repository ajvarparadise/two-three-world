package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

/*
  - connect to db

  - setup routes

  - start server

  - Game logic

  - is action allowed based on game state
*/

type Player struct {
	ID       int32  `json:"id"`
	Username string `json:"username"`
}

type GamePlayer struct {
	ID       int32 `json:"id"`
	GameID   int32 `json:"game"`
	PlayerId int32 `json:"playerId"`
}

type Game struct {
	ID     int32  `json:"id"`
	Status int    `json:"status"`
	Title  string `json:"title"`
}

var games = []Game{
	{ID: 1, Status: 0, Title: "Happy"},
	{ID: 2, Status: 1, Title: "Pappi"},
}

func CreateNewGame(c *gin.Context) {
	var newGame Game
	if err := c.ShouldBindJSON(&newGame); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	newGame.ID = xid.New().Counter()
	games = append(games, newGame)
	c.JSON(http.StatusCreated, newGame)
}

func IndexRoute(c *gin.Context) {
	c.Status(http.StatusOK)
}
func PingRoute(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
func main() {
	r := gin.Default()
	r.GET("/", IndexRoute)
	r.GET("/ping", PingRoute)
	r.POST("/new-game", CreateNewGame)

	r.Run() // listen and serve on 0.0.0.0:8080
}
