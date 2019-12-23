//go:generate go run github.com/99designs/gqlgen -v

package graphql

import (
	"context"
	"errors"
	"github.com/tom-klimovski-kappa/graphql/models"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

var meetups = []*models.Meetup{
	{
		ID:          "1",
		Name:        "A Meetup",
		Description: "A description",
		UserID:      "1",
	},
	{
		ID:          "2",
		Name:        "A Second Meetup",
		Description: "A second description",
		UserID:      "2",
	},
}

var users = []*models.User{
	{
		ID:       "1",
		Username: "tom",
		Email:    "bob@gmail.com",
	},
	{
		ID:       "2",
		Username: "alice",
		Email:    "alice@gmail.com",
	},
}

// Resolver is a comment here
type Resolver struct{}

type queryResolver struct{ *Resolver }

type mutationResolver struct{ *Resolver }

type meetupResolver struct{ *Resolver }

type userResolver struct{ *Resolver }

//Query func for resolver struct
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{Resolver: r}
}

//Mutation should have stuff here
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
	user := new(models.User)

	for _, u := range users {
		if u.ID == obj.UserID {
			user = u
			break
		}
	}

	if user == nil {
		return nil, errors.New("user id does not exist")
	}

	return user, nil
}

func (u *userResolver) Meetups(ctx context.Context, obj *models.User) ([]*models.Meetup, error) {
	var m []*models.Meetup

	for _, meetup := range meetups {
		if meetup.UserID == obj.ID {
			m = append(m, meetup)
		}
	}

	return m, nil
}

//Meetup resolver
func (r *Resolver) Meetup() MeetupResolver {
	return &meetupResolver{ Resolver: r}
}

//User resolver
func (r *Resolver) User() UserResolver {
	return &userResolver{ Resolver: r}
}

