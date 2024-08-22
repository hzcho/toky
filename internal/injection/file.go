package injection

import (
	"log/slog"
	"toky/internal/handler"
	"toky/internal/repository"
	"toky/internal/usecase"

	"github.com/google/wire"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitializeFile(log *slog.Logger, db *pgxpool.Pool) *handler.Handler {
	wire.Bind(
		repository.NewFileMetadata,
		repository.NewFileStorage,
		repository.NewUser,
		usecase.NewFile,
		usecase.NewAuth,
		handler.NewHandler,
	)

	return &handler.Handler{}
}
