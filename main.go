package main

import (
	"royan/cleanarch/app/config"
	"royan/cleanarch/app/database"
	"royan/cleanarch/app/router"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.InitConfig()
	dbMysql := database.InitDBPostgreSQL(cfg)
	database.InitialMigration(dbMysql)

	// create a new gin instance
	r := gin.Default()

	// Middleware CORS
	r.Use(ginCors())

	// Middleware RemoveTrailingSlash
	r.Use(ginRemoveTrailingSlash())

	// Middleware Logger
	r.Use(ginLogger())

	// Initialize router
	router.InitRouter(dbMysql, r)

	// start server and port
	r.Run(":8080")
}

func ginCors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Authorization, Content-Type, Content-Length, Accept-Encoding")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func ginRemoveTrailingSlash() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != "GET" && c.Request.Method != "HEAD" {
			c.Next()
			return
		}

		if len(c.Request.URL.Path) > 1 && c.Request.URL.Path[len(c.Request.URL.Path)-1] == '/' {
			c.Redirect(301, c.Request.URL.Path[:len(c.Request.URL.Path)-1])
			return
		}

		c.Next()
	}
}

func ginLogger() gin.HandlerFunc {
	return gin.Logger()
}
