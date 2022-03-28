package metrics

import (
	"fmt"
	"runtime"
	"time"
)

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func (m *MetricsStructure) GetAllocMem(instance string) uint64 {
	for {
		var mt runtime.MemStats
		runtime.ReadMemStats(&mt)
		m.MChannel <- Metric{fmt.Sprintf("device_MEM_usage{instance=\"%s\"}", instance), float64(bToMb(mt.Alloc))}
		time.Sleep(5 * time.Second)
	}
}
