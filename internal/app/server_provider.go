package app

import (
	"log"

	"github.com/go-jedi/osmoview-task/internal/config"
)

type serverProvider struct {
	loggerConfig config.LoggerConfig
}

func newServerProvider() *serverProvider {
	return &serverProvider{}
}

func (s *serverProvider) LoggerConfig() config.LoggerConfig {
	if s.loggerConfig == nil {
		cfg, err := config.NewLoggerConfig()
		if err != nil {
			log.Fatalf("failed to get logger config: %s", err.Error())
		}

		s.loggerConfig = cfg
	}

	return s.loggerConfig
}
