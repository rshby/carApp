package service

import (
	"carApp/graph/model"
	"context"
)

type ICarService interface {
	Insert(ctx context.Context, request *model.CreateCarRequest) (*model.CarDetailResponse, error)
	GetById(ctx context.Context, id string) (*model.CarDetailResponse, error)
}
