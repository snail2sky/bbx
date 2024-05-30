package server

import (
	"fmt"
	get2 "github.com/snail2sky/bbx/app/server/get"
	"github.com/spf13/cobra"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong\n")
	})

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/ping")
	})

	get := r.Group("/get")

	get.GET("/ai8", get2.Ai8Num)
	return r
}

func RunServer(cmd *cobra.Command) {
	r := setupRouter()

	host := cmd.Flag("host").Value.String()
	port := cmd.Flag("port").Value.String()
	addr := fmt.Sprintf("%s:%s", host, port)

	err := r.Run(addr)
	if err != nil {
		log.Fatal(err)
	}
}
