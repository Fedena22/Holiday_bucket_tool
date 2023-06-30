// `go generate`:
//
//	go install github.com/valyala/quicktemplate/qtc
//
// Package main contains the main loop
//
//go:generate qtc -dir=../internal/templates
package main

import (
	"log"

	"github.com/Fedena22/Holiday_bucket_tool/internal/templates"
	"github.com/valyala/fasthttp"
)

func main() {
	log.Printf("starting the server at http://localhost:8080 ...")
	err := fasthttp.ListenAndServe(":8080", requestHandler)
	if err != nil {
		log.Fatalf("unexpected error in server: %s", err)
	}
}

func requestHandler(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case "/":
		mainPageHandler(ctx)
	default:
		errorPageHandler(ctx)
	}
	ctx.SetContentType("text/html; charset=utf-8")
}

func mainPageHandler(ctx *fasthttp.RequestCtx) {
	p := &templates.MainPage{
		CTX: ctx,
		Data: []templates.Data{
			{
				Number:    0,
				Placename: "Kyoto palace",
				Latitude:  35.02509,
				Longitude: 135.76193,
				Visited:   true,
			},
			{
				Number:    1,
				Placename: "test2",
				Latitude:  1234,
				Longitude: 4321,
				Visited:   true,
				Username:  "admin",
			},
			{
				Number:    2,
				Placename: "Osaka trainstation",
				Latitude:  34.7332,
				Longitude: 135.49928,
			},
			{
				Number:    3,
				Placename: "test3",
				Latitude:  1234,
				Longitude: 4321,
				Visited:   true,
			},
		},
	}
	templates.WritePageTemplate(ctx, p)
}

func errorPageHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(fasthttp.StatusBadRequest)
}
