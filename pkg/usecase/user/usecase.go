package user

import (
	"context"

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

	// mapping *ent.User to entity.User
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