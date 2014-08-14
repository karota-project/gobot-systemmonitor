package systemmonitor

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/syslog"
	"os/exec"
	"strconv"
	"strings"
)

const (
	// From /usr/include/sys/syslog.h.
	// These are the same on Linux, BSD, and OS X.
	LOG_EMERG Priority = iota
	LOG_ALERT
	LOG_CRIT
	LOG_ERR
	LOG_WARNING
	LOG_NOTICE
	LOG_INFO
	LOG_DEBUG
)

type SysInfo struct {
	MemUsed int     `json:"mem_used"`
	MemFree int     `json:"mem_free"`
	CpuUsed float64 `json:"cpu_used"`
}

type SysInfoWithFunc struct {
	*SysInfo
	Func string `json:"func"`
}

type Priority int

func convToSyslogPriority(p Priority) (priority syslog.Priority) {
	switch p {
	case LOG_EMERG:
		return syslog.LOG_EMERG

	case LOG_ALERT:
		return syslog.LOG_ALERT

	case LOG_CRIT:
		return syslog.LOG_CRIT

	case LOG_ERR:
		return syslog.LOG_ERR

	case LOG_WARNING:
		return syslog.LOG_WARNING

	case LOG_NOTICE:
		return syslog.LOG_NOTICE

	case LOG_INFO:
		return syslog.LOG_INFO

	case LOG_DEBUG:
		return syslog.LOG_DEBUG

	default:
		return syslog.LOG_DEBUG
	}
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

func logLine(sysInfo interface{}) ([]byte, error) {
	return json.Marshal(sysInfo)
}
