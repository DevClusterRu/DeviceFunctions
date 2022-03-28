package metrics

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func (m *MetricsStructure) CountOpenFiles(instance string) {

	for {

		out, err := exec.Command("/bin/sh", "-c", fmt.Sprintf("lsof -p %v", os.Getpid())).Output()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		lines := strings.Split(string(out), "\n")
		m.MChannel <- Metric{fmt.Sprintf("cloud_descriptors_usage{instance=\"%s\"}", instance), float64(int64(len(lines) - 1))}
		time.Sleep(5 * time.Second)
	}

}
