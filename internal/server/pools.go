package server

import (
	"bytes"
	"sync"
	"time"

	"github.com/valyala/fasthttp"
)

var timerPool = sync.Pool{
	New: func() any {
		return time.NewTimer(-1)
	},
}

var bufferPool = sync.Pool{
	New: func() any {
		return new(bytes.Buffer)
	},
}

var cookieExpire = time.Now().Add(365 * 24 * time.Hour)

var cookiePool = sync.Pool{
	New: func() any {
		c := fasthttp.Cookie{}
		c.SetKey("whoami")
		c.SetExpire(cookieExpire)
		c.SetPath("/")
		return &c
	},
}
