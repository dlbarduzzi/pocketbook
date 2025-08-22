package registry

import "github.com/spf13/viper"

var (
	configPath = "."
	configType = "env"
	configName = ".env"
)

func NewRegistry() (*viper.Viper, error) {
	reg := viper.New()
	reg.AutomaticEnv()

	reg.AddConfigPath(configPath)
	reg.SetConfigType(configType)
	reg.SetConfigName(configName)

	if err := reg.ReadInConfig(); err != nil {
		return nil, err
	}

	return reg, nil
}

func SetConfigPath(s string) {
	configPath = s
}

func SetConfigType(s string) {
	configType = s
}

func SetConfigName(s string) {
	configName = s
}
