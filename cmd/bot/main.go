package main

import (
	"fmt"
	"github.com/leveebreaks/fobattle/config"
	tb "github.com/tucnak/telebot"
	"time"
)

func main() {
	var env config.Env = config.DEV

	conf := config.Build(env)
	b, err := tb.NewBot(tb.Settings{
		Token:  conf.Token,
		Poller: &tb.LongPoller{Timeout: 2 * time.Second},
	})

	if err != nil {
		panic(err)
	}
	b.Start()

	fmt.Println(conf)
}
