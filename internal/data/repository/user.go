package repository

import (
	"context"
	"fmt"
	"log"
	"session-26/internal/dto"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepoInterface interface {
	GetUser(ctx context.Context) (dto.ResponseUser, error)
}

type UserRepo struct {
	DB *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) UserRepoInterface {
	return &UserRepo{DB: db}
}

func (userRepo *UserRepo) GetUser(ctx context.Context) (dto.ResponseUser, error) {
	query := `SELECT name, email FROM users WHERE id = 1 AND pg_sleep(10) IS NULL`

	reqID := ctx.Value("ctxid").(string)
	fmt.Println(reqID)

	var user dto.ResponseUser
	err := userRepo.DB.QueryRow(ctx, query).Scan(&user.Name, &user.Email)
	if err != nil {
		log.Println(err)
		return user, err
	}

	return user, nil
}
