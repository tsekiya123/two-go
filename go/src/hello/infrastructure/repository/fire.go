package repository

import "hello/models"

type FireRepository interface {
	GetFire() models.Fire
}
