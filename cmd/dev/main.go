package main

import (
	"runtime"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/zephinzer/dev/internal/constants"
	"github.com/zephinzer/dev/internal/log"
)

func main() {
	timeStarted := time.Now()
	go func(ticker <-chan time.Time) {
		for {
			<-ticker
			if log.Trace != nil {
				var memoryStatistics runtime.MemStats
				runtime.ReadMemStats(&memoryStatistics)
				log.Tracef("profile as of %s (uptime: %v)", time.Now().Format(constants.DevTimeFormat), humanize.Time(timeStarted))
				log.Tracef("goroutines  : %v", runtime.NumGoroutine())
				totalAlloc := humanize.Bytes(memoryStatistics.TotalAlloc)
				totalFrees := humanize.Bytes(memoryStatistics.Frees)
				heapAlloc := humanize.Bytes(memoryStatistics.HeapAlloc)
				sys := humanize.Bytes(memoryStatistics.Sys)
				log.Tracef("alloc/frees/heap/sys : %s / %s / %s / %s", totalAlloc, totalFrees, heapAlloc, sys)
			}
		}
	}(time.Tick(time.Second * 5))
	cmd := GetCommand()
	cmd.Execute()
}
