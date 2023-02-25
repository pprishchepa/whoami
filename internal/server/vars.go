package server

import (
	"time"

	"github.com/docker/go-units"
)

var MaxBodySize int64 = 1 * units.MiB
var MaxHeaderSize int64 = 8 * units.KiB
var MaxDelayTime = 5 * time.Minute
