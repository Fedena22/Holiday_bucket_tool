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
	}
	templates.WritePageTemplate(ctx, p)
}

func errorPageHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(fasthttp.StatusBadRequest)
}
