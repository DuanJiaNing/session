package conf

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type config struct {
	DataSource *datasource `yaml:"datasource"`
	GRpc       *grpc       `yaml:"grpc"`
}

type grpc struct {
	ListenAddress string `yaml:"listen-address"`
}

type datasource struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

var (
	conf *config
)

func Init(path string) {
	fmt.Println("parse config file at: ", path)

	conf = &config{}
	if f, err := os.Open(path); err != nil {
		panic(err)
	} else {
		if err = yaml.NewDecoder(f).Decode(conf); err != nil {
			panic(err)
		}
	}
}

func DataSource() *datasource {
	return conf.DataSource
}

func GRpc() *grpc {
	return conf.GRpc
}
