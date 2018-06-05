package controllers

import (
	"strings"
	"wmqx-ui/app/models"
	"time"
	"wmqx-ui/app/remotes"
	"strconv"
	"github.com/astaxie/beego/httplib"
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

	var count int64
	var nodes []map[string]string
	number := 20
	limit := (page - 1) * number
	var err error
	if this.roleIsUser() {
		nodeIds := []string{}
		userNodes, err := models.UserNodeModel.GetUserNodeByUserId(this.UserID)
		if err != nil {
			this.ErrorLog("获取节点失败："+err.Error())
			this.viewError("获取节点失败", "template")
			return
		}
		for _, userNode := range userNodes {
			nodeIds = append(nodeIds, userNode["node_id"])
		}
		if len(nodeIds) == 0 {
			count = 0
		}else {
			if keywords["comment"] != "" {
				count, err = models.NodeModel.CountNodesByKeywordsInNodeIds(keywords, nodeIds)
				nodes, err = models.NodeModel.GetNodesByKeywordsAndLimitInNodeIds(keywords, nodeIds, limit, number)
			} else {
				count = int64(len(nodeIds))
				nodes, err = models.NodeModel.GetNodesByLimitInNodeIds(limit, nodeIds, number)
			}
		}
	}else {
		if keywords["comment"] != "" {
			count, err = models.NodeModel.CountNodesByKeywords(keywords)
			nodes, err = models.NodeModel.GetNodesByKeywordsAndLimit(keywords, limit, number)
		} else {
			count, err = models.NodeModel.CountNodes()
			nodes, err = models.NodeModel.GetNodesByLimit(limit, number)
		}
	}
	if err != nil {
		this.ErrorLog("获取节点失败："+err.Error())
		this.viewError(err.Error(), "template")
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
	nodes, _ := models.NodeModel.GetNodeByManagerUri(managerURI)
	if len(nodes) > 0 {
		this.jsonError("管理URI已经存在")
	}
	nodes, _ = models.NodeModel.GetNodeByPublishUri(publishURI)
	if len(nodes) > 0 {
		this.jsonError("发布URI已经存在")
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
	nodeId, err := models.NodeModel.Insert(nodeValue)
	if err != nil {
		this.ErrorLog("添加节点失败: "+err.Error())
		this.jsonError("添加节点失败！")
	}

	// add user_node
	userNode := map[string]interface{}{
		"user_id": this.UserID,
		"node_id": nodeId,
		"create_time": time.Now().Unix(),
	}
	_, err = models.UserNodeModel.Insert(userNode)
	if err != nil {
		this.ErrorLog("插入用户 "+this.UserID+" 节点 "+strconv.FormatInt(nodeId, 10)+" 关系失败: "+err.Error())
		this.jsonError("添加节点失败")
	}
	this.InfoLog("添加节点成功")
	this.jsonSuccess("添加节点成功", nil, "/node/list")
}

func (this *NodeController) Edit() {

	nodeId := this.GetString("node_id", "")
	if nodeId == "" {
		this.viewError("节点不存在", "template")
	}

	node, err := models.NodeModel.GetNodeByNodeId(nodeId)
	if err != nil {
		this.ErrorLog("查找节点 "+nodeId+" 失败: "+err.Error())
		this.viewError("节点不存在", "template")
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
		this.viewError("节点不存在", "template")
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
	isExists, _ := models.NodeModel.HaveNodeByManagerUriAndNodeId(managerURI, nodeId)
	if isExists {
		this.jsonError("管理URI已经存在")
	}
	isExists, _ = models.NodeModel.HaveNodeByPublishUriAndNodeId(publishURI, nodeId)
	if isExists {
		this.jsonError("发布URI已经存在")
	}

	nodeValue := map[string]interface{}{
		"manager_uri":      strings.TrimRight(managerURI, "/"),
		"publish_uri":      strings.TrimRight(publishURI, "/"),
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
		this.viewError("获取节点消息列表失败", "default")
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

func (this *NodeController) Status() {

	nodeId := this.GetString("node_id", "")
	if nodeId == "" {
		this.jsonError("节点不存在！")
	}
	node, err := models.NodeModel.GetNodeByNodeId(nodeId)
	if err != nil{
		this.jsonError("节点不存在！")
	}
	if len(node) == 0 {
		this.jsonError("节点不存在！")
	}
	v := map[string]interface{}{}
	err = httplib.Get(node["manager_uri"]).ToJSON(&v)
	if err != nil {
		this.jsonError("server error")
	}
	if v["code"].(float64) == 0 {
		this.jsonError(v["message"].(string))
	}

	data := v["data"].(map[string]interface{})
	this.jsonSuccess("success", data)
}
