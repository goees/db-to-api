package dbtoapi

import "github.com/spf13/viper"

type Config struct {
	DBType         string
	DBServerIP     string
	DBServerPort   string
	DBName         string
	DBUsername     string
	DBPassword     string
	HttpServerPort string
}

var conf *Config

/*func init() {
	loadConfig()
}*/

func loadConfig() {
	//viper读取配置文件
	viper.SetConfigFile("config.yml")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	checkErr(err)

	conf = &Config{
		DBType:         viper.GetString("database.type"),
		DBServerIP:     viper.GetString("database.server-ip"),
		DBServerPort:   viper.GetString("database.server-port"),
		DBName:         viper.GetString("database.name"),
		DBUsername:     viper.GetString("database.username"),
		DBPassword:     viper.GetString("database.password"),
		HttpServerPort: viper.GetString("server.port"),
	}
}
