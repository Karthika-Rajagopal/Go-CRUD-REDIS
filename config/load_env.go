package config

import(
	"github.com/spf13/viper"   //this package is to load the configuration from an environment variable file
)

type Config struct{
	//MySQL Setup, defines the configuration options for the application, which include MySQL and Redis setup details
	DBHost string `mapstructure:"SQL_HOST"`
	DBUsername string `mapstructure:"SQL_USER"`
	DBPassword string `mapstructure:"SQL_PASSWORD"`
	DBName string `mapstructure:"SQL_DB"`
	DBPort string `mapstructure:"SQL_PORT"`

	//Redis Setup
	RedisUrl string `mapstructure:"REDIS_URL"`

}

func LoadConfig(path string)(config Config, err error){  //path to the environment variable file as a parameter
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	//handle null
	err = viper.ReadInConfig()  /// package to read the configuration from the environment variable file, unmarshal it into the Config struct, and return it along with any errors
	if err!=nil{
		return
	}

	err = viper.Unmarshal(&config)
	return
}
