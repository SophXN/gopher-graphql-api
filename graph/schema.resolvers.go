package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/sophxn/hackernews/graph/generated"
	"github.com/sophxn/hackernews/graph/model"
)

// CreateLink is the resolver for the createLink field.
func (r *mutationResolver) CreateLink(ctx context.Context, input model.NewLink) (*model.Link, error) {
	var link model.Link
	var user model.User
	link.Address = input.Address
	link.Title = input.Title
	user.Name = "test"
	link.User = &user
	return &link, nil

	//OR can also be written like this
	// var link model.Link
	// var user model.User
	// link := model.Link{
	// 	Title:   input.Title,
	// 	Address: input.Address,
	// 	User: &model.User{
	// 		Name: "test user",
	// 	},
	// }
	// return &link, nil

}

// CreateUSer is the resolver for the createUSer field.
func (r *mutationResolver) CreateUSer(ctx context.Context, input model.NewUser) (string, error) {
	panic(fmt.Errorf("not implemented: CreateUSer - createUSer"))
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	panic(fmt.Errorf("not implemented: Login - login"))
}

// RefreshToken is the resolver for the refreshToken field.
func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented: RefreshToken - refreshToken"))
}

// Links is the resolver for the links field.
// It takes the context (for the current user) and returns slice of Links and error(IF Applicable)
func (r *queryResolver) Links(ctx context.Context) ([]*model.Link, error) {
	var links []*model.Link
	dummyLink := model.Link{
		Title:   "our dummy dummy link",
		Address: "https://address.org",
		User:    &model.User{Name: "admin"},
	}
	links = append(links, &dummyLink)
	return links, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
