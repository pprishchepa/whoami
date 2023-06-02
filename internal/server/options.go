package server

import (
	"math"
	"time"

	"github.com/docker/go-units"
)

type Options struct {
	Name               string
	Addr               string
	Debug              bool
	Concurrency        int
	ReadBufferSize     int64
	WriteBufferSize    int64
	ReadTimeout        time.Duration
	WriteTimeout       time.Duration
	MaxRequestBodySize int64
}

//goland:noinspection GoUnusedGlobalVariable
var DefaultOptions = Options{
	Name:               "undefined",
	Addr:               ":8081",
	Debug:              false,
	Concurrency:        math.MaxInt, //unlimited
	ReadBufferSize:     10 * units.KiB,
	WriteBufferSize:    10 * units.KiB,
	ReadTimeout:        5 * time.Second,
	WriteTimeout:       5 * time.Second,
	MaxRequestBodySize: 10 * units.MiB,
}
