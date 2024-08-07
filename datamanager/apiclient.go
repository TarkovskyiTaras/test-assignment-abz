package datamanager

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"test-assignment-abz/core"
)

const baseURL = "https://api.privatbank.ua/p24api/exchange_rates?date="

type APIClient struct {
	httpClient *http.Client
}

func NewAPIClient() *APIClient {
	return &APIClient{
		httpClient: &http.Client{
			Timeout: time.Second * 30,
		},
	}
}

func (a *APIClient) FetchCurrencyRates() (*core.CurrencyDataBankAPI, error) {
	date := time.Now()
	formattedDate := date.Format("02.01.2006")

	url := baseURL + formattedDate
	response, err := a.httpClient.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}

	var payload core.CurrencyDataBankAPI
	err = json.Unmarshal(body, &payload)
	if err != nil {
		log.Println(err)
	}

	return &payload, nil
}

func (a *APIClient) FetchHistoricalData() ([]core.CurrencyDataBankAPI, []core.CurrencyDataBankAPI, error) {
	lastMonth := LastMonthDates()
	thisMonth := ThisMonthDates()

	lastMonthCurrencies := make([]core.CurrencyDataBankAPI, 0)
	for _, date := range lastMonth {
		url := baseURL + date
		response, err := a.httpClient.Get(url)
		if err != nil {
			log.Println(err)
		}

		body, err := io.ReadAll(response.Body)
		if err != nil {
			log.Println(err)
		}
		response.Body.Close()

		var payload core.CurrencyDataBankAPI
		err = json.Unmarshal(body, &payload)
		if err != nil {
			log.Println(err)
		}

		lastMonthCurrencies = append(lastMonthCurrencies, payload)
		fmt.Println(payload)
		time.Sleep(10 * time.Second)
	}

	thisMonthCurrencies := make([]core.CurrencyDataBankAPI, 0)
	for _, date := range thisMonth {
		url := baseURL + date
		response, err := a.httpClient.Get(url)
		if err != nil {
			log.Println(err)
		}

		body, err := io.ReadAll(response.Body)
		if err != nil {
			log.Println(err)
		}
		response.Body.Close()

		var payload core.CurrencyDataBankAPI
		err = json.Unmarshal(body, &payload)
		if err != nil {
			log.Println(err)
		}

		thisMonthCurrencies = append(thisMonthCurrencies, payload)
		fmt.Println(payload)

		time.Sleep(10 * time.Second)
	}

	return lastMonthCurrencies, thisMonthCurrencies, nil
}

func LastMonthDates() []string {
	n := time.Now()
	s := n.AddDate(0, -1, 0)
	start := time.Date(s.Year(), s.Month(), 1, 0, 0, 0, 0, time.Local)
	e := start.AddDate(0, 1, -1)
	end := time.Date(e.Year(), e.Month(), e.Day(), 0, 0, 0, 0, time.Local)

	var lastMonth []string
	for d := start; d.After(end) == false; d = d.AddDate(0, 0, 1) {
		date := d.Format("02.01.2006")
		lastMonth = append(lastMonth, date)
	}

	return lastMonth
}

func ThisMonthDates() []string {
	n := time.Now()
	start := time.Date(n.Year(), n.Month(), 1, 0, 0, 0, 0, time.Local)
	end := n

	var thisMonth []string
	for d := start; d.Before(end) || d.Equal(end); d = d.AddDate(0, 0, 1) {
		date := d.Format("02.01.2006")
		thisMonth = append(thisMonth, date)
	}
	return thisMonth
}
