package controllers

import (
	"github.com/astaxie/beego"
	"hello/service"
)

type FireController struct {
	beego.Controller
}

func (c *FireController) Get() {

	beego.Info("FireController: hogehoge")
	beego.Info("Get: ", c.Ctx.Input.GetData("requestid"))

	//dog := service.Dog{}
	//dog.Ctx = c.Ctx
	service.CryAll(c.Ctx)

	c.Data["json"] = map[string]interface{}{"name": "fire🔥"}
	c.ServeJSON()
}
