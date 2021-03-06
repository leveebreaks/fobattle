package main

import (
	"github.com/leveebreaks/fobattle/internal/handlers"
	"github.com/leveebreaks/fobattle/internal/service"
	"github.com/spf13/viper"
	tb "gopkg.in/telebot.v3"
	"time"
)

var (
	botToken string
	b        *tb.Bot
)

func main() {
	setupConfig()

	initBot()

	setupHandlers()

	b.Start()
}

func setupConfig() {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	botToken = viper.GetString("bot.token")
}

func setupHandlers() {
	picksUrl := viper.GetString("fpl.api_urls.picks")
	standingsUrl := viper.GetString("fpl.api_urls.standings")
	lfs := service.NewService(picksUrl, standingsUrl) // TODO: USE INTERFACE!

	commandHandler := handlers.NewCommand(lfs)

	//b.Handle("/start", func(c tb.Context) error {
	//	s := service.LeagueFetchService{PicksUrl: viper.GetString("fpl.api_urls.picks")}
	//	p := s.Picks("4935817", "24")
	//	return c.Reply(p)
	//})
	b.Handle("/start", commandHandler.Standings)
	b.Handle("/picks", commandHandler.Picks)
}

func initBot() {
	var err error
	b, err = tb.NewBot(tb.Settings{
		Token:  botToken,
		Poller: &tb.LongPoller{Timeout: 2 * time.Second},
	})

	if err != nil {
		panic(err)
	}
}
