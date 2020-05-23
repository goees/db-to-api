package dbtoapi

import (
	"github.com/spf13/viper"
	"testing"
)

func TestServer(t *testing.T) {
	println("******************test server start******************")
	cfg := &Config{
		DBType:         viper.GetString("database.type"),
		DBServerIP:     viper.GetString("database.server-ip"),
		DBServerPort:   viper.GetString("database.server-port"),
		DBName:         viper.GetString("database.name"),
		DBUsername:     viper.GetString("database.username"),
		DBPassword:     viper.GetString("database.password"),
		HttpServerPort: viper.GetString("server.port"),
	}
	RunServer(cfg)
}
