package metrics

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type Metric struct {
	Name  string  //4 example temperature{"device=1"}
	Value float64 //Float value
}

type MetricsStructure struct {
	MChannel    chan Metric
	AChannel    chan Metric
	Metrics     map[string]float64
	Accumulator map[string]float64
	M           *sync.Mutex
}

func (mt *MetricsStructure) ShowMetrics(w http.ResponseWriter, r *http.Request) {
	mt.M.Lock()
	for k, v := range mt.Metrics {
		fmt.Fprintf(w, fmt.Sprintf("%s %f\n", k, v))
	}
	mt.M.Unlock()
}

func (mt *MetricsStructure) MetricsProcessor() {

	go func() {
		for {
			mt.M.Lock()
			for k, v := range mt.Accumulator {
				mt.Metrics[k] = v
				mt.Accumulator[k] = 0
			}
			mt.M.Unlock()
			time.Sleep(10 * time.Second)
		}
	}()

	go func() {
		for {
			select {
			case m := <-mt.MChannel:
				mt.M.Lock()
				mt.Metrics[m.Name] = m.Value
				mt.M.Unlock()
			case a := <-mt.AChannel:
				mt.M.Lock()
				mt.Accumulator[a.Name] += a.Value
				mt.M.Unlock()
			}
		}
	}()

}

func RunMetrics(port string) *MetricsStructure {

	metricsObject := MetricsStructure{}
	metricsObject.MChannel = make(chan Metric, 1000)
	metricsObject.AChannel = make(chan Metric, 1000)
	metricsObject.Metrics = make(map[string]float64)
	metricsObject.Accumulator = make(map[string]float64)
	metricsObject.M = &sync.Mutex{}

	go func() {
		http.HandleFunc("/metrics", metricsObject.ShowMetrics)
		log.Println("Starting webserver...")
		err := http.ListenAndServe(":"+port, nil)
		if err != nil {
			log.Fatal(err)
		}
	}()

	metricsObject.MetricsProcessor()
	return &metricsObject
}
