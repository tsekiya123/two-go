package controllers

import (
	"github.com/astaxie/beego"
	"hello/service"
)

type FireController struct {
	BaseController
}

func (c *FireController) Get() {
	beego.Info("requestid in controller: ", c.Ctx.Input.GetData("requestid"))

	//service.CryAll(c.Ctx)
	service.GetFire(c.Ctx)

	c.Data["json"] = map[string]interface{}{"name": "fire🔥"}
	c.ServeJSON()
}
