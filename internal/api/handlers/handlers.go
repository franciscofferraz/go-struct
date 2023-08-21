package handlers

import "github.com/franciscofferraz/go-struct/internal/db/repositories"

type Handlers struct {
	UserHandler *UserHandler
}

func NewHandlers(ur *repositories.UserRepository) *Handlers {
	return &Handlers{
		UserHandler: NewUserHandler(ur),
	}
}
