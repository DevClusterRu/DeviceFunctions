package test

import (
	"fmt"
	"github.com/DevClusterRu/DeviceFunctions/config"
	"log"
	"testing"
)

func TestConfig(t *testing.T) {
	conf, err := config.NewConfig("../local.yml")
	if err != nil {
		log.Fatal("Conf broken: ", err)
	}

	fmt.Println(conf)
}
