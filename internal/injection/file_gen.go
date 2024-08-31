package injection

import (
	"log/slog"
	"toky/internal/handler"
	"toky/internal/repository"
	"toky/internal/usecase"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

func InitializeFile(e *echo.Echo, log *slog.Logger, db *pgxpool.Pool, uploadDir string, ttl int64) *handler.Handler {
	repo := repository.NewRepositories(db, uploadDir)
	usecases := usecase.NewUsecases(repo, log, ttl)
	groups := handler.NewGroups(usecases)
	handler := handler.NewHandler(e, groups.FileGroup, groups.AuthGroup, groups.Middleware)

	return handler
}
