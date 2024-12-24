package initialize

import (
	"fmt"

	"github.com/spf13/viper"

	"github.com/augustus281/trackingcoin/global"
)

func LoadConfig() {
	viper := viper.New()
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to read configuration %w", err))
	}

	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Unable to decode configuration %v", err)
	}
}
