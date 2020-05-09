package main

import (
	"runtime"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/usvc/dev/internal/constants"
	"github.com/usvc/dev/internal/log"
)

func main() {
	go func(ticker <-chan time.Time) {
		for {
			<-ticker
			if log.Trace != nil {
				var memoryStatistics runtime.MemStats
				runtime.ReadMemStats(&memoryStatistics)
				log.Tracef("profile as of %s", time.Now().Format(constants.DevTimeFormat))
				log.Tracef("goroutines  : %v", runtime.NumGoroutine())
				log.Tracef("alloc/total : %v / %v", humanize.Bytes(memoryStatistics.TotalAlloc), humanize.Bytes(memoryStatistics.Sys))
			}
		}
	}(time.Tick(time.Second * 5))
	cmd := GetCommand()
	cmd.Execute()
}
