package controllers

import (
	"strings"
	"wmqx-ui/app/models"
	"time"
	"wmqx-ui/app/remotes"
)

type NodeController struct {
	BaseController
}

func (this *NodeController) Add() {
	this.viewTemplate("node/form")
}

func (this *NodeController) List() {

	page, _ := this.GetInt("page", 1)
	comment := strings.TrimSpace(this.GetString("keyword", ""))
	keywords := map[string]string{
		"comment": comment,
	}

	number := 20
	limit := (page - 1) * number
	var err error
	var count int64
	var nodes []map[string]string
	if keywords["comment"] != "" {
		count, err = models.NodeModel.CountNodesByKeywords(keywords)
		nodes, err = models.NodeModel.GetNodesByKeywordsAndLimit(keywords, limit, number)
	} else {
		count, err = models.NodeModel.CountNodes()
		nodes, err = models.NodeModel.GetNodesByLimit(limit, number)
	}
	if err != nil {
		this.ErrorLog("获取节点失败："+err.Error())
		this.viewError(err.Error(), "/node/list")
	}
	this.Data["nodes"] = nodes
	this.Data["comment"] = comment
	this.SetPaginator(number, count)

	this.viewTemplate("node/list")
}

func (this *NodeController) Save() {

	managerURI := strings.TrimSpace(this.GetString("manager_uri"))
	publishURI := strings.TrimSpace(this.GetString("publish_uri"))
	token := strings.TrimSpace(this.GetString("token"))
	tokenHeaderName := strings.TrimSpace(this.GetString("token_header_name"))
	comment := strings.TrimSpace(this.GetString("comment"))

	if managerURI == "" {
		this.jsonError("管理URI不能为空")
	}
	if publishURI == "" {
		this.jsonError("发布URI不能为空")
	}
	if token == "" {
		this.jsonError("Token 不能为空")
	}
	if tokenHeaderName == "" {
		this.jsonError("TokenHeader 不能为空")
	}
	if comment ==  "" {
		this.jsonError("备注不能为空")
	}

	nodeValue := map[string]interface{}{
		"manager_uri":      strings.TrimRight(managerURI, "/"),
		"publish_uri":      strings.TrimRight(publishURI, "/"),
		"token":            token,
		"token_header_name":tokenHeaderName,
		"comment":          comment,
		"is_delete":        models.NODE_NORMAL,
		"create_time":      time.Now().Unix(),
		"update_time":      time.Now().Unix(),
	}
	_, err := models.NodeModel.Insert(nodeValue)
	if err != nil {
		this.ErrorLog("添加节点失败: "+err.Error())
		this.jsonError("添加节点失败！")
	}
	this.InfoLog("添加节点成功")
	this.jsonSuccess("添加节点成功", nil, "/node/list")
}

func (this *NodeController) Edit() {

	nodeId := this.GetString("node_id", "")
	if nodeId == "" {
		this.viewError("节点不存在", "/node/list")
	}

	node, err := models.NodeModel.GetNodeByNodeId(nodeId)
	if err != nil {
		this.ErrorLog("查找节点 "+nodeId+" 失败: "+err.Error())
		this.viewError("节点不存在", "/node/list")
	}

	this.Data["node"] = node
	this.viewLayout("node/edit", "default")
}

func (this *NodeController) Modify() {

	nodeId := strings.TrimSpace(this.GetString("node_id", ""))
	managerURI := strings.TrimSpace(this.GetString("manager_uri"))
	publishURI := strings.TrimSpace(this.GetString("publish_uri"))
	tokenHeaderName := strings.TrimSpace(this.GetString("token_header_name"))
	token := strings.TrimSpace(this.GetString("token"))
	comment := strings.TrimSpace(this.GetString("comment"))
	if nodeId == "" {
		this.viewError("节点不存在", "/node/list")
	}
	if managerURI == "" {
		this.jsonError("管理URI不能为空")
	}
	if publishURI == "" {
		this.jsonError("发布URI不能为空")
	}
	if tokenHeaderName == "" {
		this.jsonError("TokenHeader 不能为空")
	}
	if token == "" {
		this.jsonError("Token 不能为空")
	}
	if comment ==  "" {
		this.jsonError("备注不能为空")
	}

	nodeValue := map[string]interface{}{
		"manager_uri":      managerURI,
		"publish_uri":      publishURI,
		"token":            token,
		"token_header_name":tokenHeaderName,
		"comment":          comment,
		"is_delete":        models.NODE_NORMAL,
		"update_time":      time.Now().Unix(),
	}

	_, err := models.NodeModel.Update(nodeId, nodeValue)
	if err != nil {
		this.ErrorLog("修改节点 "+nodeId+" 失败: "+err.Error())
		this.jsonError("修改节点失败！")
	}
	this.InfoLog("修改节点 "+nodeId+" 成功")
	this.jsonSuccess("修改节点成功", nil, "/node/list")
}

func (this *NodeController) Delete() {

	nodeId := this.GetString("node_id", "")
	if nodeId == "" {
		this.jsonError("节点不存在！")
	}

	node, err := models.NodeModel.GetNodeByNodeId(nodeId)
	if err != nil || len(node) == 0 {
		this.jsonError("节点不存在！")
	}

	nodeValue := map[string]interface{}{
		"is_delete":   models.NODE_DELETE,
		"update_time": time.Now().Unix(),
	}

	_, err = models.NodeModel.Update(nodeId, nodeValue)
	if err != nil {
		this.ErrorLog("删除节点 "+nodeId+" 失败: "+err.Error())
		this.jsonError("删除节点失败！")
	}
	this.InfoLog("删除节点 "+nodeId+" 成功")
	this.jsonSuccess("删除节点成功", nil, "/node/list")
}

func (this *NodeController) Message() {

	nodeId := this.GetString("node_id", "")
	if nodeId == "" {
		this.viewError("节点不存在！", "default")
	}

	node, err := models.NodeModel.GetNodeByNodeId(nodeId)
	if err != nil {
		this.ErrorLog("查找节点 "+nodeId+" 失败: "+err.Error())
		this.viewError("节点不存在！", "default")
	}
	if len(node) == 0 {
		this.viewError("节点不存在！", "default")
	}

	messages, err := remotes.NewMessageByNode(node).GetMessages()
	if err != nil {
		this.viewError(err.Error(), "default")
	}

	this.Data["messages"] = messages
	this.viewLayout("node/message", "default")
}

func (this *NodeController) Reload() {

	nodeId := this.GetString("node_id", "")
	if nodeId == "" {
		this.jsonError("节点不存在！")
	}

	node, err := models.NodeModel.GetNodeByNodeId(nodeId)
	if err != nil{
		this.ErrorLog("查找节点 "+nodeId+" 失败: "+err.Error())
		this.jsonError("节点不存在！")
	}
	if len(node) == 0 {
		this.jsonError("节点不存在！")
	}

	err = remotes.NewSystemByNode(node).ReloadSystem()
	if err != nil {
		this.ErrorLog("重载节点 "+nodeId+" 失败")
		this.jsonError("重载失败！")
	}
	this.InfoLog("重载节点 "+nodeId+" 成功")
	this.jsonSuccess("重载节点成功", nil, "message/list?node_id="+nodeId)
}
