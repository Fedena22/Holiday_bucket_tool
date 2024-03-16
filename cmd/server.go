// `go generate`:
//
//	go install github.com/valyala/quicktemplate/qtc
//
// Package main contains the main loop
//
//go:generate qtc -dir=../internal/templates
package main

import (
	"database/sql"
	"encoding/json"
	"log"

	"github.com/Fedena22/Holiday_bucket_tool/internal/database"
	"github.com/Fedena22/Holiday_bucket_tool/internal/templates"
	"github.com/valyala/fasthttp"
)

type BaseHandler struct {
	db *sql.DB
}

func main() {
	var handler BaseHandler
	log.Printf("Initialize the database")
	db, err := database.Open()
	if err != nil {
		log.Fatalf("unexpected error when opening the database: %s", err)
	}
	err = database.Initialize(db)
	if err != nil {
		log.Fatalf("failed to initialize database: %s", err)
	}
	err = database.TempData(db)
	if err != nil {
		log.Fatalf("failed to initialize temp data: %s", err)
	}
	handler.db = db
	log.Printf("starting the server at http://localhost:8080 ...")
	err = fasthttp.ListenAndServe(":8080", handler.requestHandler)
	if err != nil {
		log.Fatalf("unexpected error in server: %s", err)
	}
}

func (handler BaseHandler) requestHandler(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case "/":
		data, err := handler.getLocations()
		if err != nil {
			log.Println(err)
			return
		}
		handler.mainPageHandler(ctx, data)
	case "/admin":
		data, err := handler.getLocations()
		if err != nil {
			log.Println(err)
			return
		}
		handler.adminPageHandler(ctx, data, "admin")
	default:
		handler.errorPageHandler(ctx)
	}
	ctx.SetContentType("text/html; charset=utf-8")
}

func (handler BaseHandler) mainPageHandler(ctx *fasthttp.RequestCtx, locations []templates.Data) {
	p := &templates.MainPage{
		CTX:  ctx,
		Data: locations,
	}
	templates.WritePageTemplate(ctx, p)
}

func (handler BaseHandler) adminPageHandler(ctx *fasthttp.RequestCtx, locations []templates.Data, username string) {
	p := &templates.MainPage{
		CTX:      ctx,
		Data:     locations,
		Username: username,
	}
	templates.WritePageTemplate(ctx, p)
}

func (handler BaseHandler) errorPageHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(fasthttp.StatusBadRequest)
}

func (handler BaseHandler) getLocations() ([]templates.Data, error) {
	var data []templates.Data
	log.Println("getting locations from database")
	locations, err := database.GetLocations(handler.db)
	if err != nil {
		return data, err
	}
	temp, _ := json.Marshal(locations)

	err = json.Unmarshal(temp, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}
