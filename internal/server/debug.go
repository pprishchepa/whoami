package server

import (
	"context"
	"fmt"
	"runtime"
	"time"

	"github.com/rs/zerolog/log"
)

func printMemUsage(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	var m runtime.MemStats
	for {
		runtime.ReadMemStats(&m)
		log.Debug().Msg(fmt.Sprint(
			fmt.Sprintf("Alloc = %v MiB", m.Alloc/1024/1024),
			fmt.Sprintf("\tTotalAlloc = %v MiB", m.TotalAlloc/1024/1024),
			fmt.Sprintf("\tSys = %v MiB", m.Sys/1024/1024),
			fmt.Sprintf("\tNumGC = %v", m.NumGC),
		))
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
		}
	}
}
