package config

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/DevClusterRu/DeviceFunctions/metrics"
	"github.com/melbahja/goph"
	"github.com/tarantool/go-tarantool"
	"golang.org/x/crypto/ssh"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"path"
	"strings"
	"sync"
)

type Config struct {
	TARANTOOL               string `yaml:"TARANTOOL"`
	TPASSWORD               string `yaml:"TPASSWORD"`
	TUSER                   string `yaml:"TUSER"`
	AWS_ACCESS_KEY_ID       string `yaml:"AWS_ACCESS_KEY_ID"`
	AWS_SECRET_ACCESS_KEY   string `yaml:"AWS_SECRET_ACCESS_KEY"`
	AWS_SCREEN_BUCKET       string `yaml:"AWS_SCREEN_BUCKET"`
	AWS_LOG_BUCKET          string `yaml:"AWS_LOG_BUCKET"`
	APP_URL                 string `yaml:"APP_URL"`
	PROMETHEUS_PORT         string `yaml:"PROMETHEUS_PORT"`
	CLOUD_NUMBER            uint64 `yaml:"CLOUD_NUMBER"`
	DELAY_BEFORE_SCREENSHOT int64  `yaml:"DELAY_BEFORE_SCREENSHOT"`
	CAMBRIONIX_URL          string `yaml:"CAMBRIONIX_URL"`
	MODE                    string `yaml:"MODE"`
	Ttool                   *tarantool.Connection
	IMEI_COLLECTION         map[string]string `yaml:"IMEI_COLLECTION"`
	S3BufferedChannel       chan struct{}
	SSHClient               *goph.Client
	SSHUser                 string `yaml:"SSHUser"`
	SSHPassword             string `yaml:"SSHPassword"`
	SSHAddr                 string `yaml:"SSHAddr"`
	SSHPort                 int    `yaml:"SSHPort"`
	SSHPath                 string `yaml:"SSHPath"`
	TELNYX_KEY              string `yaml:"TELNYX_KEY"`
	TELNYX_CONNECTION       string `yaml:"TELNYX_CONNECTION"`
	TELNYX_CLIENT           *http.Client
}

func NewConfig(fileName string) (c *Config, err error) {
	log.Printf("reading config from '%s'", fileName)
	if ext := path.Ext(fileName); ext != ".yaml" && ext != ".yml" {
		err = fmt.Errorf("invalid  file '%s' extenstion, expected 'yaml' or 'yml'", ext)
		return
	}

	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		err = fmt.Errorf("can't read file '%s'", fileName)
		return
	}

	conf := new(Config)
	if err = yaml.Unmarshal(file, conf); err != nil {
		err = fmt.Errorf("file %s yaml unmarshal error: %v", fileName, err)
	}

	if conf.MODE != "debug" {
		f, err := os.OpenFile("logfile.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		defer f.Close()
		log.SetOutput(f)
	}

	conf.Ttool, err = tarantool.Connect(conf.TARANTOOL, tarantool.Opts{
		User: conf.TUSER,
		Pass: conf.TPASSWORD,
	})
	if err != nil {
		log.Fatal("Connection ttool refused")
	}

	//Get Telynx API creds ready to go
	conf.TELNYX_CLIENT = &http.Client{}

	metricsObject := metrics.MetricsStructure{}
	metricsObject.MChannel = make(chan metrics.Metric, 1000)
	metricsObject.AChannel = make(chan metrics.Metric, 1000)
	metricsObject.Metrics = make(map[string]float64)
	metricsObject.Accumulator = make(map[string]float64)
	metricsObject.M = &sync.Mutex{}

	conf.S3BufferedChannel = make(chan struct{}, 1000)

	auth := goph.Password(conf.SSHPassword)
	if err != nil {
		log.Fatal(err)
	}

	conf.SSHClient, err = goph.NewConn(
		&goph.Config{
			User:     conf.SSHUser,
			Addr:     conf.SSHAddr,
			Port:     uint(conf.SSHPort),
			Auth:     auth,
			Callback: verifyHost,
		})
	if err != nil {
		log.Fatal(err)
	}

	return conf, err
}

func verifyHost(host string, remote net.Addr, key ssh.PublicKey) error {
	hostFound, err := goph.CheckKnownHost(host, remote, key, "")
	if hostFound && err != nil {
		return err
	}
	if hostFound && err == nil {
		return nil
	}
	if askIsHostTrusted(host, key) == false {
		return errors.New("you typed no, aborted!")
	}
	// Add the new host to known hosts file.
	return goph.AddKnownHost(host, remote, key, "")
}

func askIsHostTrusted(host string, key ssh.PublicKey) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Unknown Host: %s \nFingerprint: %s \n", host, ssh.FingerprintSHA256(key))
	fmt.Print("Would you like to add it? type yes or no: ")
	a, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return strings.ToLower(strings.TrimSpace(a)) == "yes"
}
