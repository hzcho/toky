package app

import (
	"context"
	"fmt"
	"toky/internal/app/server"
	"toky/internal/injection"

	"log/slog"
	"toky/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

type App struct {
	server *server.Server
	pool   *pgxpool.Pool
}

func (a *App) Start() {
	a.server.Run()
}

func (a *App) Stop(ctx context.Context) {
	a.server.Shutdown(ctx)
	a.pool.Close()
}

func New(ctx context.Context, log *slog.Logger, cfg *config.Config) *App {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", cfg.DB.Username, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.DBName)

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		panic(err)
	}

	e := echo.New()

	_ = injection.InitializeFile(e, log, pool, cfg.UploadDir, int64(cfg.TokenTTL))

	server := server.New(cfg.Server, e)

	return &App{
		server: server,
		pool:   pool,
	}
}
