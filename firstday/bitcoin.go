package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

type BitcoinData struct {
	Time struct {
		Updated    string    `json:"updated"`
		UpdatedISO time.Time `json:"updatedISO"`
		Updateduk  string    `json:"updateduk"`
	} `json:"time"`
	Disclaimer string `json:"disclaimer"`
	ChartName  string `json:"chartName"`
	Bpi        struct {
		USD struct {
			Code        string  `json:"code"`
			Symbol      string  `json:"symbol"`
			Rate        string  `json:"rate"`
			Description string  `json:"description"`
			RateFloat   float64 `json:"rate_float"`
		} `json:"USD"`
		GBP struct {
			Code        string  `json:"code"`
			Symbol      string  `json:"symbol"`
			Rate        string  `json:"rate"`
			Description string  `json:"description"`
			RateFloat   float64 `json:"rate_float"`
		} `json:"GBP"`
		EUR struct {
			Code        string  `json:"code"`
			Symbol      string  `json:"symbol"`
			Rate        string  `json:"rate"`
			Description string  `json:"description"`
			RateFloat   float64 `json:"rate_float"`
		} `json:"EUR"`
	} `json:"bpi"`
}

func main() {
	var url string = "https://api.coindesk.com/v1/bpi/currentprice.json"
	//var currency string = "EUR"
	//var price float64 = 0.0
	var bitcoin BitcoinData
	client := &http.Client{}
	res, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(body, &bitcoin)
	log.Println("Bitcoin Daten vom: ", bitcoin.Time.Updated)
	log.Println(bitcoin.Bpi.USD.Rate)
	log.Println(bitcoin.Bpi.GBP.Rate)
	log.Println(bitcoin.Bpi.EUR.Rate)
	log.Println(bitcoin.Bpi.USD.RateFloat)
	log.Println(bitcoin.Bpi.GBP.RateFloat)
	log.Println(bitcoin.Bpi.EUR.RateFloat)

}
