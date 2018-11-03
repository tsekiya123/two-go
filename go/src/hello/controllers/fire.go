package controllers

import (
	"github.com/astaxie/beego"
	"hello/service"
)
// import (
// 	"github.com/astaxie/beego/context"
// 	"github.com/satori/go.uuid"
// )

type FireController struct {
	beego.Controller
}

func (c *FireController) Get() {

	beego.Info("FireController: hogehoge")
	beego.Info("Get: ", c.Ctx.Input.GetData("requestid"))

	dog := service.Dog{}
	dog.Ctx = c.Ctx
	dog.Cry()

	c.Data["json"] = map[string]interface{}{"name": "fire🔥"}
	c.ServeJSON()
}
