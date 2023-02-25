package server

import (
	"math"
	"time"

	"github.com/alecthomas/units"
)

type Options struct {
	Addr               string
	Debug              bool
	Concurrency        int
	ReadBufferSize     units.Base2Bytes
	WriteBufferSize    units.Base2Bytes
	ReadTimeout        time.Duration
	WriteTimeout       time.Duration
	MaxRequestBodySize units.Base2Bytes
}

//goland:noinspection GoUnusedGlobalVariable
var DefaultOptions = Options{
	Addr:               ":8081",
	Debug:              false,
	Concurrency:        math.MaxInt, //unlimited
	ReadBufferSize:     10 * units.KiB,
	WriteBufferSize:    10 * units.KiB,
	ReadTimeout:        5 * time.Second,
	WriteTimeout:       5 * time.Second,
	MaxRequestBodySize: 10 * units.MiB,
}
