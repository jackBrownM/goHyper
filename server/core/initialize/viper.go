package initialize

import (
	"fmt"
	"github.com/spf13/viper"
	"goHyper/core/global"
)

func ViperInit() {
	viper.SetConfigFile("./resource/config.yaml") // 指定 YAML 文件路径
	if err := viper.ReadInConfig(); err != nil {
		global.HyperLog.Error(fmt.Sprintf("Error reading config file, %s", err))
	}

	// 使用 viper 获取配置
	port := viper.GetInt("server.port")
	host := viper.GetString("server.host")

	fmt.Printf("Server will run on %s:%d\n", host, port)
}
