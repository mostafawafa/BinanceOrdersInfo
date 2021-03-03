package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func sendRequest(route string, currencyPair string) []order {

	timeStamp := strconv.FormatInt(time.Now().UnixNano()/1000000, 10)
	query := "timestamp=" + timeStamp + "&symbol=" + currencyPair
	signtaure := getSignature(query)
	client := &http.Client{}
	url := apibase + route + "?" + query + "&signature=" + signtaure
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("X-MBX-APIKEY", os.Getenv("API_KEY"))

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return convertBodyToOrders(body)
}

func convertBodyToOrders(body []byte) []order {
	var orders []order

	error := json.Unmarshal(body, &orders)
	if error != nil {
		log.Fatal(error)
	}
	return orders
}
