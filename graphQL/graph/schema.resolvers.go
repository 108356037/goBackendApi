package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"graphql/example.com/graph/generated"
	"graphql/example.com/graph/model"
	"graphql/example.com/internal/links"
	"graphql/example.com/internal/users"
	"graphql/example.com/pkg/jwt"
	"strconv"
)

func (r *mutationResolver) CreateLink(ctx context.Context, input model.NewLink) (*model.Link, error) {
	var link links.Link
	link.Address = input.Address
	link.Title = input.Title
	linkID, err := link.Save()
	if err != nil {
		return nil, err
	}
	return &model.Link{ID: strconv.FormatInt(linkID, 10), Title: link.Title, Address: link.Address}, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	var user users.User
	user.Username = input.Username
	user.Password = input.Password
	if err := user.Create(); err != nil {
		return "", err
	}
	token, err := jwt.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Links(ctx context.Context) ([]*model.Link, error) {
	resultLinks := []*model.Link{}
	dbLinks, err := links.GetAll()
	if err != nil {
		return nil, err
	}
	for _, val := range dbLinks {
		resultLinks = append(resultLinks, &model.Link{
			ID:      val.ID,
			Title:   val.Title,
			Address: val.Address,
		})
	}
	return resultLinks, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
