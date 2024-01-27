package repository

import (
	"carApp/app/model/entity"
	"context"
)

type ICarRepository interface {
	Insert(ctx context.Context, input *entity.Car) (*entity.Car, error)
	GetById(ctx context.Context, id string) (*entity.Car, error)
}
