package config

import (
	"fmt"
	"github.com/spf13/viper"
)

// 包级别的 Viper 实例变量
var devConfig *viper.Viper //开发测试用的配置文件
//var config *viper.Viper    //正式版用的配置文件

type Config struct {
	User   string
	Pass   string
	Addr   string
	Port   string
	Dbname string
}

func init() {
	// 初始化第一个 Viper 实例
	devConfig = viper.New()
	devConfig.SetConfigName("dev_config") //如果要切换成正式版，改这里就行
	devConfig.SetConfigType("json")
	devConfig.AddConfigPath("./config")
	if err := devConfig.ReadInConfig(); err != nil {
		fmt.Printf("Error reading dev_config file, %s\n", err)
	}

}

// 读区mysql配置
func GetMysqlConfig() Config {
	host := devConfig.GetString("mysql.host")
	port := devConfig.GetString("mysql.port")
	username := devConfig.GetString("mysql.username")
	password := devConfig.GetString("mysql.password")
	database := devConfig.GetString("mysql.database")
	myConfig := Config{
		User:   username,
		Pass:   password,
		Addr:   host,
		Port:   port,
		Dbname: database,
	}
	return myConfig
}

func GetVideoPath() string {
	return devConfig.GetString("video_files.path")
}

func Test() {
	// 设置配置文件名和存放路径

	// 读取配置文件
	if err := devConfig.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	// 获取并打印 MySQL 配置信息
	host := devConfig.GetString("mysql.host")
	port := devConfig.GetInt("mysql.port")
	username := devConfig.GetString("mysql.username")
	password := devConfig.GetString("mysql.password")
	database := devConfig.GetString("mysql.database")

	fmt.Println("MySQL Configuration:")
	fmt.Println("Host:", host)
	fmt.Println("Port:", port)
	fmt.Println("Username:", username)
	fmt.Println("Password:", password)
	fmt.Println("Database:", database)
}
func GetAIUrl() string {
	return devConfig.GetString("AI.url")
}
func GetReport() string {
	return devConfig.GetString("video_files.report_path")
}
