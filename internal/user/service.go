package user

import (
	"context"
	"rest-app/package/logging"
)

type Service struct {
	storage Storage
	logger  *logging.Logger
}

func (s Service) Create(ctx context.Context, dto CreateUserDTO) (u User, err error) {
	//ToDO next one
	return
}
