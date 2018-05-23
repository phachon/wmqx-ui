package controllers

import (
	"encoding/json"
	"strings"
	"github.com/astaxie/beego"
	"wmqx-ui/app/utils"
	"fmt"
	"wmqx-ui/app/models"
)

type BaseController struct {
	beego.Controller
	UserID string
	User   map[string]string
	controllerName string
}

type JsonResponse struct {
	Code     int                    `json:"code"`
	Message  interface{}            `json:"message"`
	Data     interface{}            `json:"data"`
	Redirect map[string]interface{} `json:"redirect"`
}

// prepare
func (this *BaseController) Prepare() {
	controllerName, _ := this.GetControllerAndAction()
	this.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])

	if this.controllerName == "author" {
		return
	}

	if !this.isLogin() {
		this.Redirect("/author/index", 302)
		this.StopRun()
	}

	this.User = this.GetSession("author").(map[string]string)
	this.UserID = this.User["user_id"]
	this.Data["user"] = this.User
	this.Layout = "layout/default.html"
}

// check is login
func (this *BaseController) isLogin() bool {
	passport := beego.AppConfig.String("author.passport")
	cookie := this.Ctx.GetCookie(passport)
	// cookie is empty
	if cookie == "" {
		return false
	}
	user := this.GetSession("author")
	// session is empty
	if user == nil {
		return false
	}
	cookieValue, _ := utils.Encrypt.Base64Decode(cookie)
	identifyList := strings.Split(cookieValue, "@")
	if cookieValue == "" || len(identifyList) != 2 {
		fmt.Println(identifyList)
		return false
	}
	username := identifyList[0]
	identify := identifyList[1]
	userValue := user.(map[string]string)

	// cookie  session name
	if username != userValue["username"] {
		return false
	}
	// UAG and IP
	if identify != utils.Encrypt.Md5Encode(this.Ctx.Request.UserAgent()+this.getClientIp()+userValue["password"]) {
		return false
	}
	// success
	return true
}

// view layout title
func (this *BaseController) viewLayoutTitle(title, viewName, layout string) {
	this.Layout = "layout/" + layout + ".html"
	this.TplName = viewName + ".html"
	this.Data["navName"] = this.controllerName
	this.Data["title"] = title
	this.Render()
}

// view layout
func (this *BaseController) viewLayout(viewName, layout string) {
	this.Layout = "layout/" + layout + ".html"
	this.TplName = viewName + ".html"
	this.Data["title"] = ""
	this.Data["navName"] = this.controllerName
	this.Render()
}

// view
func (this *BaseController) viewTemplate(viewName string) {
	this.Layout = "layout/template.html"
	this.TplName = viewName + ".html"
	this.Data["title"] = ""
	this.Data["navName"] = this.controllerName
	this.Render()
}

// error view
func (this *BaseController) viewError(errorMessage string, layout string, data ...interface{}) {
	this.Layout = "layout/" + layout + ".html"
	redirect := "/"
	sleep := 2000
	if len(data) > 0 {
		redirect = data[0].(string)
	}
	if len(data) > 1 {
		sleep = data[1].(int)
	}
	this.Data["navName"] = this.controllerName
	this.TplName = "error/error.html"
	this.Data["title"] = "error"
	this.Data["message"] = errorMessage
	this.Data["redirect"] = redirect
	this.Data["sleep"] = sleep
	this.Render()
}

// view title
func (this *BaseController) viewTitle(title, layout string, viewName string) {
	this.Layout = "layout/" + layout + ".html"
	this.TplName = viewName + ".html"
	this.Data["title"] = title
	this.Render()
}

// return json success
func (this *BaseController) jsonSuccess(message interface{}, data ...interface{}) {
	url := ""
	sleep := 2000
	var _data interface{}
	if len(data) > 0 {
		_data = data[0]
	}
	if len(data) > 1 {
		url = data[1].(string)
	}
	if len(data) > 2 {
		sleep = data[2].(int)
	}
	this.Data["json"] = JsonResponse{
		Code:    1,
		Message: message,
		Data:    _data,
		Redirect: map[string]interface{}{
			"url":   url,
			"sleep": sleep,
		},
	}

	j, err := json.MarshalIndent(this.Data["json"], "", "\t")
	if err != nil {
		this.Abort(err.Error())
	} else {
		this.Abort(string(j))
	}
}

// return json error
func (this *BaseController) jsonError(message interface{}, data ...interface{}) {
	url := ""
	sleep := 2000
	var _data interface{}
	if len(data) > 0 {
		_data = data[0]
	}
	if len(data) > 1 {
		url = data[1].(string)
	}
	if len(data) > 2 {
		sleep = data[2].(int)
	}
	this.Data["json"] = JsonResponse{
		Code:    0,
		Message: message,
		Data:    _data,
		Redirect: map[string]interface{}{
			"url":   url,
			"sleep": sleep,
		},
	}
	j, err := json.MarshalIndent(this.Data["json"], "", " \t")
	if err != nil {
		this.Abort(err.Error())
	} else {
		this.Abort(string(j))
	}
}

// get client ip
func (this *BaseController) getClientIp() string {
	s := strings.Split(this.Ctx.Request.RemoteAddr, ":")
	return s[0]
}

// paginator
func (this *BaseController) SetPaginator(per int, nums int64) *utils.Paginator {
	p := utils.NewPaginator(this.Ctx.Request, per, nums)
	this.Data["paginator"] = p
	return p
}

func (this *BaseController) roleIsRoot() bool {
	return this.User["role"] == fmt.Sprintf("%d", models.USER_ROLE_ROOT)
}

func (this *BaseController) roleIsAdmin() bool {
	return this.User["role"] == fmt.Sprintf("%d", models.USER_ROLE_ADMIN)
}

func (this *BaseController) roleIsUser() bool {
	return this.User["role"] == fmt.Sprintf("%d", models.USER_ROLE_USER)
}