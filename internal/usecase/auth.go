package usecase

import (
	"errors"
	"log/slog"
	"toky/internal/domain/model"
	"toky/internal/domain/repository"
	"toky/internal/token"

	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	userRepo repository.User
	log *slog.Logger
	exp int64
}

func NewAuth(userRepo repository.User, log *slog.Logger, exp int64)*Auth{
	return &Auth{
		userRepo: userRepo,
		log: log,
		exp: exp,
	}
}

func (u *Auth) CreateUser(email, password string) (int64, error) {
	const op="internal/usecase/auth/CreateUser"
	log:=u.log.With(
		slog.String("operation", op),
		slog.String("email", email),
	)

	passHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Error(err.Error())
		return 0, err
	}

	userId, err := u.userRepo.Save(email, string(passHash))
	if err != nil {
		log.Error(err.Error())
		return 0, err
	}

	return userId, nil
}

func (u *Auth) GenerateToken(email, password string) (string, error) {
	const op="internal/usecase/auth/GenerateToken"
	log:=u.log.With(
		slog.String("operation", op),
		slog.String("email", email),
	)

	passHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	user, err := u.userRepo.User(email, string(passHash))
	if user == (model.User{}) {
		err:=errors.New("the user does not exist")

		log.Error(err.Error())

		return "", err
	}
	if err!=nil{
		log.Error(err.Error())
		return "", err
	}

	token, err := token.New(token.TokenClaims{
		Email: email,
		Exp:   u.exp,
	})

	return token, err
}

func (u *Auth) VerifyToken(tkn string) (string, error) {
	const op="internal/usecase/auth/VerifyToken"
	log:=u.log.With(
		slog.String("operation", op),
	)

	claims, err:=token.ExtractClaims(tkn)
	if err!=nil{
		log.Error(err.Error())
		return "", err
	}

	return claims.Email, nil
}
