package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"context"
	"toky/internal/domain/model"
)

type User struct{
	db *pgxpool.Pool
}

func NewUser(db *pgxpool.Pool)*User{
	return &User{
		db: db,
	}
}

func (r *User) User(ctx context.Context, name, passHash string) (model.User, error){
	var user model.User

	query:="select * from users where name=$1, pass_hash=$2"

	if err:=r.db.QueryRow(ctx, query, name, passHash).Scan(&user); err!=nil{
		return model.User{}, err
	}

	return user, nil
}

func (r *User) Save(ctx context.Context, name, passHash string) (uint64, error){
	var userID uint64

	query:="insert into users (name, pass_hash) values($1, $2) returning id"

	if err:=r.db.QueryRow(ctx, query, name, passHash).Scan(&userID); err!=nil{
		return 0, err
	}

	return userID, nil
}