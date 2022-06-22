package injector

import (
	"api-golang-hex/config"
	"log"
	"os"

	"github.com/spf13/viper"
)

func ConfigProvider() (*config.Global, error) {

	viper.SetConfigFile("./config.yml")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal(err)
	}

	var config config.Global
	err = viper.Unmarshal(&config)

	os.Setenv("JWTKey", config.BaseVariables.JWTKey)

	if err != nil {
		log.Fatal(err)
	}

	return &config, err
}
