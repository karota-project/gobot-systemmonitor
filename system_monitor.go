package systemmonitor

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

type SysInfo struct {
	MemUsed int     `json:"mem_used"`
	MemFree int     `json:"mem_free"`
	CpuUsed float64 `json:"cpu_used"`
}

// exec vmstat command
func getSystemInfo() (sysInfo []*SysInfo, err error) {
	cmd := exec.Command("vmstat")
	stdout := bytes.Buffer{}
	stderr := bytes.Buffer{}

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
	}

	sysInfo = make([]*SysInfo, 0)
	for {
		line, err := stdout.ReadString('\n')
		if err != nil {
			break
		}

		// split string into array
		ft := make([]string, 0)
		tokens := strings.Split(line, " ")
		for _, t := range tokens {
			if t != "" && t != "\t" {
				ft = append(ft, t)
			}
		}

		// memUsed : swapd + buffer + cached
		swap, err := strconv.Atoi(ft[2])
		buf, err := strconv.Atoi(ft[4])
		cach, err := strconv.Atoi(ft[5])
		memUsed := swap + buf + cach
		if err != nil {
			continue
		}

		memFree, err := strconv.Atoi(ft[3])
		if err != nil {
			continue
		}

		cpuUsed, err := strconv.ParseFloat(ft[12], 64)
		if err != nil {
			continue
		}

		sysInfo = append(sysInfo, &SysInfo{memUsed, memFree, cpuUsed})
	}

	return sysInfo, err
}
