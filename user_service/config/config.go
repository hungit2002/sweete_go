package config

import (
	"fmt"
	"go.uber.org/zap"
	"os"
	mysqlConfig "sweete/core/pkg/mysql/config"
	"sweete/core/pkg/viper"
)

type EnvStruct struct {
	App             appStruct             `mapstructure:"app"`
	MySql           *mysqlConfig.Config   `mapstructure:"mysql"`
	Storage         storageStruct         `mapstructure:"storage"`
	ExternalService externalServiceStruct `mapstructure:"externalService"`
}
type storageStruct struct {
	Disk string `mapstructure:"disk"`
}

type Curl struct {
	ClassroomUrl string `mapstructure:"classroomUrl"`
	CourseUrl    string `mapstructure:"courseUrl"`
	UserAdminUrl string `mapstructure:"userAdminUrl"`
}
type Sms struct {
	ServiceEmail string `mapstructure:"serviceEmail"`
}

type Google struct {
	ClientID        string `mapstructure:"clientID"`
	ClientSecret    string `mapstructure:"clientSecret"`
	AuthStateString string `mapstructure:"authStateString"`
}

type appStruct struct {
	Port          string `mapstructure:"port"`
	Name          string `mapstructure:"name"`
	Env           string `mapstructure:"env"`
	Key           string `mapstructure:"key"`
	Timezone      string `mapstructure:"timezone"`
	SecretKey     string `mapstructure:"secretKey"`
	TokenToServer string `mapstructure:"tokenToServer"`
	GrpcPort      string `mapstructure:"grpcPort"`
	Url           string `mapstructure:"url"`
	Curl          Curl   `mapstructure:"curl"`
	Sms           Sms    `mapstructure:"sms"`
	Google        Google `mapstructure:"google"`
}

type externalServiceStruct struct {
	Develop string `mapstructure:"develop"`
	Media   string `mapstructure:"media"`
	Product string `mapstructure:"product"`

	App      string `mapstructure:"app"`
	User     string `mapstructure:"user"`
	Award    string `mapstructure:"award"`
	Lesson   string `mapstructure:"lesson"`
	Story    string `mapstructure:"story"`
	Platform string `mapstructure:"platform"`

	Report string `mapstructure:"report"`

	Sentry       string `mapstructure:"sentry"`
	Telegram     string `mapstructure:"telegram"`
	ServiceEmail string `mapstructure:"serviceEmail"`
}

func NewConfig() error {
	path, err := os.Getwd()
	if err != nil {
		return err
	}
	path = fmt.Sprintf("%s/user_service/config", path)
	env, err := viper.ReadFile[EnvStruct](path, "config", "yaml")
	if err != nil {
		zap.S().Error(err)
		return err
	}

	instance = env
	return nil
}
