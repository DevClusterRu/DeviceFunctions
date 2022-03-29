package platforms

import (
	"context"
	"os"
	"sync"
	"time"
)

type Device struct {
	Id             uint64
	SilentFlag     bool
	SilentFlagAt   int
	Status         string
	Ttl            uint64
	Attempts       uint64
	SerialNumber   string
	DeviceId       uint64
	DeviceStatus   string
	IncomingNumber string
	DeviceNumber   string
	Session        string
	Platform       string
	Carrier        string
	Hiya           bool
	Product        string
	JobId          uint64
	Temperature    float64
	Volts          float64
	SilentScore    time.Time
	TestLog        bool
	LogFile        *os.File
	Cnam           map[string]int //Catch CNAM here
	MessageId      string
	Ringing        bool      //if true - ringing right now!
	RingingTime    time.Time //Time ringing start
	ContextCancel  context.CancelFunc
	Mtx            *sync.Mutex
	RebootCase     uint64
}
