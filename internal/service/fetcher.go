package service

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func NewService(picksUrl string, standingsUrl string) LeagueFetchService {
	return LeagueFetchService{PicksUrl: picksUrl, StandingsUrl: standingsUrl}
}

type LeagueFetchService struct {
	PicksUrl     string
	StandingsUrl string
}

func (s *LeagueFetchService) Picks(managerId string, gameWeek string) string {
	replacer := strings.NewReplacer("{manager_id}", managerId, "{event_id}", gameWeek)
	url := replacer.Replace(s.PicksUrl)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Add("User-Agent", "TelebotFetcher")
	if err != nil {
		log.Println("LeagueFetchService", "Picks", err)
	}
	client := &http.Client{}
	response, err := client.Do(request)
	defer response.Body.Close()
	respBody, _ := ioutil.ReadAll(response.Body)
	return string(respBody)
}

func (s *LeagueFetchService) Standings(leagueId string) []byte {
	replacer := strings.NewReplacer("{league_id}", leagueId)
	url := replacer.Replace(s.StandingsUrl)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Add("User-Agent", "TelebotFetcher")
	if err != nil {
		log.Println("LeagueFetchService", "Picks", err)
	}
	client := &http.Client{}
	response, err := client.Do(request)
	defer response.Body.Close()
	respBody, _ := ioutil.ReadAll(response.Body)
	return respBody
}
