package controllers

import (
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"wmqx-ui/app"
)

type MainController struct {
	BaseController
}

func (this *MainController) Index() {

	this.Data["version"] = app.Version
	this.viewTemplate("main/index")
}

// ajax 获取服务器状态
func (this *MainController) ServerStatus() {
	vm, _ := mem.VirtualMemory()
	cpuPercent, _ := cpu.Percent(0, false)
	d, _ := disk.Usage("/")

	data := map[string]interface{}{
		"memory_used_percent": int(vm.UsedPercent),
		"cpu_used_percent":    int(cpuPercent[0]),
		"disk_used_percent":   int(d.UsedPercent),
	}

	this.jsonSuccess("ok", data)
}