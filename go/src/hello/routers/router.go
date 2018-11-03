package routers

import (
	"hello/controllers"
	"github.com/astaxie/beego"
	"github.com/satori/go.uuid"
	"github.com/astaxie/beego/context"
)

// import (
// 	"github.com/astaxie/beego/context"
// 	"github.com/satori/go.uuid"
// )

func init() {
	beego.InsertFilter("*", beego.BeforeExec, func(ctx *context.Context){
		ctx.Input.SetData("requestid", uuid.Must(uuid.NewV4()))
	})
  beego.Router("/", &controllers.MainController{})
	beego.Router("/fire", &controllers.FireController{})
}
