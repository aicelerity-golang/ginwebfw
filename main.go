package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"webfw/util"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	log.Println("..Application running in environment: ", config.RuntimeSetup, " and on port: ", config.AppPort)

	router := gin.Default()

	// # Static Files
	router.StaticFile("/", "./public/index.html")
	router.Static("/public", "./public")

	// # Http methods
	router.GET("/employee", func(ctx *gin.Context) {
		ctx.File("./public/employee.html")
	})

	router.POST("/employee", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "New request POSted successfully")
	})

	// # Parameterised Routes
	router.GET("/employee/:username/*rest", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"username": ctx.Param("username"),
			"rest":     ctx.Param("rest"),
		})
	})

	// # Route Groups
	adminGroup := router.Group("/admin")

	adminGroup.GET("/users", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Page to administer users")
	})

	adminGroup.GET("/roles", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Page to administer roles")
	})

	adminGroup.GET("/policies", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Page to administer policies")
	})

	log.Fatal(router.Run(":3000"))
}
