package config

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/viper"
)

type Info struct {
	Instance string `json:"instance"`
	Version  string `json:"version"`
}

var Params ConfigParams

var Version string

type ConfigParams struct {
	Logger struct {
		Logfile string `json:"logfile"`
	} `json:"logger"`
	Mongo struct {
		Database string `json:"database"`
		Host     string `json:"host"`
		Password string `json:"password"`
		Port     string `json:"port"`
		User     string `json:"user"`
	} `json:"mongo"`
	Port string `json:"port"`
}

func init() {
	fmt.Printf("%v\n", Version)

	viper.AddConfigPath("./configs")
	viper.AddConfigPath("$HOME/configs")
	// And then register config file name (no extension)
	viper.SetConfigName("env")

	//Register Environment Prefix
	viper.SetEnvPrefix("APP")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	// Optionally we can set specific config type
	//viper.SetConfigType("json")

	viper.AutomaticEnv()

	// Find and read the config file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	// Confirm which config file is used
	fmt.Printf("Using config: %s\n", viper.ConfigFileUsed())

	// Check if a particular key is set
	// Notice that we can trverse nested configuration e.g. prod.port
	if !viper.IsSet("port") {
		log.Fatal("missing port number")
	}

	// https://godoc.org/github.com/spf13/viper#Get
	// Get can retrieve any value given the key to use.
	// Get is case-insensitive for a key.
	// Get has the behavior of returning the value associated with the first place from where it is set.
	// Viper will check in the following order: override, flag, env, config file, key/value store, default
	// Get returns an interface. For a specific value use one of the Get____ methods.

	port := viper.Get("port") // returns string
	//port := viper.GetInt("prod.port") // returns integer
	fmt.Printf("GET port Value: %v, Type: %T\n", port, port)

	// Extract sub-tree using `Sub`

	//mongo := viper.Sub("mongo")

	//fmt.Printf("%v \n", mongo)

	err := viper.Unmarshal(&Params)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	fmt.Printf("mongo.host Value: %v, Type: %T\n", Params.Mongo.Host, Params.Mongo.Host)
	fmt.Printf("mongo.port Value: %v, Type: %T\n", Params.Mongo.Port, Params.Mongo.Port)
	fmt.Printf("mongo.database Value: %v, Type: %T\n", Params.Mongo.Database, Params.Mongo.Database)

	if viper.IsSet("mongo.user") {
		Params.Mongo.User = fmt.Sprintf("%v", viper.Get("mongo.user"))
		fmt.Printf("mongo.user Value: %v, Type: %T\n", viper.Get("mongo.user"), viper.Get("mongo.user"))

	}

	if viper.IsSet("mongo.password") {
		Params.Mongo.Password = fmt.Sprintf("%v", viper.Get("mongo.password"))
		fmt.Printf("mongo.password Value: %v, Type: %T\n", viper.Get("mongo.password"), viper.Get("mongo.password"))
	}

}
