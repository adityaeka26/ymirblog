// Package rest is port handler.
package rest

import (
	"fmt"
	"net/http"

	"github.com/abialemuel/ymirblog/pkg/entity"
	"github.com/abialemuel/ymirblog/pkg/usecase/user"
	"github.com/go-chi/chi/v5"

	"github.com/abialemuel/ymirblog/pkg/persist/ymirblog"
	"github.com/kubuskotak/asgard/rest"
	pkgRest "github.com/kubuskotak/asgard/rest"
)

// UserOption is a struct holding the handler options.
type UserOption func(User *User)

// User handler instance data.
type User struct {
	UcUser user.T
	DB     *ymirblog.Database
}

// NewUser creates a new User handler instance.
//
//	var UserHandler = rest.NewUser()
//
//	You can pass optional configuration options by passing a Config struct:
//
//	var adaptor = &adapters.Adapter{}
//	var UserHandler = rest.NewUser(rest.WithUserAdapter(adaptor))
func NewUser(opts ...UserOption) *User {
	// Create a new handler.
	var handler = &User{}

	// Assign handler options.
	for o := range opts {
		var opt = opts[o]
		opt(handler)
	}

	// Return handler.
	return handler
}

func WithUserUsecase(u user.T) UserOption {
	return func(a *User) {
		a.UcUser = u
	}
}

// Register is endpoint group for handler.
func (u *User) Register(router chi.Router) {
	// PLEASE EDIT THIS EXAMPLE, how to register handler to router
	router.Post("/users", pkgRest.HandlerAdapter[CreateUserRequest](u.CreateUser).JSON)

}

// Create User handler
func (u *User) CreateUser(w http.ResponseWriter, r *http.Request) (GetUserResponse, error) {
	request, err := rest.GetBind[CreateUserRequest](r)

	if err != nil {
		fmt.Println(err, "error getbind")
		return GetUserResponse{}, rest.ErrBadRequest(w, r, err)
	}

	payload := entity.CreateUserPayload{
		Name:  request.Name,
		Email: request.Email,
	}

	//create user
	user, err := u.UcUser.CreateUser(r.Context(), payload)

	if err != nil {
		return GetUserResponse{
			Message: err.Error(),
		}, rest.ErrBadRequest(w, r, err)
	}

	return GetUserResponse{
		Message: "success",
		User:    &user,
	}, nil
}

