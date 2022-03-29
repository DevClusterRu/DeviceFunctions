package test

import (
	"fmt"
	"github.com/DevClusterRu/DeviceFunctions/config"
	"github.com/DevClusterRu/DeviceFunctions/platforms"
	"log"
	"testing"
)

func TestDeviceCollection(t *testing.T) {
	conf, err := config.NewConfig("../local.yml")
	if err != nil {
		log.Fatal("Conf broken: ", err)
	}
	fmt.Println(conf.GetDevicesCollection())
}

func TestNewDevice(t *testing.T) {
	d := platforms.NewDevice("ce071717c8ae0b02047e", "android")
	fmt.Println(d)
}
