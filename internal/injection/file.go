package injection

import (
	"log/slog"
	"toky/internal/handler"
	"toky/internal/repository"
	"toky/internal/usecase"

	"github.com/google/wire"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ProvideRepositories = wire.NewSet(
	repository.NewFileMetadata,
	repository.NewFileStorage,
	repository.NewUser,
)

var ProvideUseCases = wire.NewSet(
	usecase.NewFile,
	usecase.NewAuth,
)

var ProvideHandler = wire.NewSet(
	handler.NewHandler,
)

func InitializeFile(log *slog.Logger, db *pgxpool.Pool) *handler.Handler {
	wire.Build(
		ProvideRepositories,
		ProvideUseCases,
		ProvideHandler,
	)
	return &handler.Handler{}
}
