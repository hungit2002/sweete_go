package config

import (
	"go.uber.org/zap"
	"sync"
)

var instance *EnvStruct
var once sync.Once

func GetInstance() *EnvStruct {
	once.Do(func() {
		if instance == nil {
			err := NewConfig()
			if err != nil {
				zap.S().Error(err)
			}
		}
	})
	return instance
}
