package domain

import (
	"kituwalk/config"

	"go.uber.org/zap"
)

type Apper interface {
	Options() *config.Configuration
	SaveOptions(key string, value interface{}) error
	SaveAllOptions() error
	Logger() *zap.SugaredLogger
	ConfigPath() string
	DefaultDbPath() string
	LogPath() string
	Pwd() string
	Output() string
}
