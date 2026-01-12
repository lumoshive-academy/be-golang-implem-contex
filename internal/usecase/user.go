package usecase

import (
	"context"
	"log"
	"session-26/internal/data/repository"
	"session-26/internal/dto"
)

type UserUseCaseInterface interface {
	GetUser(ctx context.Context) ([]dto.ResponseUser, error)
}

type UserUseCase struct {
	UserRepo repository.Repository
}

func NewUserUseCase(userRepo repository.Repository) UserUseCaseInterface {
	return &UserUseCase{UserRepo: userRepo}
}

func (userUseCase *UserUseCase) GetUser(ctx context.Context) ([]dto.ResponseUser, error) {
	users, err := userUseCase.UserRepo.UserRepo.GetUser(ctx)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return users, nil
}
