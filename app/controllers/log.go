package controllers

import (
	"wmqx-ui/app/models"
	"strings"
	"wmqx-ui/app/remotes"
)

type LogController struct {
	BaseController
}

func (this *LogController) System() {

	page, _ := this.GetInt("page", 1)
	level := strings.Trim(this.GetString("level", ""), "")
	message := strings.Trim(this.GetString("message", ""), "")
	username := strings.Trim(this.GetString("username", ""), "")

	number := 15
	limit := (page - 1) * number
	var err error
	var count int64
	var logs []map[string]string
	if level != "" || message != "" || username != "" {
		count, err = models.LogModel.CountLogsByKeyword(level, message, username)
		logs, err = models.LogModel.GetLogsByKeywordAndLimit(level, message, username, limit, number)
	} else {
		count, err = models.LogModel.CountLogs()
		logs, err = models.LogModel.GetLogsByLimit(limit, number)
	}
	if err != nil {
		this.viewError(err.Error(), "template")
	}

	this.Data["logs"] = logs
	this.Data["username"] = username
	this.Data["level"] = level
	this.Data["message"] = message
	this.SetPaginator(number, count)
	this.viewLayout("log/system", "template")
}

func (this *LogController) Info() {

	logId := this.GetString("log_id", "")
	if logId == "" {
		this.viewError("日志不存在", "default")
	}

	log, err := models.LogModel.GetLogByLogId(logId)
	if err != nil {
		this.viewError("日志不存在", "default")
	}


	this.Data["log"] = log

	this.viewLayout("log/info", "default")
}

func (this *LogController) Node() {

	nodeId := this.GetString("node_id", "")
	keyword := this.GetString("keyword", "")
	logType := this.GetString("type", "error")

	nodes := []map[string]string{}
	var err error
	if this.roleIsUser() {
		userNodes, err := models.UserNodeModel.GetUserNodeByUserId(this.UserID)
		if err != nil {
			this.viewError("查找节点失败", "template")
		}

		nodeIds := []string{}
		for _, userNode := range userNodes {
			nodeIds = append(nodeIds, userNode["node_id"])
		}
		nodes, err = models.NodeModel.GetNodesByNodeIds(nodeIds)
		if err != nil {
			this.viewError("查找节点失败", "template")
		}
		if nodeId == "" && len(nodeIds) > 0 {
			nodeId = nodeIds[0]
		}
	} else {
		nodes, err = models.NodeModel.GetNodes()
		if err != nil {
			this.viewError("查找节点失败", "template")
		}
		if nodeId == "" && len(nodes) > 0 {
			nodeId = nodes[0]["node_id"]
		}
	}

	defaultNode := map[string]string{}
	for _, node := range nodes {
		if node["node_id"] == nodeId {
			defaultNode = node
			break
		}
	}
	if len(defaultNode) == 0 {
		this.viewError("没有选择节点", "template")
	}

	logs, err := remotes.NewLogByNode(defaultNode).Index()
	if err != nil {
		this.ErrorLog("获取节点日志失败："+err.Error())
		this.viewError("获取节点日志失败", "default")
	}

	this.Data["nodes"] = nodes
	this.Data["node_id"] = nodeId
	this.Data["log_type"] = logType
	this.Data["keyword"] = keyword
	this.Data["logs"] = logs

	this.viewLayout("log/node", "template")
}

func (this *LogController) Download()  {

	//nodeId, _ := this.GetInt("node_id", 1);
	//
	//_, logDownloads := models.LogDownload(nodeId)
	//
	//this.Data["logDownloads"] = logDownloads;
	//
	//this.layoutHtml = "layout/template";
	//
	//this.display("log/download");
}