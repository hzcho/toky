//go:build wireinject
// +build wireinject

package injection

import (
	"log/slog"
	"toky/internal/handler"
	"toky/internal/repository"
	"toky/internal/usecase"

	domainr "toky/internal/domain/repository"
	domainuc "toky/internal/domain/usecase"

	"github.com/google/wire"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

var ProvideRepositories = wire.NewSet(
	repository.NewFileMetadata,
	wire.Bind(new(domainr.FileMetadata), new(repository.FileMetadata)),
	repository.NewFileStorage,
	wire.Bind(new(domainr.FileStorage), new(repository.FileStorage)),
	repository.NewUser,
	wire.Bind(new(domainr.User), new(repository.User)),
)

var ProvideUseCases = wire.NewSet(
	usecase.NewFile,
	wire.Bind(new(domainuc.File), new(usecase.File)),
	usecase.NewAuth,
	wire.Bind(new(domainuc.Auth), new(usecase.Auth)),
)

var ProvideHandlerGroup = wire.NewSet(
	handler.NewFileGroup,
	handler.NewAuthGroup,
)

var ProvideMiddleware = wire.NewSet(
	handler.NewMiddleware,
)
var ProvideHandler = wire.NewSet(
	handler.NewHandler,
)

func InitializeFile(e *echo.Echo, log *slog.Logger, db *pgxpool.Pool, uploadDir string, ttl int64) *handler.Handler {
	wire.Build(
		ProvideRepositories,
		ProvideUseCases,
		ProvideHandlerGroup,
		ProvideMiddleware,
		ProvideHandler,
	)
	return &handler.Handler{}
}
