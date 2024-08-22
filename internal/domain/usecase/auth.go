package usecase

type Auth interface {
	CreateUser(username, password string) (int, error)
	GenerateToken(username, password string) (string, error)
	VerifyToken(token string) (int, error)
}
