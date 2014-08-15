package systemmonitor

import "github.com/hybridgroup/gobot"

type SystemmonitorDriver struct {
	gobot.Driver
}

type SystemmonitorInterface interface {
}

func NewSystemmonitorDriver(a *SystemmonitorAdaptor, name string) *SystemmonitorDriver {
	s := &SystemmonitorDriver{
		Driver: *gobot.NewDriver(
			name,
			"systemmonitor.SystemmonitorDriver",
			a,
		),
	}

	s.AddCommand("SystemInfo", func(params map[string]interface{}) interface{} {
		return resultSystemInfos(s.SystemInfo())
	})

	return s
}

func (s *SystemmonitorDriver) adaptor() *SystemmonitorAdaptor {
	return s.Driver.Adaptor().(*SystemmonitorAdaptor)
}

func (s *SystemmonitorDriver) Start() bool {
	return true
}

func (s *SystemmonitorDriver) Halt() bool {
	return true
}

func (s *SystemmonitorDriver) SystemInfo() (sysInfo []*SysInfo, err error) {
	return getSystemInfo()
}

func resultSystemInfos(s []SystemInfo, err error) interface{} {
	if err == nil {
		return struct {
			Result []SystemInfo `json:"result"`
		}{s}
	} else {
		return struct {
			Result error `json:"result"`
		}{err}
	}
}
