package controllers

import (
	"strings"
	"time"
	"wmqx-ui/app/models"
	"fmt"
)

type UserController struct {
	BaseController
}

func (this *UserController) List() {

	page, _ := this.GetInt("page", 1)
	keyword := strings.TrimSpace(this.GetString("keyword", ""))

	number := 20
	limit := (page - 1) * number
	var err error
	var count int64
	var users []map[string]string
	if keyword != "" {
		count, err = models.UserModel.CountUsersByKeyword(keyword)
		users, err = models.UserModel.GetUsersByKeywordAndLimit(keyword, limit, number)
	} else {
		count, err = models.UserModel.CountUsers()
		users, err = models.UserModel.GetUsersByLimit(limit, number)
	}
	if err != nil {
		this.ErrorLog("查找用户失败: "+err.Error())
		this.viewError(err.Error(), "/user/list")
	}

	this.Data["users"] = users
	this.Data["keyword"] = keyword
	this.SetPaginator(number, count)

	this.viewTemplate("user/list")
}

func (this *UserController) Add() {
	this.viewTemplate("user/form")
}

func (this *UserController) Save() {

	username := strings.TrimSpace(this.GetString("username"))
	password := strings.TrimSpace(this.GetString("password"))
	email := strings.TrimSpace(this.GetString("email"))
	mobile := strings.TrimSpace(this.GetString("mobile"))

	if username == "" {
		this.jsonError("用户名不能为空")
	}
	if password == "" {
		this.jsonError("密码不能为空")
	}
	if email == "" {
		this.jsonError("邮箱不能为空")
	}
	if mobile == "" {
		this.jsonError("手机号不能为空")
	}

	isExists, err := models.UserModel.HasUsername(username)
	if err != nil {
		this.ErrorLog("查找用户 "+username+" 失败: "+err.Error())
		this.jsonError("添加用户失败！")
	}
	if isExists {
		this.jsonError("该用户名已存在！")
	}

	userValue := map[string]interface{}{
		"username":    username,
		"password":    models.UserModel.EncodePassword(password),
		"email":       email,
		"mobile":      mobile,
		"role":        models.USER_ROLE_USER,
		"create_time": time.Now().Unix(),
		"update_time": time.Now().Unix(),
	}

	_, err = models.UserModel.Insert(userValue)
	if err != nil {
		this.ErrorLog("添加用户失败: "+err.Error())
		this.jsonError("添加用户失败！")
	}

	this.InfoLog("添加用户成功")
	this.jsonSuccess("添加用户成功", nil, "/user/list")
}

func (this *UserController) Edit() {

	userId := this.GetString("user_id", "")
	if userId == "" {
		this.viewError("用户不存在", "/user/list")
	}

	user, err := models.UserModel.GetUserByUserId(userId)
	if err != nil {
		this.ErrorLog("查找用户 "+userId+" 失败: "+err.Error())
		this.viewError("用户不存在", "default", "/user/list")
	}
	if user["role"] == fmt.Sprintf("%d", models.USER_ROLE_ROOT) {
		this.viewError("不能修改超级管理员！", "default")
	}

	this.Data["user"] = user
	this.viewLayout("user/edit", "default")
}

func (this *UserController) Modify() {

	userId := strings.TrimSpace(this.GetString("user_id"))
	email := strings.TrimSpace(this.GetString("email"))
	mobile := strings.TrimSpace(this.GetString("mobile"))
	role, _ := this.GetInt("role")

	if email == "" {
		this.jsonError("邮箱不能为空")
	}
	if mobile == "" {
		this.jsonError("手机号不能为空")
	}
	if (role != models.USER_ROLE_ADMIN) && (role != models.USER_ROLE_ROOT) && (role != models.USER_ROLE_USER) {
		this.jsonError("角色选择错误")
	}

	userValue := map[string]interface{}{
		"email":       email,
		"mobile":      mobile,
		"role":        role,
		"create_time": time.Now().Unix(),
		"update_time": time.Now().Unix(),
	}

	_, err := models.UserModel.Update(userId, userValue)
	if err != nil {
		this.ErrorLog("修改用户 "+userId+" 失败: "+err.Error())
		this.jsonError("修改用户失败！")
	}
	this.InfoLog("修改用户 "+userId+" 成功")
	this.jsonSuccess("修改用户成功", nil, "/user/list")
}

