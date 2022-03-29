package adb

import (
	"github.com/DevClusterRu/DeviceFunctions/config"
	"github.com/DevClusterRu/DeviceFunctions/platforms"
	"regexp"
	"strings"
	"sync"
	"time"
)

func getPhoneNumber(serial string) string {
	codes := []string{"17", "15"}
	for _, v := range codes {
		numberStr := string(GoBash("adb", "-s", serial, "shell", "service call iphonesubinfo "+v+" |awk -F \"'\" '{print $2}'|sed '1 d'|tr -d '.'|awk '{print}' ORS=|awk '{print $1}'"))
		r := regexp.MustCompile(`\d{10,}`)
		number := r.FindString(numberStr)
		if len(number) == 10 || len(number) == 11 {
			if len(number) == 10 {
				number = "+1" + number
			}
			if len(number) == 11 {
				number = "+" + number
			}
			return number
		}
	}
	return ""
}

func getImei(serial string) string {
	imeiStr := string(GoBash("adb", "-s", serial, "shell", "service call iphonesubinfo 1 |awk -F \"'\" '{print $2}'|sed '1 d'|tr -d '.'|awk '{print}' ORS=|awk '{print $1}'"))
	//imeiStr := GoBash("./imei.sh")
	r := regexp.MustCompile(`\d{10,}`)
	return r.FindString(imeiStr)
}

func getCarrier(serial string) string {
	return carrierRewrite(string(GoBash("adb", "-s", serial, "shell", "getprop gsm.operator.alpha")))
}

func getModel(serial string) string {
	return strings.TrimSpace(string(GoBash("adb", "-s", serial, "shell", "getprop ro.product.model")))
}

func GetValuesList(serial string) map[string]string {
	m := make(map[string]string)
	m["deviceNumber"] = getPhoneNumber(serial)
	m["carrier"] = getCarrier(serial)
	m["model"] = getModel(serial)
	return m
}

func carrierRewrite(str string) string {
	strUp := strings.ToUpper(str)
	if strings.Contains(strUp, "TMOBILE") || strings.Contains(strUp, "T-MOBILE") {
		return "tmobile"
	}
	if strings.Contains(strUp, "VERIZON") {
		return "verizon"
	}
	if strings.Contains(strUp, "ATT") || strings.Contains(strUp, "AT&T") {
		return "att"
	}
	if strings.Contains(strUp, "TING") {
		return "ting"
	}
	return str
}

func IsOnline(serial string) bool {
	answer := string(GoBash("adb", "-s", serial, "shell", "date"))
	if strings.TrimSpace(answer) != "" {
		return true
	}
	return false
}

func NewDevice(serial string, c config.Config) *platforms.Device {

	if !IsOnline(serial) {
		return nil
	}

	dev := &platforms.Device{
		SerialNumber: serial,
		Platform:     "android",
		SilentFlag:   false,
		SilentFlagAt: 0,
		Ringing:      false,
		SilentScore:  time.Now(),
		Cnam:         make(map[string]int),
		Mtx:          new(sync.Mutex),
	}

	properties := GetValuesList(serial)

	if number, ok := properties["deviceNumber"]; ok {
		dev.DeviceNumber = number
	}

	if model, ok := properties["model"]; ok {
		dev.Product = model
	}

	if carrier, ok := properties["carrier"]; ok {
		dev.Carrier = carrier
	}

	return dev
}
