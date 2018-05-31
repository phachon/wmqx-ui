package controllers

import (
	"strings"
	"wmqx-ui/app/models"
)

type ProfileController struct {
	BaseController
}

func (this *ProfileController) Index() {
	this.viewTemplate("profile/index")
}

func (this *ProfileController) Save() {

	email := strings.Trim(this.GetString("email", ""), "")
	mobile := strings.Trim(this.GetString("mobile", ""), "")

	if email == "" {
		this.jsonError("邮箱不能为空！")
	}
	if mobile == "" {
		this.jsonError("手机号不能为空！")
	}

	_, err := models.UserModel.Update(this.Data["user"].(map[string]string)["user_id"], map[string]interface{}{
		"email":      email,
		"mobile":     mobile,
	})

	if err != nil {
		this.ErrorLog("更新信息失败: "+err.Error())
		this.jsonError("更新信息失败")
	}

	this.User["email"] = email
	this.User["mobile"] = mobile
	this.InfoLog("更新信息成功")
	this.jsonSuccess("更新信息成功", nil, "/profile/index")
}

func (this *ProfileController) Repass() {
	this.viewTemplate("profile/repass")
}

func (this *ProfileController) Savepass() {

	pwd := strings.Trim(this.GetString("old_password", ""), "")
	pwdNew := strings.Trim(this.GetString("new_password", ""), "")
	pwdConfirm := strings.Trim(this.GetString("re_password", ""), "")

	if (pwd == "") || (pwdNew == "") || (pwdConfirm == "") {
		this.jsonError("密码不能为空！")
	}

	p := models.UserModel.EncodePassword(pwd)
	if p != this.User["password"] {
		this.jsonError("旧密码错误")
	}
	if pwdConfirm != pwdNew {
		this.jsonError("新密码和确认密码不一致")
	}

	_, err := models.UserModel.Update(this.Data["user"].(map[string]string)["user_id"], map[string]interface{}{
		"password": models.UserModel.EncodePassword(pwdNew),
	})

	// 阻止日志记录 password
	this.Ctx.Request.PostForm.Del("pwd")
	this.Ctx.Request.PostForm.Del("pwd_new")
	this.Ctx.Request.PostForm.Del("pwd_confirm")

	if err != nil {
		this.ErrorLog("修改密码失败: "+err.Error())
		this.jsonError("修改密码失败")
	}
	this.InfoLog("修改密码成功")
	this.jsonSuccess("修改密码成功", nil, "/profile/repass")
}