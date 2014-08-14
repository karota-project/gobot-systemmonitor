package systemmonitor

import (
	"github.com/hybridgroup/gobot"
	"testing"
)

func initTestSystemmonitorDriver() *SystemmonitorDriver {
	return NewSystemmonitorDriver(NewSystemmonitorAdaptor("myAdaptor"), "myDriver")
}

func TestSystemmonitorDriverStart(t *testing.T) {
	d := initTestSystemmonitorDriver()
	gobot.Expect(t, d.Start(), true)
}

func TestSystemmonitorDriverHalt(t *testing.T) {
	d := initTestSystemmonitorDriver()
	gobot.Expect(t, d.Halt(), true)
}
