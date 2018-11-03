package service

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"hello/infrastructure/repository"
	"hello/models"
)

type AnimalActionService interface {
	GetFire() models.Fire
	Cry() string
}

func CryAll(ctx *context.Context) {
	var dog, cat AnimalActionService
	dog = &models.Dog{}
	cat = &models.Cat{}
	beego.Info("温度：", dog.GetFire())
	beego.Info("鳴声：", dog.Cry())

	beego.Info("温度：", cat.GetFire())
	beego.Info("鳴声：", cat.Cry())

	beego.Info("requestid in service: ", ctx.Input.GetData("requestid"))
}

func GetFire(ctx *context.Context) models.Fire {
	beego.Info("requestid in service: ", ctx.Input.GetData("requestid"))
	return repository.GetFireForRepository(ctx)
}