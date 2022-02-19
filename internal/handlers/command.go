package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/leveebreaks/fobattle/internal/service"
	"github.com/spf13/viper"
	tb "gopkg.in/telebot.v3"
	"log"
	"strconv"
	"time"
)

type Command struct {
	lfs service.LeagueFetchService
}

func NewCommand(lfs service.LeagueFetchService) Command {
	return Command{lfs: lfs}
}

func (cmd Command) Picks(c tb.Context) error {
	return c.Send(cmd.lfs.PicksUrl)
}

func (cmd Command) Standings(c tb.Context) error {
	leagueId := getLeagueId()
	standingsJson := cmd.lfs.Standings(leagueId)
	var s standings
	err := json.Unmarshal(standingsJson, &s)
	if err != nil {
		log.Println("handlers", "Standings", err)
	}

	var msg string
	for _, v := range s.Standings.Results {
		msg += fmt.Sprintf("%v. %v | %v | %v \n", strconv.Itoa(v.Rank), v.PlayerName, v.EntryName, strconv.Itoa(v.Total))
	}

	return c.Send(msg)
}

func getLeagueId() string {
	return viper.GetString("fpl.league_id")
}

type standings struct {
	NewEntries struct {
		HasNext bool          `json:"has_next"`
		Page    int           `json:"page"`
		Results []interface{} `json:"results"`
	} `json:"new_entries"`
	LastUpdatedData time.Time `json:"last_updated_data"`
	League          struct {
		Id          int         `json:"id"`
		Name        string      `json:"name"`
		Created     time.Time   `json:"created"`
		Closed      bool        `json:"closed"`
		MaxEntries  interface{} `json:"max_entries"`
		LeagueType  string      `json:"league_type"`
		Scoring     string      `json:"scoring"`
		AdminEntry  int         `json:"admin_entry"`
		StartEvent  int         `json:"start_event"`
		CodePrivacy string      `json:"code_privacy"`
		HasCup      bool        `json:"has_cup"`
		CupLeague   interface{} `json:"cup_league"`
		Rank        interface{} `json:"rank"`
	} `json:"league"`
	Standings struct {
		HasNext bool `json:"has_next"`
		Page    int  `json:"page"`
		Results []struct {
			Id         int    `json:"id"`
			EventTotal int    `json:"event_total"`
			PlayerName string `json:"player_name"`
			Rank       int    `json:"rank"`
			LastRank   int    `json:"last_rank"`
			RankSort   int    `json:"rank_sort"`
			Total      int    `json:"total"`
			Entry      int    `json:"entry"`
			EntryName  string `json:"entry_name"`
		} `json:"results"`
	} `json:"standings"`
}
