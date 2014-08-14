package systemmonitor

import (
	"log"
	"log/syslog"
	"os"
	"runtime"

	"text/template"
	"time"

	"github.com/hybridgroup/gobot"
)

type SystemmonitorDriver struct {
	gobot.Driver
}

type SystemmonitorInterface interface {
}

func NewSystemmonitorDriver(a *SystemmonitorAdaptor, name string) *SystemmonitorDriver {
	return &SystemmonitorDriver{
		Driver: *gobot.NewDriver(
			name,
			"systemmonitor.SystemmonitorDriver",
			a,
		),
	}
}

func (s *SystemmonitorDriver) adaptor() *SystemmonitorAdaptor {
	return s.Driver.Adaptor().(*SystemmonitorAdaptor)
}

func (s *SystemmonitorDriver) Start() bool { return true }

func (s *SystemmonitorDriver) Halt() bool { return true }

// Printlog for linux
func (s *SystemmonitorDriver) Stdout(functionName string) (err error) {
	_, file, line, _ := runtime.Caller(1)

	sysInfo, err := getSystemInfo()
	if err != nil {
		return err
	}

	for _, s := range sysInfo {
		v, err := logLine(&SysInfoWithFunc{s, functionName})
		if err != nil {
			return err
		}

		t, err := template.New("template").Parse(`{{.File}} (line {{.Line}}) {{.Json}}`)
		if err != nil {
			return err
		}

		t.Execute(os.Stdout, struct {
			File string
			Line int
			Json string
		}{file, line, string(v)})
	}

	return nil
}

// Write file
func (s *SystemmonitorDriver) WriteFile(functionName string, outputfile string) (err error) {
	_, file, line, _ := runtime.Caller(1)

	f, err := os.OpenFile(outputfile, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		return err
	}

	sysInfo, err := getSystemInfo()
	if err != nil {
		return err
	}

	for _, s := range sysInfo {
		v, err := logLine(&SysInfoWithFunc{s, functionName})
		if err != nil {
			return err
		}

		t, err := template.New("template").Parse(`[{{.Date}}] {{.File}} (line {{.Line}}) {{.Json}}`)
		if err != nil {
			return err
		}

		t.Execute(f, struct {
			Date time.Time
			File string
			Line int
			Json string
		}{time.Now(), file, line, string(v)})
	}

	defer f.Close()

	return nil
}

// Print syslog for unix
func (s *SystemmonitorDriver) Syslog(priority Priority, facility string) (err error) {
	p := convToSyslogPriority(priority)

	// Configure logger to write to the syslog. You could do this in init(), too.
	logWriter, err := syslog.New(p, facility)
	if err != nil {
		return err
	}

	logger := log.New(logWriter, "", log.LstdFlags)

	sysInfo, err := getSystemInfo()
	if err != nil {
		return err
	}

	for _, s := range sysInfo {
		v, err := logLine(s)

		if err != nil {
			return err
		}

		logger.Print(string(v))
	}

	return nil
}
