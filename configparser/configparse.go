package configparser

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// var (
// 	s3bucketname string
// 	s3region     string
// )

func ConfigParser() {

	viper.SetConfigName("config") // config file name without extension
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}

	s3bucketname := viper.GetString("awss3config.s3bucket")
	s3region := viper.GetString("awss3config.s3region")

	fmt.Println(s3bucketname)
	fmt.Println(s3region)

}
