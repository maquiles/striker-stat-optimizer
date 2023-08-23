package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

type StatPriorityList struct {
	Stats []string `json:"stats"`
}

func (app *App) HealthHandler(writer http.ResponseWriter, request *http.Request) {
	log.Println("Received GET /health")
	writer.Write([]byte("VIBEY AF"))
}

func (app *App) CalculateHandler(writer http.ResponseWriter, request *http.Request) {
	log.Println("Received GET /calculate")
	CalcAllStrikerGearStatCombos(app.DB)
	writer.Write([]byte("Striker gear combo calculations completed successfully!"))
}

func (app *App) StrikerHandler(writer http.ResponseWriter, request *http.Request) {
	requestVars := mux.Vars(request)
	log.Printf("Received GET /striker/%s", requestVars["strikerName"])
	strikerData, err := GetStrikerFromDB(app.DB, requestVars["strikerName"])
	if err != nil {
		if err.Error() == "InvalidStrikerName" {
			http.Error(writer, "InvalidStrikerName", http.StatusBadRequest)
		}
		http.Error(writer, fmt.Sprintf("Interal Server Error >> %s", err), http.StatusInternalServerError)
	}

	json.NewEncoder(writer).Encode(strikerData)
}

func (app *App) StatsTotalHandler(writer http.ResponseWriter, request *http.Request) {
	log.Println("Received GET /stats/total")
	json.NewEncoder(writer).Encode(CalcStatTotal())
}

func (app *App) EvenStatsHandler(writer http.ResponseWriter, request *http.Request) {
	log.Println("Received GET /stats/even")

	results, err := GetBalancedStrikers(app.DB)
	if err != nil {
		http.Error(writer, "InternalServerError", http.StatusInternalServerError)
	}

	json.NewEncoder(writer).Encode(results)
}

func (app *App) StatPrioritzeHandler(writer http.ResponseWriter, request *http.Request) {
	log.Println("Receieved POST /stats/prioritize")

	var statsList StatPriorityList
	err := json.NewDecoder(request.Body).Decode(&statsList)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	results, err := GetStrikersForStatPriority(app.DB, statsList.Stats)
	if err != nil {
		if err.Error() == "InvalidStrikerStat" {
			http.Error(writer, "InvalidStrikerStat", http.StatusBadRequest)
		}
		http.Error(writer, fmt.Sprintf("Internal Server Error >> %s", err), http.StatusInternalServerError)
	}

	json.NewEncoder(writer).Encode(results)
}

func (app *App) initRoutes() {
	app.Router.HandleFunc("/health", app.HealthHandler).Methods("GET")
	app.Router.HandleFunc("/calculate", app.CalculateHandler).Methods("GET")
	app.Router.HandleFunc("/striker/{strikerName}", app.StrikerHandler).Methods("GET")
	app.Router.HandleFunc("/stats/total", app.StatsTotalHandler).Methods("GET")
	app.Router.HandleFunc("/stats/even", app.EvenStatsHandler).Methods("GET")
	app.Router.HandleFunc("/stats/prioritize", app.StatPrioritzeHandler).Methods("POST")
}

func (app *App) Init() {
	var err error
	app.DB, err = sql.Open("postgres", "host=localhost port=5432 user=postgres password=docker dbname=strikers_data sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	app.Router = mux.NewRouter()
	app.initRoutes()
}

func (app *App) Run() {
	log.Fatal(http.ListenAndServe(":9000", app.Router))
}

func main() {
	app := App{}
	app.Init()
	app.Run()
}
