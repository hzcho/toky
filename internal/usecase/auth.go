package usecase

import (
	"toky/internal/domain/repository"
	"golang.org/x/crypto/bcrypt"
	"toky/internal/domain/model"
)

type Auth struct{
	UserRepo repository.UserRepository
}

func (u *Auth) CreateUser(user model.User) (int64, error){
	passHash, err:=bcrypt.GenerateFromPassword([]byte(user.Password),  bcrypt.DefaultCost)
	if err!=nil{
		return 0, err
	}

	userId, err:=u.UserRepo.Save(user.Name, string(passHash))
	if err!=nil{
		return 0, err
	}

	return userId, nil
}

func (u *Auth) GenerateToken(username, password string) (string, error){

}

func (u *Auth)VerifyToken(token string) (int, error){

}