//go:generate go run github.com/99designs/gqlgen -v

package graphql

import (
	"context"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

var meetups = []*models.Meetup{
	{
		ID: "1",
		Name: "A Meetup",
		Description: "A description",
		UserID: "1",
	},
	{
		ID: "2",
		Name: "A Second Meetup",
		Description: "A second description",
		UserID: "2",
	},
}

var users = []*models.User{
	{
		ID: "1",
		Username: "tom",
		Email: "bob@gmail.com",
	},
	{
		ID: "2",
		Username: "alice",
		Email: "alice@gmail.com",
	},
}

type Resolver struct{}

type queryResolver struct{ *Resolver }

type mutationResolver struct { *Resolver }

type meetupResolver struct { *Resolver }

type userResolver struct { *Resolver }

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{Resolver: r}
}

func (r *queryResolver) Meetups(ctx context.Context) ([]*models.Meetup, error) {
	return meetups, nil
}

func (m *mutationResolver) CreateMeetup(ctx context.Context, input NewMeetup) (*models.Meetup, error) {
	panic("not implemented")
}

func (m *meetupResolver) User(ctx context.Context, obj *models.Meetup) (*models.User, error) {
	panic("not implemented")
}

func (u *userResolver) Meetups(ctx context.Context, obj *models.Meetup) (*models.User, error){
	return nil, nil
}