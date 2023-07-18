package user

import (
	"context"
	"errors"

	"github.com/abialemuel/ymirblog/pkg/entity"
	"github.com/kubuskotak/asgard/tracer"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel/trace"
)

func (i *impl) CreateUser(ctx context.Context, newUser entity.CreateUserPayload) (entity.User, error) {
	//create user
	entUser, err := i.adapter.YmirblogPersist.User.Create().
		SetName(newUser.Name).
		SetEmail(newUser.Email).
		Save(ctx)

	// mapping *entUser to entity.User
	res := entity.User{
		ID:    entUser.ID,
		Name:  entUser.Name,
		Email: entUser.Email,
	}
	return res, err
}

// get all user usecase
func (i *impl) GetAllUser(ctx context.Context) ([]entity.User, error) {
	span := trace.SpanFromContext(ctx)
	defer span.End()
	l := log.Hook(tracer.TraceContextHook(ctx))

	users, err := i.adapter.YmirblogPersist.User.Query().All(ctx)
	if err != nil {
		l.Error().Err(err).Msg("GetAll")
		return nil, err
	}

	getAllUser := []entity.User{}
	for _, user := range users {
		entityUser := entity.User{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		}
		getAllUser = append(getAllUser, entityUser)
	}

	return getAllUser, nil
}

// Get User By Id
func (i *impl) GetUserID(ctx context.Context, ID int) (entity.User, error) {
	span := trace.SpanFromContext(ctx)
	defer span.End()
	l := log.Hook(tracer.TraceContextHook(ctx))

	user, err := i.adapter.YmirblogPersist.User.Get(ctx, ID)
	if err != nil {
		l.Error().Err(err).Msg("GetBy ID")
		return entity.User{}, err
	}

	userID := entity.User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	return userID, nil
}

// Update User By Id
func (i *impl) UpdateUser(ctx context.Context, ID int, updateUser entity.UpdateUserPayload) (entity.User, error) {
	// Update User
	user, err := i.adapter.YmirblogPersist.User.UpdateOneID(ID).
		SetName(updateUser.Name).
		SetEmail(updateUser.Email).
		Save(ctx)
	if err != nil {
		return entity.User{}, err
	}

	// mapping *User to resUpdateUser
	resUpdateUser := entity.User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
	return resUpdateUser, err
}

// delete user
func (i *impl) DeleteUser(ctx context.Context, ID int) error  {
	span := trace.SpanFromContext(ctx)
	defer span.End()
	l := log.Hook(tracer.TraceContextHook(ctx))

	// validate persist connection
	if i.adapter.YmirblogPersist == nil {
		return errors.New("ymir blog persistence connection is nil")
	}

	err := i.adapter.YmirblogPersist.User.DeleteOneID(ID).Exec(ctx)
	if err != nil {
		l.Error().Err(err).Msg("Delete ID")
		return err
	}

	return err
}