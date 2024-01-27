package graph

import service "carApp/app/service/interface"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CarService service.ICarService
}

func NewResolver(carService service.ICarService) ResolverRoot {
	return &Resolver{
		CarService: carService,
	}
}
