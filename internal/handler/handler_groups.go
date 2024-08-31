package handler

import(
	"toky/internal/usecase"
)

type Groups struct{
	AuthGroup
	FileGroup
	Middleware
}

func NewGroups(usecases *usecase.Usecases)*Groups{
	return &Groups{
		AuthGroup: NewAuthGroup(usecases.Auth),
		FileGroup: NewFileGroup(usecases.File),
		Middleware: NewMiddleware(usecases.Auth),
	}
}