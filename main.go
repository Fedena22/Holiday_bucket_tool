package main

import (
	"net/http"

	"github.com/Fedena22/Holiday_bucket_tool/db"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.LoadHTMLGlob("public/*.html")
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	admin := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"admin": "admin",
	}))

	admin.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin.html", nil)
	})
	admin.GET("/visetedlocations", db.GetVisitedLocations)
	admin.GET("/notvisetedlocations", db.GetNotVisitedLocations)
	admin.DELETE("/deletelocations", db.DeleteLocations)
	admin.POST("/addlocations", db.UpdateLocations)
	admin.PUT("/newlocations", db.InsertLocations)

	db.Migrate()
	r.Run(":3000")
}
