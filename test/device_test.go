package test

import (
	"fmt"
	"github.com/DevClusterRu/DeviceFunctions/config"
	"github.com/DevClusterRu/DeviceFunctions/platforms/adb"
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
	d := adb.NewDevice("ce071717c8ae0b02047e")
	fmt.Println(d)
}
