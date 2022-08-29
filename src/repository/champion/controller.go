package champion

import (
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lol-data/src/api"
	httperror "github.com/lol-data/src/utils/http"
)

func GetFreeChampions(w http.ResponseWriter, r *http.Request) {
	res := api.GetFreeChampions()

	body, err := ioutil.ReadAll(res.Body)

	httperror.RequestError(w, r, body, err)
}

func GetSummonerByName(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")

	res, err := api.GetSummonerByName(name)

	httperror.RequestError(w, r, res, err)
}
