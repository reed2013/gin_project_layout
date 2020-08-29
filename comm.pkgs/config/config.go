package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	vp *viper.Viper
}

func NewConfig(configPath string, configFileName string, configFileExt string) (*Config, error) {
	vp := viper.New()
	vp.SetConfigName(configFileName)
	vp.SetConfigType(configFileExt)
	vp.AddConfigPath(configPath)
	if err := vp.ReadInConfig(); err != nil {
		return nil, err
	}

	return &Config{vp}, nil
}

func (conf *Config) ReadSegment(key string, segment interface{}) error {
	if err := conf.vp.UnmarshalKey(key, segment); err != nil {
		return err
	}

	return nil
}