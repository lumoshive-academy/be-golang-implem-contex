package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	UserRepo UserRepoInterface
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		UserRepo: NewUserRepo(db),
	}
}
