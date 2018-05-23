package controllers

import (
	"wmqx-ui/app/models"
	"wmqx-ui/app/remotes"
)

type MessageController struct {
	BaseController
}

func (this *MessageController) Add() {

	nodeId := this.GetString("node_id")

	node, err := models.NodeModel.GetNodeByNodeId(nodeId)
	if err != nil {
		this.viewError("节点不存在", "default")
	}

	this.Data["node"] = node
	this.viewLayout("message/form", "default")
}

func (this *MessageController) List() {

	nodeId := this.GetString("node_id")

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

	messages, err := remotes.NewMessageByNode(defaultNode).GetMessages()
	if err != nil {
		this.viewError(err.Error(), "template")
	}
	this.Data["nodes"] = nodes
	this.Data["node_id"] = nodeId
	this.Data["messages"] = messages
	this.viewTemplate("message/list")
}

func (this *MessageController) Save() {

	nodeId := this.GetString("node_id", "")
	name := this.GetString("name")
	mode := this.GetString("mode", "topic")
	durable := this.GetString("durable", "0")
	isNeedToken := this.GetString("is_need_token", "1")
	token := this.GetString("token")
	comment := this.GetString("comment", "")

	if nodeId == "" {
		this.jsonError("没有选择节点")
	}
	if name == "" {
		this.jsonError("消息名称不能为空")
	}
	if mode == "" || (mode != "topic" && mode != "fanout" && mode != "direct"){
		this.jsonError("消息模式错误")
	}
	if name == "" {
		this.jsonError("消息名称不能为空")
	}
	if isNeedToken == "1" && token == "" {
		this.jsonError("Token不能为空")
	}
	if comment == "" {
		this.jsonError("消息备注不能为空")
	}

	node, err := models.NodeModel.GetNodeByNodeId(nodeId)
	if err != nil {
		this.jsonError("节点错误")
	}
	if len(node) == 0 {
		this.jsonError("节点不存在")
	}

	err = remotes.NewMessageByNode(node).AddMessage(map[string]string{
		"name": name,
		"comment": comment,
		"durable": durable,
		"is_need_token": isNeedToken,
		"mode": mode,
		"token": token,
	})
	if err != nil {
		this.jsonError(err.Error())
	}

	this.jsonSuccess("添加消息成功", nil, "/message/list?node_id="+nodeId)
}

func (this *MessageController) Edit() {

	nodeId := this.GetString("node_id", "")
	name := this.GetString("message_name")

	if nodeId == "" {
		this.viewError("没有选择节点", "default")
	}
	if name == "" {
		this.viewError("没有选择消息", "default")
	}
	node, err := models.NodeModel.GetNodeByNodeId(nodeId)
	if err != nil {
		this.viewError("节点错误", "default")
	}
	if len(node) == 0 {
		this.viewError("节点不存在", "default")
	}

	message, err := remotes.NewMessageByNode(node).GetMessageByName(name)
	if err != nil {
		this.viewError(err.Error(), "default")
	}

	this.Data["message"] = message
	this.Data["nodeId"] = nodeId
	this.viewLayout("message/edit", "default")
}

func (this *MessageController) Modify() {

	nodeId := this.GetString("node_id", "")
	name := this.GetString("name")
	mode := this.GetString("mode", "topic")
	durable := this.GetString("durable", "0")
	isNeedToken := this.GetString("is_need_token", "1")
	token := this.GetString("token")
	comment := this.GetString("comment", "")

	if nodeId == "" {
		this.jsonError("没有选择节点")
	}
	if name == "" {
		this.jsonError("消息名称不能为空")
	}
	if mode == "" || (mode != "topic" && mode != "fanout" && mode != "direct"){
		this.jsonError("消息模式错误")
	}
	if name == "" {
		this.jsonError("消息名称不能为空")
	}
	if isNeedToken == "1" && token == "" {
		this.jsonError("Token不能为空")
	}
	if comment == "" {
		this.jsonError("消息备注不能为空")
	}

	node, err := models.NodeModel.GetNodeByNodeId(nodeId)
	if err != nil {
		this.jsonError("节点错误")
	}
	if len(node) == 0 {
		this.jsonError("节点不存在")
	}

	err = remotes.NewMessageByNode(node).UpdateMessage(map[string]string{
		"name": name,
		"comment": comment,
		"durable": durable,
		"is_need_token": isNeedToken,
		"mode": mode,
		"token": token,
	})
	if err != nil {
		this.jsonError(err.Error())
	}

	this.jsonSuccess("修改消息成功", nil, "/message/list?node_id="+nodeId)
}

func (this *MessageController) Delete() {

	nodeId := this.GetString("node_id", "")
	name := this.GetString("message_name")

	if nodeId == "" {
		this.jsonError("没有选择节点")
	}
	if name == "" {
		this.jsonError("消息名称不能为空")
	}

	node, err := models.NodeModel.GetNodeByNodeId(nodeId)
	if err != nil {
		this.jsonError("节点错误")
	}
	if len(node) == 0 {
		this.jsonError("节点不存在")
	}

	err = remotes.NewMessageByNode(node).DeleteMessage(name)
	if err != nil {
		this.jsonError(err.Error())
	}

	this.jsonSuccess("删除消息成功", nil, "/message/list?node_id="+nodeId)
}