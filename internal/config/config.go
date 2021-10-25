package config

import (
	"github.com/spf13/viper"
	"sync"
)

type cfg struct {
	NeoApiKey string
}

var doOnce sync.Once
var config *cfg

func Get() *cfg {
	doOnce.Do(func() {
		config = &cfg{}
		viper.AutomaticEnv()
		config.NeoApiKey = viper.GetString("RIOT_API_KEY")
	})
	return config
}
