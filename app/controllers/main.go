package controllers

import (
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"wmqx-ui/app"
	"github.com/astaxie/beego"
	"encoding/json"
	"wmqx-ui/app/models"
)

type MainController struct {
	BaseController
}

func (this *MainController) Index() {

	// get my node
	nodes := []map[string]string{}
	if this.roleIsUser() {
		userNodes, err := models.UserNodeModel.GetUserNodeByUserId(this.UserID)
		if err == nil {
			nodeIds := []string{}
			for _, userNode := range userNodes {
				nodeIds = append(nodeIds, userNode["node_id"])
			}
			nodes, _ = models.NodeModel.GetNodesByNodeIds(nodeIds)
		}
	} else {
		nodes, _ = models.NodeModel.GetNodes()
	}

	// get contacts
	contacts := []map[string]string{}
	contactsConf, err := beego.AppConfig.GetSection("contact")
	if err == nil {
		for name, contactConf := range contactsConf {
			if contactConf == "" {
				continue
			}
			contact := map[string]string{}
			err := json.Unmarshal([]byte(contactConf), &contact)
			if err != nil {
				continue
			}
			contact["name"]=name
			contacts = append(contacts, contact)
		}
	}

	// my system log
	var logs []map[string]string
	logs, _ = models.LogModel.GetLogsByKeywordAndLimit("", "", this.User["username"], 0, 15)

	this.Data["version"] = app.Version
	this.Data["nodes"] = nodes
	this.Data["logs"] = logs
	this.Data["contacts"] = contacts
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