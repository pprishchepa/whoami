package server

import (
	"context"
	"time"

	"github.com/fasthttp/router"
	"github.com/rs/zerolog/log"
	"github.com/valyala/fasthttp"
)

func Serve(ctx context.Context, opts Options) error {
	if opts.Debug {
		go printMemUsage(ctx, time.Second)
	}

	r := router.New()
	r.RedirectTrailingSlash = false
	r.RedirectFixedPath = false
	r.NotFound = func(ctx *fasthttp.RequestCtx) {
		whoamiHandler(ctx, opts.Name)
	}

	r.GET("/", func(ctx *fasthttp.RequestCtx) {
		whoamiHandler(ctx, opts.Name)
	})
	r.GET("/bench", benchHandler)
	r.GET("/data", dataHandler)
	r.GET("/random", dataHandler)

	srv := &fasthttp.Server{
		Handler:                       r.Handler,
		Name:                          "whoami",
		Concurrency:                   opts.Concurrency,
		ReadBufferSize:                int(opts.ReadBufferSize),
		WriteBufferSize:               int(opts.WriteBufferSize),
		ReadTimeout:                   opts.ReadTimeout,
		WriteTimeout:                  opts.WriteTimeout,
		MaxRequestBodySize:            int(opts.MaxRequestBodySize),
		DisableHeaderNamesNormalizing: true,
	}

	go func() {
		<-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.ShutdownWithContext(shutdownCtx); err != nil {
			log.Err(err).Msgf("failed to stop server gracefully ")
		}
	}()

	log.Info().Msgf("server \"%s\" serve on %v", opts.Name, opts.Addr)
	return srv.ListenAndServe(opts.Addr)
}
