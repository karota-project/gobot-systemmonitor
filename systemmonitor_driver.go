package systemmonitor

import (
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
func (s *SystemmonitorDriver) Halt() bool  { return true }
