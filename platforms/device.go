package platforms

import (
	"context"
	"github.com/DevClusterRu/DeviceFunctions/platforms/adb"
	"os"
	"sync"
	"time"
)

type Device struct {
	Id                 uint64
	SilentFlag         bool
	SilentFlagAt       int
	Status             string
	SerialNumber       string
	DeviceId           uint64
	DeviceStatus       string
	IncomingNumber     string
	DeviceNumber       string
	Session            string
	Platform           string
	Carrier            string
	Hiya               bool
	Product            string
	JobId              uint64
	SilentScore        time.Time
	TestLog            bool
	LogFile            *os.File
	Cnam               map[string]int //Catch CNAM here
	MessageId          string
	Ringing            bool      //if true - ringing right now!
	RingingTime        time.Time //Time ringing start
	ContextCancel      context.CancelFunc
	Mtx                *sync.Mutex
	RebootCase         uint64
	SuccessCallCounter uint64
	FailCallCounter    uint64
}

func NewDevice(serial string, platform string) *Device {

	dev := &Device{
		SerialNumber: serial,
		Platform:     platform,
		SilentFlag:   false,
		SilentFlagAt: 0,
		Ringing:      false,
		SilentScore:  time.Now(),
		Cnam:         make(map[string]int),
		Mtx:          new(sync.Mutex),
	}

	switch platform {
	case "android":
		dev.DeviceNumber = adb.GetPhoneNumber(serial)
		dev.Product = adb.GetModel(serial)
		dev.Carrier = adb.GetCarrier(serial)
		dev.DeviceNumber = adb.GetPhoneNumber(serial)
	}

	return dev
}
