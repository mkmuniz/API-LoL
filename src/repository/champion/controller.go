package champion

import (
	"net/http"

	"github.com/lol-data/src/api"
)

func GetFreeChampions(w http.ResponseWriter, r *http.Request) {
	api.GetFreeChampions()
}

func GetItemByName(w http.ResponseWriter, r *http.Request) {

}
