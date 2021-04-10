package main

import (
	"net/http"
	"os"
	"time"

	"github.com/Fedena22/Holiday_bucket_tool/middleware"
	"github.com/gin-contrib/cors"

	"github.com/Fedena22/Holiday_bucket_tool/db"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

func router01() http.Handler {
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
	return router
}

func router02() http.Handler {
	e := gin.Default()

	e.Use(cors.Default())

	whitelist := make(map[string]bool)
	whitelist["127.0.0.1"] = true

	e.Use(middleware.IPWhiteList(whitelist))

	api := e.Group("/api")
	api.GET("/visetedlocations", db.GetVisitedLocations)
	api.GET("/notvisetedlocations", db.GetNotVisitedLocations)
	api.GET("/alllocations", db.GetAllLocations)
	api.DELETE("/deletelocations", db.DeleteLocations)
	api.POST("/addlocations", db.UpdateLocations)
	api.PUT("/newlocations", db.InsertLocations)

	return e

}
func main() {
	// gin.SetMode(gin.ReleaseMode)
	var webport string
	if os.Getenv("WEBPORT") != "" {
		webport = os.Getenv("WEBPORT")
	} else {
		webport = ":3000"
	}

	db.Migrate()
	log.Printf("Webport: %v", webport)

	server01 := &http.Server{
		Addr:         webport,
		Handler:      router01(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server02 := &http.Server{
		Addr:         ":1234",
		Handler:      router02(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		return server01.ListenAndServe()
	})
	g.Go(func() error {
		return server02.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
