package configs

import (
	"github.com/reed2013/gin_project_layout/pkgs/config"
	"time"
)

type ServerConf struct {
	RunMode string
	HttpPort string
	ReadTimeout time.Duration
	WriteTimeout time.Duration
}

type AppConf struct {
	PageSize int
	LogPath string
	LogFileName string
	LogFileExt string
}

func NewConfig() (*config.Config, error) {
	var ConfigManager *config.Config
	var err error
	if ConfigManager, err = config.NewConfig("../configs/", "config", "yaml"); err != nil {
		return nil, err
	}

	return ConfigManager, nil
}