func (this *UserController) Node() {

	userId := this.GetString("user_id", "")
	if userId == "" {
		this.viewError("用户不存在", "/user/list")
	}

	user, err := models.UserModel.GetUserByUserId(userId)
	if err != nil {
		this.ErrorLog("查找用户 "+userId+" 失败: "+err.Error())
		this.viewError("用户不存在", "default", "/user/list")
	}
	if user["role"] == fmt.Sprintf("%d", models.USER_ROLE_ROOT) {
		this.viewError("不能修改超级管理员！", "default")
	}

	nodes, err := models.NodeModel.GetNodes()
	if err != nil {
		this.ErrorLog("查找节点失败: "+err.Error())
		this.viewError("查找节点失败", "default")
	}
	userNodes, err := models.UserNodeModel.GetUserNodeByUserId(userId)
	if err != nil {
		this.ErrorLog("查找用户 "+userId+" 节点失败: "+err.Error())
		this.viewError("查找节点失败", "default")
	}
	for _, node := range nodes {
		node["is_default"] = "0"
		for _, userNode := range userNodes {
			if userNode["node_id"] == node["node_id"] {
				node["is_default"] = "1"
				break
			}
		}
	}

	this.Data["user"] = user
	this.Data["nodes"] = nodes
	this.viewLayout("user/node", "default")
}

func (this *UserController) Privilege() {

	nodeIds := this.GetStrings("node_id")
	userId := this.GetString("user_id", "")
	if userId == "" {
		this.jsonError("用户不存在", "/user/list")
	}
	if len(nodeIds) == 0 {
		this.jsonError("没有选择节点")
	}

	_, err := models.UserModel.GetUserByUserId(userId)
	if err != nil {
		this.ErrorLog("查找用户 "+userId+" 失败: "+err.Error())
		this.jsonError("用户不存在", "default", "/user/list")
	}

	err = models.UserNodeModel.DeleteUserNodeByUserId(userId)
	if err != nil {
		this.ErrorLog("删除用户 "+userId+" 节点失败: "+err.Error())
		this.jsonError("节点授权失败")
	}

	userNodes := make([]map[string]interface{}, 0)
	for _, nodeId := range nodeIds {
		userNode := map[string]interface{}{
			"user_id": userId,
			"node_id": nodeId,
			"create_time": time.Now().Unix(),
		}
		userNodes = append(userNodes, userNode)
	}
	_, err = models.UserNodeModel.InsertBatch(userNodes)
	if err != nil {
		this.ErrorLog("插入用户 "+userId+" 节点失败: "+err.Error())
		this.jsonError("节点授权失败")
	}

	this.InfoLog("用户 "+userId+"节点授权成功")
	this.jsonSuccess("节点授权成功", nil, "/user/list")
}

func (this *UserController) Remove() {

	userId := this.GetString("user_id", "")

	if userId == "" {
		this.jsonError("没有选择用户！")
	}

	user, err := models.UserModel.GetUserByUserId(userId)
	if err != nil {
		this.ErrorLog("查找用户 "+userId+" 失败: "+err.Error())
		this.jsonError("用户不存在！")
	}
	if len(user) == 0 {
		this.jsonError("用户不存在！")
	}

	userValue := map[string]interface{}{
		"is_delete":   models.USER_DELETE,
		"update_time": time.Now().Unix(),
	}

	_, err = models.UserModel.Update(userId, userValue)
	if err != nil {
		this.ErrorLog("删除用户 "+userId+" 失败: "+err.Error())
		this.jsonError("删除用户失败！")
	}
	this.InfoLog("删除用户 "+userId+" 成功")
	this.jsonSuccess("删除用户成功", nil, "/user/list")
}