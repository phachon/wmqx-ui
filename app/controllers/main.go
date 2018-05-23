package controllers

type MainController struct {
	BaseController
}

func (this *MainController) Index() {
	this.viewTemplate("main/index")
}
