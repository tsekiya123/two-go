package repositoryImpl

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"hello/models"
)

type FireRepositoryImpl struct {
	ctx *context.Context
}

func NewFireRepositoryImpl(ctx *context.Context) *FireRepositoryImpl{
	return &FireRepositoryImpl{
		ctx: ctx,
	}
}

func (c *FireRepositoryImpl) GetFire() models.Fire {
	//time.Sleep(10 * time.Second)
	beego.Info("requestid in impl: ", c.ctx.Input.GetData("requestid"))
	return models.Fire{
		Size: "サイズ",
		Temp: "温度",
	}
}