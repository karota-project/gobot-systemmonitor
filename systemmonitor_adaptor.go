package systemmonitor

import (
  "github.com/hybridgroup/gobot"
)

type SystemmonitorAdaptor struct {
  gobot.Adaptor
}

func NewSystemmonitorAdaptor(name string) *SystemmonitorAdaptor {
  return &SystemmonitorAdaptor{
    Adaptor: *gobot.NewAdaptor(
      name,
      "systemmonitor.SystemmonitorAdaptor",
    ),
  }
}

func (s *SystemmonitorAdaptor) Connect() bool {
  return true
}

func (s *SystemmonitorAdaptor) Finalize() bool {
  return true
}
