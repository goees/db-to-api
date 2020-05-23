package dbtoapi

import (
	"github.com/spf13/viper"
	"testing"
)

func TestServer(t *testing.T) {
	println("******************test server start******************")
	cfg := &Config{
		databaseType:       viper.GetString("database.type"),
		databaseServerIP:   viper.GetString("database.server-ip"),
		databaseServerPort: viper.GetString("database.server-port"),
		databaseName:       viper.GetString("database.name"),
		databaseUsername:   viper.GetString("database.username"),
		databasePassword:   viper.GetString("database.password"),
		httpServerPort:     viper.GetString("server.port"),
	}
	RunServer(cfg)
}
