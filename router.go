package main

import (
	"github.com/astaxie/beego"
	"wmqx-ui/app/controllers"
	"wmqx-ui/app/utils"
	"net/http"
)

func init()  {
	initRouter()
}

func initRouter() {
	// router
	beego.BConfig.WebConfig.AutoRender = false
	beego.BConfig.RouterCaseSensitive = false

	beego.Router("/", &controllers.MainController{}, "*:Index")
	beego.Router("/author", &controllers.AuthorController{}, "*:Index")
	beego.AutoRouter(&controllers.AuthorController{})
	beego.AutoRouter(&controllers.MainController{})
	beego.AutoRouter(&controllers.UserController{})
	beego.AutoRouter(&controllers.NodeController{})
	beego.AutoRouter(&controllers.MessageController{})
	beego.AutoRouter(&controllers.ConsumerController{})
	beego.AutoRouter(&controllers.LogController{})
	beego.AutoRouter(&controllers.ProfileController{})
	beego.ErrorHandler("404", http_404)
	beego.ErrorHandler("500", http_500)

	// add template func
	beego.AddFuncMap("dateFormat", utils.NewDate().Format)
}

func http_404(rs http.ResponseWriter, req *http.Request) {
	rs.Write([]byte("404 not found!"))
}

func http_500(rs http.ResponseWriter, req *http.Request) {
	rs.Write([]byte("500 server error!"))
}
