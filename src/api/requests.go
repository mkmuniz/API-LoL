package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetItems() {
	url := "https://perodriguezl-league-of-legends-v1.p.rapidapi.com/lol/items?lang=en_US"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "86d88e9702msh84ebc415d46748bp196e77jsn45c9bcf71972")
	req.Header.Add("X-RapidAPI-Host", "perodriguezl-league-of-legends-v1.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}
