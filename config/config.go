package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var (
	cfg Config
)

// Config
// @Description: stores telegram bot config and
type Config struct {
	TgCfg
	CommonCfg
}

type TgCfg struct {
	Token string
}

type CommonCfg struct {
	LogPath string
}

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./config/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fail to get config.toml: %v", err))
	}
	cfg.Token = viper.GetString("telegram.token")
	cfg.LogPath = viper.GetString("common.logPath")
}

// GetConfig
//
//	@Description: Get config instance
//	@return Config
func GetConfig() Config {
	return cfg
}
