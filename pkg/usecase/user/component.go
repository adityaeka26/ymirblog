// Package user is implements component logic.
package user

import (
	"context"
	"fmt"
	"reflect"

	"entgo.io/ent/dialect"
	"github.com/abialemuel/ymirblog/pkg/adapters"
	"github.com/abialemuel/ymirblog/pkg/entity"
	"github.com/abialemuel/ymirblog/pkg/persist/ymirblog"
	"github.com/abialemuel/ymirblog/pkg/usecase"
)

func init() {
	usecase.Register(usecase.Registration{
		Name: "user",
		Inf:  reflect.TypeOf((*T)(nil)).Elem(),
		New: func() any {
			return &impl{}
		},
	})
}

// T is the interface implemented by all user Component implementations.
type T interface {
	CreateUser(ctx context.Context, newUser entity.CreateUserPayload) (entity.User, error)
}

type impl struct {
	adapter *adapters.Adapter
}

// Init initializes the execution of a process involved in a user Component usecase.
func (i *impl) Init(adapter *adapters.Adapter) error {
	i.adapter = adapter
	return nil
}

func WithYmirBlogPersist() adapters.Option {
	return func(a *adapters.Adapter) {
		// adapter ymirblog sqlite
		if a.YmirblogMySQL == nil {
			panic(fmt.Errorf("%s is not found", "YmirBlogMySQL"))
		}
		// persist ymirblog driver
		var c = ymirblog.Driver(
			ymirblog.WithDriver(a.YmirblogMySQL, dialect.MySQL),
		)
		a.YmirblogPersist = c
	}
}
