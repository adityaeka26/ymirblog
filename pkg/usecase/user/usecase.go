package user

import (
	"context"

	"github.com/abialemuel/ymirblog/pkg/entity"
)

func (i *impl) CreateUser(ctx context.Context, newUser entity.CreateUserPayload) (entity.User, error) {
	//create user
	entUser, err := i.adapter.YmirblogPersist.User.Create().
		SetName(newUser.Name).
		SetEmail(newUser.Email).
		Save(ctx)

	// mapping *ent.User to entity.User
	res := entity.User{
		ID:    entUser.ID,
		Name:  entUser.Name,
		Email: entUser.Email,
	}
	return res, err
}
