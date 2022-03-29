package test

import (
	"fmt"
	"github.com/DevClusterRu/DeviceFunctions/metrics"
	"testing"
	"time"
)

func TestMetrics(t *testing.T) {
	port := "7777"
	m := metrics.RunMetrics(port)
	m.MChannel <- metrics.Metric{"Hello", 1.234}
	time.Sleep(10 * time.Millisecond)
	fmt.Println(m.Metrics)
}
