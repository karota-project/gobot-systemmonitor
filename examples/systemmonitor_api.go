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

				path := "" // ex /dev/disk0s2
				pid := 19136

				if cpu, err := systemmonitorDriver.Cpu(); err == nil {
					fmt.Println("--- Cpu ---")
					fmt.Println("User: ", cpu.User)
					fmt.Println("Nice: ", cpu.Nice)
					fmt.Println("Sys: ", cpu.Sys)
					fmt.Println("Idle: ", cpu.Idle)
					fmt.Println("Wait: ", cpu.Wait)
					fmt.Println("Irq: ", cpu.Irq)
					fmt.Println("SoftIrq: ", cpu.SoftIrq)
					fmt.Println("Stolen: ", cpu.Stolen)
					fmt.Println()
				}

				if cpuList, err := systemmonitorDriver.CpuList(); err == nil {
					fmt.Println("--- CpuList ---")
					for i, v := range cpuList.List {
						fmt.Println("--- Cpu", i, " ---")
						fmt.Println("User: ", v.User)
						fmt.Println("Nice: ", v.Nice)
						fmt.Println("Sys: ", v.Sys)
						fmt.Println("Idle: ", v.Idle)
						fmt.Println("Wait: ", v.Wait)
						fmt.Println("Irq: ", v.Irq)
						fmt.Println("SoftIrq: ", v.SoftIrq)
						fmt.Println("Stolen: ", v.Stolen)
						fmt.Println()
					}
				}

				if fileSystemList, err := systemmonitorDriver.FileSystemList(); err == nil {
					fmt.Println("--- FileSystemList ---")
					for i, v := range fileSystemList.List {
						fmt.Println("--- FileSystem", i, " ---")
						fmt.Println("DirName: ", v.DirName)
						fmt.Println("DevName: ", v.DevName)
						fmt.Println("TypeName: ", v.TypeName)
						fmt.Println("SysTypeName: ", v.SysTypeName)
						fmt.Println("Options: ", v.Options)
						fmt.Println("Flags: ", v.Flags)
						fmt.Println()
					}
				}

				if fileSystemUsage, err := systemmonitorDriver.FileSystemUsage(path); err == nil {
					fmt.Println("--- FileSystemUsage ---")
					fmt.Println("Total: ", fileSystemUsage.Total)
					fmt.Println("Used: ", fileSystemUsage.Used)
					fmt.Println("Free: ", fileSystemUsage.Free)
					fmt.Println("Avail: ", fileSystemUsage.Avail)
					fmt.Println("Files: ", fileSystemUsage.Files)
					fmt.Println("FreeFiles: ", fileSystemUsage.FreeFiles)
					fmt.Println()
				}

				if loadAverage, err := systemmonitorDriver.LoadAverage(); err == nil {
					fmt.Println("--- LoadAverage ---")
					fmt.Println("One: ", loadAverage.One)
					fmt.Println("Five: ", loadAverage.Five)
					fmt.Println("Fifteen: ", loadAverage.Fifteen)
					fmt.Println()
				}

				if mem, err := systemmonitorDriver.Mem(); err == nil {
					fmt.Println("--- Mem ---")
					fmt.Println("Total: ", mem.Total)
					fmt.Println("Used: ", mem.Used)
					fmt.Println("Free: ", mem.Free)
					fmt.Println("ActualFree: ", mem.ActualFree)
					fmt.Println("ActualUsed: ", mem.ActualUsed)
					fmt.Println()
				}

				if procArgs, err := systemmonitorDriver.ProcArgs(pid); err == nil {
					fmt.Println("--- ProcArgs ---")
					if len(procArgs.List) > 0 {
						for i, v := range procArgs.List {
							fmt.Println(i, ".  ", v)
						}
					} else {
						fmt.Println("none")
					}
					fmt.Println()
				}

				if procExe, err := systemmonitorDriver.ProcExe(pid); err == nil {
					fmt.Println("--- ProcExe ---")
					fmt.Println("Name: ", procExe.Name)
					fmt.Println("Cwd: ", procExe.Cwd)
					fmt.Println("Root: ", procExe.Root)
					fmt.Println()
				}

				if procList, err := systemmonitorDriver.ProcList(); err == nil {
					fmt.Println("--- ProcList ---")
					fmt.Println(procList.List)
					fmt.Println()
				}

				if procMem, err := systemmonitorDriver.ProcMem(pid); err == nil {
					fmt.Println("--- ProcMem ---")
					fmt.Println("Size: ", procMem.Size)
					fmt.Println("Resident: ", procMem.Resident)
					fmt.Println("Share: ", procMem.Share)
					fmt.Println("MinorFaults: ", procMem.MinorFaults)
					fmt.Println("MajorFaults: ", procMem.MajorFaults)
					fmt.Println("PageFaults: ", procMem.PageFaults)
					fmt.Println()
				}

				if procState, err := systemmonitorDriver.ProcState(pid); err == nil {
					fmt.Println("--- ProcState ---")
					fmt.Println("Name: ", procState.Name)
					fmt.Println("State: ", procState.State)
					fmt.Println("Ppid: ", procState.Ppid)
					fmt.Println("Tty: ", procState.Tty)
					fmt.Println("Priority: ", procState.Priority)
					fmt.Println("Nice: ", procState.Nice)
					fmt.Println("Processor: ", procState.Processor)
					fmt.Println()
				}

				if procTime, err := systemmonitorDriver.ProcTime(pid); err == nil {
					fmt.Println("--- ProcTime ---")
					fmt.Println("StartTime: ", procTime.StartTime)
					fmt.Println("User: ", procTime.User)
					fmt.Println("Sys: ", procTime.Sys)
					fmt.Println("Total: ", procTime.Total)
					fmt.Println()
				}

				if swap, err := systemmonitorDriver.Swap(); err == nil {
					fmt.Println("--- Swap ---")
					fmt.Println("Total: ", swap.Total)
					fmt.Println("Used: ", swap.Used)
					fmt.Println("Free: ", swap.Free)
					fmt.Println()
				}

				if uptime, err := systemmonitorDriver.Uptime(); err == nil {
					fmt.Println("--- Uptime ---")
					fmt.Println("Length: ", uptime.Length)
					fmt.Println()
				}
			}))

	master.Start()
}
