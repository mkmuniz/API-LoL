package api

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/lol-data/configs"
)

func GetFreeChampions() *http.Response {
	configs.Load()
	url := fmt.Sprintf("https://br1.api.riotgames.com/lol/platform/v3/champion-rotations?api_key=" + configs.GetRiotAPIKey())

	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)

	return res
}

func GetSummonerByName(name string) ([]byte, error) {
	configs.Load()
	url := fmt.Sprintf("https://br1.api.riotgames.com/lol/summoner/v4/summoners/by-name/" + name + "?api_key=" + configs.GetRiotAPIKey())

	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)

	body, err := ioutil.ReadAll(res.Body)

	return body, err
}
