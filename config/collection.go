package config

import (
	"log"
	"strconv"
	"strings"
)

type Collection struct {
	Platform  string
	DeviceId  int
	HubSerial string
	Port      int
	Number    string
}

func (c *Config) GetDevicesCollection() map[string]Collection {
	collection := make(map[string]Collection)
	for k, v := range c.DEVICES_COLLECTION {
		elements := strings.Split(v, "-")
		if len(elements) < 5 {
			log.Println("Wrong element count: ", v)
			continue
		}
		deviceId, err := strconv.Atoi(elements[1])
		if err != nil {
			log.Println("Wrong numeric: ", elements[1])
			continue
		}
		devicePort, err := strconv.Atoi(elements[3])
		if err != nil {
			log.Println("Wrong numeric: ", elements[1])
			continue
		}

		collection[k] = Collection{
			Platform:  elements[0],
			DeviceId:  deviceId,
			HubSerial: elements[2],
			Port:      devicePort,
			Number:    elements[4],
		}
	}

	return collection
}
