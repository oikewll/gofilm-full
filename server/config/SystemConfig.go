package config

import (
	"github.com/spf13/viper"
)

var (
	IsDev bool
)

func init() {
	IsDev = viper.GetBool("system.development")
}
