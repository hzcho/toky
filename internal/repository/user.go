package repository

import (
	"context"
	"toky/internal/domain/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	db *pgxpool.Pool
}

func NewUser(db *pgxpool.Pool) *User {
	return &User{
		db: db,
	}
}

func (r *User) User(ctx context.Context, email string) (model.User, error) {
	var user model.User

	query := "select id, email, pass_hash from users where email=$1"

	if err := r.db.QueryRow(ctx, query, email).Scan(&user.Id, &user.Email, &user.Password); err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (r *User) Save(ctx context.Context, email, passHash string) (uint64, error) {
	var userID uint64

	query := "insert into users (email, pass_hash) values($1, $2) returning id"

	if err := r.db.QueryRow(ctx, query, email, passHash).Scan(&userID); err != nil {
		return 0, err
	}

	return userID, nil
}
