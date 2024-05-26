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
	"flag"
	"log/slog"
	"os"

	"github.com/Fedena22/Holiday_bucket_tool/internal/database"
	"github.com/Fedena22/Holiday_bucket_tool/internal/templates"
	"github.com/valyala/fasthttp"
)

var dFlag = flag.Bool("d", false, "enable debug log messages")

type BaseHandler struct {
	db     *sql.DB
	logger *slog.Logger
}

func main() {
	flag.Parse()
	opts := slog.HandlerOptions{}
	if *dFlag {
		opts = slog.HandlerOptions{Level: slog.LevelDebug}
	}
	log := slog.New(slog.NewJSONHandler(os.Stdout, &opts))
	log.Debug("debug enabled")
	db, err := database.Open()
	if err != nil {
		log.Error("unexpected error when opening the database", "error", err)
	}
	err = database.Initialize(db)
	if err != nil {
		log.Error("failed to initialize database", "error", err)
	}
	err = database.TempData(db)
	if err != nil {
		log.Error("failed to initialize temp data", "error", err)
	}
	handler := BaseHandler{
		db:     db,
		logger: log,
	}

	log.Info("starting the server at http://localhost:8080 ...")
	err = fasthttp.ListenAndServe(":8080", handler.requestHandler)
	if err != nil {
		log.Error("unexpected error in server", "error", err)
	}
}

func (handler BaseHandler) requestHandler(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case "/":
		data, err := handler.getLocations()
		if err != nil {
			handler.logger.Error("cant get locations", "error", err)
			return
		}
		handler.mainPageHandler(ctx, data)
	case "/admin":
		data, err := handler.getLocations()
		if err != nil {
			handler.logger.Error("cant get locations admin ctx", "error", err)
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
	handler.logger.Debug("admin page entered")
	templates.WritePageTemplate(ctx, p)
}

func (handler BaseHandler) errorPageHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(fasthttp.StatusBadRequest)
}

func (handler BaseHandler) getLocations() ([]templates.Data, error) {
	var data []templates.Data
	handler.logger.Info("getting locations from database")
	locations, err := database.GetLocations(handler.db)
	if err != nil {
		return data, err
	}
	temp, _ := json.Marshal(locations)

	err = json.Unmarshal(temp, &data)
	if err != nil {
		return data, err
	}
	slog.Debug("getLocations finished", "locations", data)
	return data, nil
}
