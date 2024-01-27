package service

import (
	"carApp/app/logging"
	"carApp/app/model/entity"
	repository "carApp/app/repository/interface"
	service "carApp/app/service/interface"
	"carApp/graph/model"
	"context"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/opentracing/opentracing-go"
	"strconv"
)

type CarService struct {
	Log           logging.ILogger
	Validate      *validator.Validate
	CarRepository repository.ICarRepository
}

func NewCarProvider(log logging.ILogger, validate *validator.Validate,
	carRepo repository.ICarRepository) service.ICarService {
	return &CarService{log, validate, carRepo}
}

// method insert
func (c *CarService) Insert(ctx context.Context, request *model.CreateCarRequest) (*model.CarDetailResponse, error) {
	span, ctxTracing := opentracing.StartSpanFromContext(ctx, "CarService Insert")
	defer span.Finish()

	// create entity
	input := entity.Car{
		Id:    request.ID,
		Name:  request.Name,
		Brand: request.Brand,
		Year:  request.Year,
		Price: fmt.Sprintf("%v", request.Price),
	}

	// call procedur to insert
	result, err := c.CarRepository.Insert(ctxTracing, &input)
	if err != nil {
		return nil, err
	}

	// mapping to response
	response := model.CarDetailResponse{
		ID:    result.Id,
		Name:  result.Name,
		Brand: result.Brand,
		Year:  result.Brand,
		Price: ToFloat(result.Price),
	}

	return &response, nil
}

func (c *CarService) GetById(ctx context.Context, id string) (*model.CarDetailResponse, error) {
	return nil, errors.New("feature belum jadi")
}

func ToFloat(price string) float64 {
	num, _ := strconv.ParseFloat(price, 64)
	return num
}
