package test

import (
	"fmt"
	"github.com/DevClusterRu/DeviceFunctions/config"
	"github.com/DevClusterRu/DeviceFunctions/metrics"
	"log"
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

func TestConfig(t *testing.T) {
	conf, err := config.NewConfig("../local.yml")
	if err != nil {
		log.Fatal("Conf broken: ", err)
	}

	fmt.Println(conf)
}

func TestDeviceCollection(t *testing.T) {
	conf, err := config.NewConfig("../local.yml")
	if err != nil {
		log.Fatal("Conf broken: ", err)
	}
	fmt.Println(conf.GetDevicesCollection())
}
