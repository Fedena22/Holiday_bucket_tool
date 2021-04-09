package main

import (
	"net/http"
	"os"

	"github.com/Fedena22/Holiday_bucket_tool/db"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	var adminuser, adminpass string

	if os.Getenv("ADMINUSER") != "" {
		adminuser = os.Getenv("ADMINUSER")
	} else {
		adminuser = "admin"
	}
	if os.Getenv("ADMINPASSWORD") != "" {
		adminpass = os.Getenv("ADMINPASSWORD")
	} else {
		adminpass = "admin"
	}

	router := gin.Default()
	router.LoadHTMLGlob("public/*.html")
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	router.Use(static.Serve("/assets", static.LocalFile("./assets", true)))
	admin := router.Group("/admin", gin.BasicAuth(gin.Accounts{
		adminuser: adminpass,
	}))

	admin.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin.html", nil)
	})
	admin.GET("/visetedlocations", db.GetVisitedLocations)
	admin.GET("/notvisetedlocations", db.GetNotVisitedLocations)
	admin.GET("/alllocations", db.GetAllLocations)
	admin.DELETE("/deletelocations", db.DeleteLocations)
	admin.POST("/addlocations", db.UpdateLocations)
	admin.PUT("/newlocations", db.InsertLocations)

	db.Migrate()
	router.Run(":3000")
}
