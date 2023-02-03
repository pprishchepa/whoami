package main

import (
	"fmt"
	"log"

	"github.com/valyala/fasthttp"
)

func main() {
	if err := fasthttp.ListenAndServe(":8081", requestHandler); err != nil {
		log.Fatalf("fasthttp: listen and serve: %v", err)
	}
}

func requestHandler(ctx *fasthttp.RequestCtx) {
	_, _ = fmt.Fprintf(ctx, "Request: %s %s\n", ctx.Method(), ctx.URI().String())
	_, _ = fmt.Fprintf(ctx, "Host: %q\n", ctx.Host())
	_, _ = fmt.Fprintf(ctx, "IP: %q\n\n", ctx.RemoteIP())

	ctx.SetContentType("text/plain; charset=utf8")
}
