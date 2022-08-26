package api

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/lol-data/configs"
)

func GetFreeChampions() {
	configs.Load()
	url := fmt.Sprintf("https://br1.api.riotgames.com/lol/platform/v3/champion-rotations?api_key=" + configs.GetRiotAPIKey())

	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}
