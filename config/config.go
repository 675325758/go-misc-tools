package config




import (
	"fmt"

	"github.com/spf13/viper"
)


type Config struct {
	Redis string
	MySQL MySQLConfig
}

type MySQLConfig struct {
	Port     int
	Host     string
	Username string
	Password string
}

func test() {
	var config Config
	write_viper := viper.New()
	write_viper.SetConfigName("config")
	write_viper.SetConfigType("json")
	write_viper.AddConfigPath("/tmp")


	write_viper.Set("redis", "this is a example of json")
	write_viper.Set("mysql.port", 3306)
	write_viper.Set("mysql.host", "192.168.1.0")
	write_viper.Set("mysql.username", "root123")
	write_viper.Set("mysql.password", "root123")
	
	if err := write_viper.WriteConfig(); err != nil {
		fmt.Println(err)
	}

	read_viper := viper.New()
	read_viper.SetConfigName("config")
	read_viper.SetConfigType("json")
	read_viper.AddConfigPath("/tmp")
	if err := read_viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		return
	}
	read_viper.Unmarshal(&config)
	fmt.Println("config: ", config)
}

func Test() {
	go test()
}