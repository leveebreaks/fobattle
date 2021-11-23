package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Env string

type builder func() Config

const (
	DEV  Env = "dev"
	PROD Env = "prod"
)

var builders map[Env]builder

type Config struct {
	Token string `json:"bot_token"`
}

func init() {
	builders = make(map[Env]builder)
	builders[DEV] = buildDev
}

func Build(env Env) Config {
	return builders[env]()
}

func buildDev() Config {
	file, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	c := Config{}
	err = json.Unmarshal(bytes, &c)
	if err != nil {
		panic(err)
	}

	return c
}
