package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/satori/go.uuid"
	"hello/controllers"
)

func init() {
	beego.InsertFilter("*", beego.BeforeExec, func(ctx *context.Context){
		ctx.Input.SetData("requestid", uuid.Must(uuid.NewV4()))
	})
  beego.Router("/", &controllers.MainController{})
	beego.Router("/fire", &controllers.FireController{})
}
