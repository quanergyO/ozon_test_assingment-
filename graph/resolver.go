package graph

import "github.com/quanergyo/ozon-test-assingment/repository"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	repo *repository.Repository
}

func NewResolver(repo *repository.Repository) *Resolver {
	return &Resolver{
		repo: repo,
	}
}
