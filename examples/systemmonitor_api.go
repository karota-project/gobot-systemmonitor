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

				err := systemmonitorDriver.Stdout("roomba2d2")
				if err != nil {
					fmt.Println(err)
				}
				err = systemmonitorDriver.WriteFile("roomba2d2", "sample.log")
				if err != nil {
					fmt.Println(err)
				}

				// on debian
				// $ cat /var/log/syslog | grep roomba2d2
				err = systemmonitorDriver.Syslog(systemmonitor.LOG_NOTICE, "roomba2d2")
				if err != nil {
					fmt.Println(err)
				}
			}))

	master.Start()
}
