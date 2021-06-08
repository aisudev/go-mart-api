package utils

import (
	"strings"

	"github.com/spf13/viper"
)

func ViperInit() {
	viper.SetConfigName("config")
	viper.AddConfigPath("config")
	viper.AutomaticEnv()

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
