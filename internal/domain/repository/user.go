package repository

import "toky/internal/domain/model"

type UserRepository interface {
	User(name, passHash string) (model.User, error)
	Save(name, passHash string) (int64, error)
}
