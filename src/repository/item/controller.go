package item

import (
	"net/http"

	"github.com/lol-data/src/api"
)

func GetAllItems(w http.ResponseWriter, r *http.Request) {
	api.GetItems()
}

func GetItemByName(w http.ResponseWriter, r *http.Request) {

}
