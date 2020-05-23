package dbtoapi

import "github.com/spf13/viper"

type Config struct {
	databaseType       string
	databaseServerIP   string
	databaseServerPort string
	databaseName       string
	databaseUsername   string
	databasePassword   string
	httpServerPort     string
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
		databaseType:       viper.GetString("database.type"),
		databaseServerIP:   viper.GetString("database.server-ip"),
		databaseServerPort: viper.GetString("database.server-port"),
		databaseName:       viper.GetString("database.name"),
		databaseUsername:   viper.GetString("database.username"),
		databasePassword:   viper.GetString("database.password"),
		httpServerPort:     viper.GetString("server.port"),
	}
}
