package viper

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func ReadFile[T any](pathFile string, nameFile string, typeFile string) (*T, error) {
	viper.AddConfigPath(pathFile)
	viper.SetConfigName(nameFile)
	viper.SetConfigType(typeFile)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		zap.S().Error(err)
		return nil, err
	}

	var result T
	err = viper.Unmarshal(&result)
	if err != nil {
		zap.S().Error(err)
		return nil, err
	}
	return &result, nil
}
