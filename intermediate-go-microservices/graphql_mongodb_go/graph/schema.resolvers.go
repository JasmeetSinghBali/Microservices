package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"

	"github.com/Jasmeet-1998/Microservices/intermediate-go-microservices/graphql_mongodb_go/graph/model"
	"github.com/Jasmeet-1998/Microservices/intermediate-go-microservices/graphql_mongodb_go/repository"
)

var db = repository.Connect()

// CreateDonutHub is the resolver for the createDonutHub field.
func (r *mutationResolver) CreateDonutHub(ctx context.Context, input model.CreateDonutHubInput) (*model.DonutHub, error) {
	return db.CreateDonutHub(input), nil
}

// UpdateDonutHub is the resolver for the updateDonutHub field.
func (r *mutationResolver) UpdateDonutHub(ctx context.Context, id string, input model.UpdateDonutHubInput) (*model.DonutHub, error) {
	return db.UpdateDonutHub(id, input), nil
}

// DeleteDonutHub is the resolver for the deleteDonutHub field.
func (r *mutationResolver) DeleteDonutHub(ctx context.Context, id string) (*model.DeleteDonutHubResponse, error) {
	return db.DeleteDonutHub(id), nil
}

// Hubs is the resolver for the hubs field.
func (r *queryResolver) Hubs(ctx context.Context) ([]*model.DonutHub, error) {
	return db.GetHubs(), nil
}

// Hub is the resolver for the hub field.
func (r *queryResolver) Hub(ctx context.Context, id string) (*model.DonutHub, error) {
	return db.GetHub(id), nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
