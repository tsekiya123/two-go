package service

import (
	"github.com/astaxie/beego"
)
import (
	"github.com/astaxie/beego/context"
	// "github.com/astaxie/beego"
)
type Dog struct {
	Ctx *context.Context
}
func (d *Dog) Cry() {
	beego.Info("わんわん", d.Ctx.Input.GetData("requestid"))
}
