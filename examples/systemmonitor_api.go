package main

import (
	"fmt"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/api"
	"github.com/karota-project/gobot-systemmonitor"
)

func main() {
	master := gobot.NewGobot()
	api.NewAPI(master).Start()

	systemmonitorAdaptor := systemmonitor.NewSystemmonitorAdaptor("systemmonitor-a01")
	systemmonitorDriver := systemmonitor.NewSystemmonitorDriver(systemmonitorAdaptor, "systemmonitor-d01")

	master.AddRobot(
		gobot.NewRobot(
			"systemmonitor",
			[]gobot.Connection{systemmonitorAdaptor},
			[]gobot.Device{systemmonitorDriver},
			func() {
				fmt.Println("work")

				systemInfo, err := systemmonitorDriver.SystemInfo()
				if err == nil {
					for _, s := range systemInfo {
						fmt.Println("mem_used: ", s.Memory.Swapd)
						fmt.Println("mem_free: ", s.Memory.Free)
						fmt.Println("cpu_used: ", s.Cpu.Us)
						fmt.Println("----------")
					}
				}
			}))

	master.Start()
}
