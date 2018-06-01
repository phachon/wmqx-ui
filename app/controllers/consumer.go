package controllers

import (
	"wmqx-ui/app/remotes"
	"wmqx-ui/app/models"
	"fmt"
)

type ConsumerController struct {
	BaseController
}

func (this *ConsumerController) Add() {
	nodeId := this.GetString("node_id", "")
	messageName := this.GetString("message_name", "")

	this.Data["node_id"] = nodeId
	this.Data["message_name"] = messageName
	this.viewLayout("consumer/form", "default")
}

func (this *ConsumerController) Save() {

	nodeId := this.GetString("node_id", "")
	messageName := this.GetString("message_name")
	url := this.GetString("url", "")
	routeKey := this.GetString("route_key", "")
	timeout := this.GetString("timeout", "2000")
	checkCode := this.GetString("check_code", "0")
	code := this.GetString("code", "200")
	comment := this.GetString("comment", "")

	if nodeId == "" {
		this.jsonError("没有选择节点")
	}
	if messageName == "" {
		this.jsonError("消息名称不能为空")
	}
	if url == ""{
		this.jsonError(" 消费URL不能为空")
	}
	if timeout == "" {
		this.jsonError("超时时间不能为0")
	}
	if checkCode == "1" && code == "" {
		this.jsonError("httpCode不能为0")
	}

	node, err := models.NodeModel.GetNodeByNodeId(nodeId)
	if err != nil {
		this.ErrorLog("获取节点 "+nodeId+" 失败: "+err.Error())
		this.jsonError("节点错误")
	}
	if len(node) == 0 {
		this.jsonError("节点不存在")
	}

	err = remotes.NewConsumerByNode(node, messageName).AddConsumer(map[string]string{
		"url": url,
		"route_key": routeKey,
		"timeout": timeout,
		"code": code,
		"check_code": checkCode,
		"comment": comment,
	})
	if err != nil {
		this.ErrorLog("添加消费者失败: "+err.Error())
		this.jsonError("添加消息失败")
	}
	this.InfoLog("添加消费者成功")
	redirect := fmt.Sprintf("/message/consumer?node_id=%s&message_name=%s", nodeId, messageName)
	this.jsonSuccess("添加消息成功", nil, redirect)
}

func (this *ConsumerController) Edit() {

	nodeId := this.GetString("node_id", "")
	messageName := this.GetString("message_name", "")
	consumerId := this.GetString("consumer_id", "")

	if consumerId == "" {
		this.viewError("消费者错误", "default")
	}
	node, err := models.NodeModel.GetNodeByNodeId(nodeId)
	if err != nil {
		this.ErrorLog("获取节点失败: "+err.Error())
		this.viewError("节点错误", "default")
	}
	if len(node) == 0 {
		this.viewError("节点不存在", "default")
	}

	consumer, err := remotes.NewConsumerByNode(node, messageName).GetConsumerByConsumerId(consumerId)
	if err != nil {
		this.ErrorLog("获取消费者失败: "+err.Error())
		this.viewError("获取消费者错误", "default")
	}
	this.Data["node_id"] = nodeId
	this.Data["message_name"] = messageName
	this.Data["consumer"] = consumer
	this.viewLayout("consumer/edit", "default")
}

func (this *ConsumerController) Modify() {

	nodeId := this.GetString("node_id", "")
	consumerId := this.GetString("consumer_id", "")
	messageName := this.GetString("message_name")
	url := this.GetString("url", "")
	routeKey := this.GetString("route_key", "")
	timeout := this.GetString("timeout", "2000")
	checkCode := this.GetString("check_code", "0")
	code := this.GetString("code", "200")
	comment := this.GetString("comment", "")

	if nodeId == "" {
		this.jsonError("没有选择节点")
	}
	if consumerId == "" {
		this.jsonError("没有选择消费者")
	}
	if messageName == "" {
		this.jsonError("消息名称不能为空")
	}
	if url == ""{
		this.jsonError(" 消费URL不能为空")
	}
	if timeout == "" {
		this.jsonError("超时时间不能为0")
	}
	if checkCode == "1" && code == "" {
		this.jsonError("httpCode不能为0")
	}

	node, err := models.NodeModel.GetNodeByNodeId(nodeId)
	if err != nil {
		this.ErrorLog("获取节点失败: "+err.Error())
		this.jsonError("节点错误")
	}
	if len(node) == 0 {
		this.jsonError("节点不存在")
	}

	err = remotes.NewConsumerByNode(node, messageName).UpdateConsumerByConsumerId(consumerId, map[string]string{
		"url": url,
		"route_key": routeKey,
		"timeout": timeout,
		"code": code,
		"check_code": checkCode,
		"comment": comment,
	})
	if err != nil {
		this.ErrorLog("修改消费者失败: "+err.Error())
		this.jsonError("修改消费者失败")
	}
	this.InfoLog("修改消费者 "+consumerId+" 成功")
	redirect := fmt.Sprintf("/message/consumer?node_id=%s&message_name=%s", nodeId, messageName)
	this.jsonSuccess("修改消费者成功", nil, redirect)
}

func (this *ConsumerController) Delete() {

	nodeId := this.GetString("node_id", "")
	consumerId := this.GetString("consumer_id", "")
	messageName := this.GetString("message_name")

	if nodeId == "" {
		this.jsonError("没有选择节点")
	}
	if consumerId == "" {
		this.jsonError("没有选择消费者")
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

	err = remotes.NewConsumerByNode(node, messageName).DeleteConsumerByConsumerId(consumerId)
	if err != nil {
		this.ErrorLog("删除消费者失败: "+err.Error())
		this.jsonError("删除消费者失败")
	}
	this.InfoLog("删除消费者 "+consumerId+" 成功")
	redirect := fmt.Sprintf("/message/consumer?node_id=%s&message_name=%s", nodeId, messageName)
	this.jsonSuccess("删除消费者成功", nil, redirect)
}