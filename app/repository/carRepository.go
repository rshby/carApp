package repository

import (
	"carApp/app/model/entity"
	repository "carApp/app/repository/interface"
	"context"
	"database/sql"
	"errors"
	"github.com/opentracing/opentracing-go"
)

type CarRepository struct {
	DB *sql.DB
}

// function provider
func NewCarRepository(db *sql.DB) repository.ICarRepository {
	return &CarRepository{db}
}

// method insert
func (c *CarRepository) Insert(ctx context.Context, input *entity.Car) (*entity.Car, error) {
	span, ctxTracing := opentracing.StartSpanFromContext(ctx, "CarRepository Insert")
	defer span.Finish()

	// create prepare statement
	statement, err := c.DB.PrepareContext(ctxTracing, "INSERT INTO car (id, name, brand, year, price) VALUES ($1, $2, $3, $4, $5)")
	defer statement.Close()
	if err != nil {
		return nil, err
	}

	// execute
	result, err := statement.ExecContext(ctxTracing, input.Id, input.Name, input.Brand, input.Year, input.Price)
	if err != nil {
		return nil, err
	}

	if row, _ := result.RowsAffected(); row == 0 {
		return nil, errors.New("failed to insert")
	}

	// success insert
	return input, nil
}

// method get by id
func (c *CarRepository) GetById(ctx context.Context, id string) (*entity.Car, error) {
	span, ctxTracing := opentracing.StartSpanFromContext(ctx, "CarRepository GetById")
	defer span.Finish()

	// create prepare statement
	statement, err := c.DB.PrepareContext(ctxTracing, "SELECT id, name, brand, year, price FROM car WHERE id=$1")
	defer statement.Close()
	if err != nil {
		return nil, err
	}

	// query
	row := statement.QueryRowContext(ctxTracing, id)
	if err = row.Err(); err != nil {
		return nil, err
	}

	// success get data
	var car entity.Car
	if err = row.Scan(&car.Id, &car.Name, &car.Brand, &car.Year, &car.Price); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("record with this id not found")
		}
		return nil, err
	}

	// success get data
	return &car, nil
}
