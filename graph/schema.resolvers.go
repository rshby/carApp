package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

import (
	"carApp/graph/model"
	"context"
	opentracing "github.com/opentracing/opentracing-go"
)

// CreateCar is the resolver for the CreateCar field.
func (r *mutationResolver) CreateCar(ctx context.Context, input model.CreateCarRequest) (*model.CarDetailResponse, error) {
	span, ctxTracing := opentracing.StartSpanFromContext(ctx, "CreateCar")
	defer span.Finish()

	return r.CarService.Insert(ctxTracing, &input)
}

// GetCar is the resolver for the GetCar field.
func (r *queryResolver) GetCar(ctx context.Context, input model.GetCar) (*model.CarDetailResponse, error) {
	span, ctxTracing := opentracing.StartSpanFromContext(ctx, "GetCar")
	defer span.Finish()

	return r.CarService.GetById(ctxTracing, input.ID)
}

// GetAll is the resolver for the GetAll field.
func (r *queryResolver) GetAll(ctx context.Context) ([]*model.CarDetailResponse, error) {
	span, ctxTracing := opentracing.StartSpanFromContext(ctx, "GetAll")
	defer span.Finish()

	return r.CarService.GetAll(ctxTracing)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
