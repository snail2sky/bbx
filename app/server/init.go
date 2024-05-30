package server

import (
	"fmt"
	get2 "github.com/snail2sky/bbx/app/server/get/car"
	"github.com/spf13/cobra"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter(debug bool) *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong\n")
	})

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/ping")
	})

	get := r.Group("/get")
	getCar := get.Group("/car")

	getCar.GET("/ai8", get2.Ai8Num)
	return r
}

func RunServer(cmd *cobra.Command) {

	host := cmd.Flag("host").Value.String()
	port := cmd.Flag("port").Value.String()
	debug, _ := cmd.Flags().GetBool("debug")

	r := setupRouter(debug)

	addr := fmt.Sprintf("%s:%s", host, port)

	err := r.Run(addr)
	if err != nil {
		log.Fatal(err)
	}
}
