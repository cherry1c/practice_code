package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("conf")
	viper.SetConfigType("json")
	viper.AddConfigPath("./conf")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("config file not found.")
			return
		}
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	f := viper.Get("local_life.mysql.locallife")
	fmt.Println(f)
	fmt.Println(viper.Get("local_life.mysql.locallife.username"))

	notExist := viper.Get("not exist")
	fmt.Println(notExist)
}
