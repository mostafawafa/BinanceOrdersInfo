package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

const apibase = "https://api.binance.com/api/v3/"

type order struct {
	Status   string
	Side     string
	Price    float64 `json:",string"`
	Quantity float64 `json:"executedQty,string"`
}

type orderInfo struct {
	Currency     string
	Quantity     float64
	AveragePrice float64
}

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	http.HandleFunc("/", getPriceInfoController)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getPriceInfoController(w http.ResponseWriter, r *http.Request) {
	currency := r.URL.Query().Get("currency")
	ordersInfo := getOrdersInfo(currency)
	js, _ := json.Marshal(ordersInfo)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
