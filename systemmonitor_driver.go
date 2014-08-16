package systemmonitor

import (
	"github.com/cloudfoundry/gosigar"
	"github.com/hybridgroup/gobot"
)

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

	s.AddCommand("GetCpu", func(params map[string]interface{}) interface{} {
		return resultApi(s.Cpu())
	})

	s.AddCommand("GetCpuList", func(params map[string]interface{}) interface{} {
		return resultApi(s.CpuList())
	})

	/*
		s.AddCommand("GetFileSystem", func(params map[string]interface{}) interface{} {
			return resultApi(s.FileSystem())
		})
	*/

	s.AddCommand("GetFileSystemList", func(params map[string]interface{}) interface{} {
		return resultApi(s.FileSystemList())
	})

	s.AddCommand("GetFileSystemUsage", func(params map[string]interface{}) interface{} {
		path := params["path"].(string)
		return resultApi(s.FileSystemUsage(path))
	})

	s.AddCommand("GetLoadAverage", func(params map[string]interface{}) interface{} {
		return resultApi(s.LoadAverage())
	})

	s.AddCommand("GethMem", func(params map[string]interface{}) interface{} {
		return resultApi(s.Mem())
	})

	s.AddCommand("GetProcArgs", func(params map[string]interface{}) interface{} {
		pid := params["pid"].(float64)
		return resultApi(s.ProcArgs(int(pid)))
	})

	s.AddCommand("GetProcExe", func(params map[string]interface{}) interface{} {
		pid := params["pid"].(float64)
		return resultApi(s.ProcExe(int(pid)))
	})

	s.AddCommand("GetProcList", func(params map[string]interface{}) interface{} {
		return resultApi(s.ProcList())
	})

	s.AddCommand("GetProcMem", func(params map[string]interface{}) interface{} {
		pid := params["pid"].(float64)
		return resultApi(s.ProcMem(int(pid)))
	})

	s.AddCommand("GetProcState", func(params map[string]interface{}) interface{} {
		pid := params["pid"].(float64)
		return resultApi(s.ProcState(int(pid)))
	})

	s.AddCommand("GetProcTime", func(params map[string]interface{}) interface{} {
		pid := params["pid"].(float64)
		return resultApi(s.ProcTime(int(pid)))
	})

	s.AddCommand("GetSwap", func(params map[string]interface{}) interface{} {
		return resultApi(s.Swap())
	})

	s.AddCommand("GetUptime", func(params map[string]interface{}) interface{} {
		return resultApi(s.Uptime())
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

func (s *SystemmonitorDriver) Cpu() (cpu sigar.Cpu, err error) {
	cpu = sigar.Cpu{}
	err = cpu.Get()

	return cpu, err
}

func (s *SystemmonitorDriver) CpuList() (cpuList sigar.CpuList, err error) {
	cpuList = sigar.CpuList{}
	err = cpuList.Get()

	return cpuList, err
}

/*
func (s *SystemmonitorDriver) FileSystem() (fileSystem sigar.FileSystem, err error) {
	fileSystem = sigar.FileSystem{}
	err = fileSystem.Get()

	return fileSystem, err
}
*/

func (s *SystemmonitorDriver) FileSystemList() (fileSystemList sigar.FileSystemList, err error) {
	fileSystemList = sigar.FileSystemList{}
	err = fileSystemList.Get()

	return fileSystemList, err
}

func (s *SystemmonitorDriver) FileSystemUsage(path string) (fileSystemUsage sigar.FileSystemUsage, err error) {
	fileSystemUsage = sigar.FileSystemUsage{}
	err = fileSystemUsage.Get(path)

	return fileSystemUsage, err
}

func (s *SystemmonitorDriver) LoadAverage() (loadAverage sigar.LoadAverage, err error) {
	loadAverage = sigar.LoadAverage{}
	err = loadAverage.Get()

	return loadAverage, err
}

func (s *SystemmonitorDriver) Mem() (mem sigar.Mem, err error) {
	mem = sigar.Mem{}
	err = mem.Get()

	return mem, err
}

func (s *SystemmonitorDriver) ProcArgs(pid int) (procArgs sigar.ProcArgs, err error) {
	procArgs = sigar.ProcArgs{}
	err = procArgs.Get(pid)

	return procArgs, err
}

func (s *SystemmonitorDriver) ProcExe(pid int) (procExe sigar.ProcExe, err error) {
	procExe = sigar.ProcExe{}
	err = procExe.Get(pid)

	return procExe, err
}

func (s *SystemmonitorDriver) ProcList() (procList sigar.ProcList, err error) {
	procList = sigar.ProcList{}
	err = procList.Get()

	return procList, err
}

func (s *SystemmonitorDriver) ProcMem(pid int) (procMem sigar.ProcMem, err error) {
	procMem = sigar.ProcMem{}
	err = procMem.Get(pid)

	return procMem, err
}

func (s *SystemmonitorDriver) ProcState(pid int) (procState sigar.ProcState, err error) {
	procState = sigar.ProcState{}
	err = procState.Get(pid)

	return procState, err
}

func (s *SystemmonitorDriver) ProcTime(pid int) (procTime sigar.ProcTime, err error) {
	procTime = sigar.ProcTime{}
	err = procTime.Get(pid)

	return procTime, err
}

func (s *SystemmonitorDriver) Swap() (swap sigar.Swap, err error) {
	swap = sigar.Swap{}
	err = swap.Get()

	return swap, err
}

func (s *SystemmonitorDriver) Uptime() (uptime sigar.Uptime, err error) {
	uptime = sigar.Uptime{}
	err = uptime.Get()

	return uptime, err
}

func resultApi(v interface{}, err error) interface{} {
	if err == nil {
		return struct {
			Result interface{} `json:"result"`
		}{v}
	} else {
		return struct {
			Result error `json:"result"`
		}{err}
	}
}
