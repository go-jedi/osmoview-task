package app

import (
	"context"

	"github.com/go-jedi/osmoview-task/internal/config"
	"github.com/go-jedi/osmoview-task/pkg/logger"
)

type App struct {
	serverProvider *serverProvider
}

func NewApp(ctx context.Context) error {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServerProvider,
		a.initLogger,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	err := config.Load(".env")
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initServerProvider(_ context.Context) error {
	a.serverProvider = newServerProvider()
	return nil
}

func (a *App) initLogger(_ context.Context) error {
	logger.Init(
		logger.GetCore(
			logger.GetAtomicLevel(a.serverProvider.LoggerConfig().Level()),
		),
	)

	logger.Info("Logger is running")

	return nil
}
