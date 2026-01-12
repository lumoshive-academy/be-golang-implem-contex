package repository

import (
	"context"
	"log"
	"session-26/internal/dto"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepoInterface interface {
	GetUser(ctx context.Context) ([]dto.ResponseUser, error)
}

type UserRepo struct {
	DB *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) UserRepoInterface {
	return &UserRepo{DB: db}
}

func (userRepo *UserRepo) GetUser(ctx context.Context) ([]dto.ResponseUser, error) {
	query := `SELECT pg_sleep(10)`

	row, err := userRepo.DB.Query(ctx, query)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	var users []dto.ResponseUser
	for row.Next() {
		// mapping
	}

	return users, nil
}
