package configparser

import (
	"fmt"
	"os"

	"example.com/fileUploadApp/logging"
	"github.com/spf13/viper"
)

func configParser() {

	viper.SetConfigName("config") // config file name without extension
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}

	logging.InfoLogger.Println("Viper configurations loaded successfully!")

}

func init() {
	configParser()
}
