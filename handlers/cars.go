package handlers

import (
	"encoding/json"
	"net/http"

	db "github.com/era-n/nic-test/database"
)

var database *db.Db

func init() {
	database = db.NewDB()
	database.InitMongo()
}

func GetCarsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	cars := database.GetCars()

	resp, err := json.Marshal(cars)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(resp)
}
