package controllers

import (
	"github.com/astaxie/beego"
)

type FourZeroController struct {
	beego.Controller
}

func (this *FourZeroController) Get() {
	this.TplName = "404.html"
}
