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

	if len(defaultNode) == 0 {
		this.viewError("没有选择节点", "template")
	}

	messages, err := remotes.NewMessageByNode(defaultNode).GetMessages()
	if err != nil {
		this.viewError("wmqx 节点 "+defaultNode["node_id"]+" 请求失败, 请检查是否正常工作", "template")
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
		this.ErrorLog("获取节点失败: "+err.Error())
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
		this.ErrorLog("添加消息失败: "+err.Error())
		this.jsonError("添加消息失败")
	}
	this.InfoLog("添加消息成功")
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
		this.ErrorLog("获取节点失败: "+err.Error())
		this.viewError("节点错误", "default")
	}
	if len(node) == 0 {
		this.viewError("节点不存在", "default")
	}

	message, err := remotes.NewMessageByNode(node).GetMessageByName(name)
	if err != nil {
		this.ErrorLog("获取消息失败: "+err.Error())
		this.viewError("获取消息失败", "default")
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
		this.ErrorLog("获取节点失败: "+err.Error())
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
		this.ErrorLog("修改消息 "+name+" 失败: "+err.Error())
		this.jsonError("修改消息失败")
	}

	this.InfoLog("修改消息 "+name+" 成功")
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
		this.ErrorLog("获取节点失败: "+err.Error())
		this.jsonError("节点错误")
	}
	if len(node) == 0 {
		this.jsonError("节点不存在")
	}

	err = remotes.NewMessageByNode(node).DeleteMessage(name)
	if err != nil {
		this.ErrorLog("删除消息 "+name+" 失败: "+err.Error())
		this.jsonError("删除失败")
	}
	this.InfoLog("删除消息 "+name+" 成功: "+err.Error())
	this.jsonSuccess("删除消息成功", nil, "/message/list?node_id="+nodeId)
}

func (this *MessageController) Consumer() {
	nodeId := this.GetString("node_id", "")
	name := this.GetString("message_name")

	if nodeId == "" {
		this.viewError("没有选择节点", "template")
	}
	if name == "" {
		this.viewError("没有选择消息", "template")
	}
	node, err := models.NodeModel.GetNodeByNodeId(nodeId)
	if err != nil {
		this.ErrorLog("获取节点失败: "+err.Error())
		this.viewError("节点错误", "template")
	}
	if len(node) == 0 {
		this.viewError("节点不存在", "template")
	}
	message, err := remotes.NewMessageByNode(node).GetMessageByName(name)
	if err != nil {
		this.ErrorLog("获取消息 "+name+" 失败: "+err.Error())
		this.viewError("获取消息失败", "template")
	}
	consumers, err := remotes.NewMessageByNode(node).GetConsumersByName(name)
	if err != nil {
		this.ErrorLog("获取消息 "+name+" 消费者失败: "+err.Error())
		this.viewError("查找消费者失败", "template")
	}

	this.Data["message"] = message
	this.Data["consumers"] = consumers
	this.Data["node"] = node
	this.viewLayout("message/consumer", "template")
}

func (this *MessageController) Reload() {

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
		this.ErrorLog("获取节点失败: "+err.Error())
		this.jsonError("节点错误")
	}
	if len(node) == 0 {
		this.jsonError("节点不存在")
	}

	err = remotes.NewMessageByNode(node).ReloadMessage(name)
	if err != nil {
		this.ErrorLog("重载消息" +name+ "失败")
		this.jsonError("重载失败")
	}
	this.InfoLog("重载消息" +name+ "成功")
	this.jsonSuccess("重载消息成功", nil, "/message/consumer?node_id="+nodeId+"&message_name="+name)
}

func (this *MessageController) ConsumerStatus() {

	nodeId := this.GetString("node_id", "")
	messageName := this.GetString("message_name")

	if nodeId == "" {
		this.jsonError("没有选择节点")
	}
	if messageName == "" {
		this.jsonError("消息名称不能为空")
	}

	node, err := models.NodeModel.GetNodeByNodeId(nodeId)
	if err != nil {
		this.ErrorLog("获取节点 "+nodeId+" 失败: "+err.Error())
		this.jsonError("节点错误")
	}
	if len(node) == 0 {
		this.jsonError("节点不存在")
	}

	status, err := remotes.NewMessageByNode(node).GetConsumersStatus(messageName)
	if err != nil {
		this.ErrorLog("获取消费者状态失败: "+err.Error())
		this.jsonError("获取消费者状态失败")
	}

	this.jsonSuccess("获取消费者状态成功", status)
}