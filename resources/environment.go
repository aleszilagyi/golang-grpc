package resources

import (
	"github.com/aleszilagyi/golang-grpc/config/logger"

	"github.com/spf13/viper"
)

var log *logger.StandardLogger

func init() {
	log = logger.NewLogger()
}

type AppEnv struct {
	AppConfig AppConfig `mapstructure:"config"`
}

type AppConfig struct {
	Hostname string `mapstructure:"hostname"`
	Port     string `mapstructure:"port"`
}

func GetConf() *AppEnv {

	viper.AddConfigPath("./resources")
	viper.SetConfigName("application")
	viper.SetConfigType("yml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Errorf("%v", err)
		panic(err)
	}

	conf := &AppEnv{}
	err = viper.UnmarshalKey("application", conf)
	if err != nil {
		log.Errorf("unable to decode into config struct, %v", err)
		panic(err)
	}
	return conf
}
