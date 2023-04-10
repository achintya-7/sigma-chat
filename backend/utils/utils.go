package utils

import (
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	MongoUri      string `mapstructure:"MONGO_URI"`
	DBName        string `mapstructure:"DB_NAME"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return config, err
}

func OSLoadConfig() (config Config, err error) {

	if MongoUri, ok := os.LookupEnv("MONGO_URI"); !ok {
		log.Fatal("MONGO_URI not set")
	} else {
		config.MongoUri = MongoUri
	}

	if DBName, ok := os.LookupEnv("DB_NAME"); !ok {
		log.Fatal("DB_NAME not set")
	} else {
		config.DBName = DBName
	}

	if ServerAddress, ok := os.LookupEnv("SERVER_ADDRESS"); !ok {
		log.Fatal("SERVER_ADDRESS not set")
	} else {
		config.ServerAddress = ServerAddress
	}

	log.Println("MONGO_URI: ", config.MongoUri)
	log.Println("DB_NAME: ", config.DBName)
	log.Println("SERVER_ADDRESS: ", config.ServerAddress)

	return config, err
}

func GetTime() int32 {
	return int32(time.Now().Unix())
}
