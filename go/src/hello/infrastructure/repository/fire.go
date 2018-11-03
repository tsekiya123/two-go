package repository

import (
	"github.com/astaxie/beego/context"
	"hello/infrastructure/repositoryImpl"
	"hello/models"
)

type FireRepository interface {
	GetFire(ctx *context.Context) models.Fire
}

func GetFireForRepository(ctx *context.Context) models.Fire {
	impl := repositoryImpl.NewFireRepositoryImpl(ctx)
	return impl.GetFire()
}
