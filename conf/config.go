package conf

import (
	"os"

	"gopkg.in/yaml.v2"
)

type config struct {
	DataSource *datasource `yaml:"datasource"`
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

func init() {
	conf = &config{}
	if f, err := os.Open("../app.yaml"); err != nil {
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
