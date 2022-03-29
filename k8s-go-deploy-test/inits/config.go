package inits

import (
	"github.com/spf13/viper"
	"kgdt/configs"
)

func InitConfig() *configs.Configs{
	// 把配置文件读取到结构体上
	var configs configs.Configs
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./configs")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
		return nil
	}
	err = viper.Unmarshal(&configs)
	if err != nil {
		panic(err)
	}
	return &configs
}