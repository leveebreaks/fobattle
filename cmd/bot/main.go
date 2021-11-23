package main

import (
	"fmt"
	"github.com/leveebreaks/fobattle/config"
)

func main() {
	var env config.Env = config.DEV

	conf := config.Build(env)

	fmt.Println(conf)

}
