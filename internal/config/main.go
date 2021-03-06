package config

import (
	"github.com/spf13/viper"
	"log"
	"strings"
)

var C config

type (
	config struct {
		Server struct {
			URL string
		}
		Database struct {
			Userdb struct {
				URL         string
				Dialect     string
				Poolsizemin int
				Poolsizemax int
			}
			Docsdb struct {
				URL         string
				Dialect     string
				Poolsizemin int
				Poolsizemax int
			}
		}
	}
)

func InitConfig() {
	viper.SetConfigFile("config.yaml")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	err := viper.Unmarshal(&C)
	if err != nil {
		log.Fatalf("unmarshal config file error - %s", err)
	}
}
