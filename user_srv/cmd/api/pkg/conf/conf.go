package conf

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)


func Init() {

}

// Conf 全局变量，用来保存应用的所有配置信息
var Conf = new(Config)


/*

	这里推荐使用 mapstructure 作为序列化标签

	yaml 不支持带有下划线的标签
		e.g.
			{
				AppSignExpire int64  `yaml:"app_sign_expire"` 
			}

 */

type Config struct {
	App    App    `mapstructure:"app"`
	JWT    JWT    `mapstructure:"jwt"`
	Redis  Redis  `mapstructure:"redis"`
	Consul Consul `mapstructure:"consul"`
	// 嵌套
}

type App struct {
	Addr string `mapstructure:"addr"`
}

type JWT struct {
	SecretKey string `mapstructure:"secret_key"`
}

type Redis struct {
	Addr string `mapstructure:"addr"`
}

type Consul struct {
	Addr      string `mapstructure:"addr"`
}


// 加载配置，失败直接 panic
func LoadConfig() {

	// 1. 创建 viper 实例
	viper := viper.New()

	// 2. 配置文件路径
	viper.SetConfigFile("./config.yaml")

	// 3. 配置读取
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	// 4. 将配置映射成结构体
	if err := viper.Unmarshal(Conf); err != nil {
		panic(err)
	}

	// 5. 监听配置文件变更，重新解析配置
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {  // 回调
		fmt.Println(e.Name)

		// Again，+1
		if err := viper.Unmarshal(Conf); err != nil {
			panic(err)
		}
	})
}