package adb

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"time"
)

func GoBash(command string, args ...string) []byte {
	var bt []byte
	cmd := exec.Command(command, args...)
	var stdout io.ReadCloser
	//var stderr io.ReadCloser
	c := 0
	for {
		var err1 error
		var err2 error
		stdout, err1 = cmd.StdoutPipe()
		//stderr, err2 = cmd.StderrPipe()

		if err1 == nil && err2 == nil {
			break
			//log.Fatal("ERROR stdout: ", err1, err2)
		}
		c++
		if c >= 5 {
			return nil
		}
		fmt.Println(" STDOUT ERROR, retrying... ", c, err1, err2)
		time.Sleep(1 * time.Second)
	}

	scanOut := bufio.NewScanner(stdout)
	//scanErr = bufio.NewScanner(stderr)

	cmd.Start()

	for scanOut.Scan() {
		m := scanOut.Bytes()
		bt = append(bt, m...)
	}

	//for scanner.Scan() {
	//	m := scanner.Text()
	//	sb.WriteString(m + "\n")
	//}
	cmd.Wait()

	scanOut = nil
	return bt
}
