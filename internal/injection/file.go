//go:build wireinject
// +build wireinject

package injection

import (
	"log/slog"
	"toky/internal/handler"
	"toky/internal/repository"
	"toky/internal/usecase"

	"github.com/google/wire"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
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

var ProvideHandlerGroup = wire.NewSet(
	handler.NewFileGroup,
	handler.NewFileGroup,
)
var ProvideHandler = wire.NewSet(
	handler.NewHandler,
)

func InitializeFile(e *echo.Echo, log *slog.Logger, db *pgxpool.Pool) *handler.Handler {
	wire.Build(
		ProvideRepositories,
		ProvideUseCases,
		ProvideHandler,
	)
	return &handler.Handler{}
}
