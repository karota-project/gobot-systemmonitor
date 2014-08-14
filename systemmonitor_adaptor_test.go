package systemmonitor

import (
	"github.com/hybridgroup/gobot"
	"testing"
)

func initTestSystemmonitorAdaptor() *SystemmonitorAdaptor {
	return NewSystemmonitorAdaptor("myAdaptor")
}

func TestSystemmonitorAdaptorConnect(t *testing.T) {
	a := initTestSystemmonitorAdaptor()
	gobot.Expect(t, a.Connect(), true)
}

func TestSystemmonitorAdaptorFinalize(t *testing.T) {
	a := initTestSystemmonitorAdaptor()
	gobot.Expect(t, a.Finalize(), true)
}
