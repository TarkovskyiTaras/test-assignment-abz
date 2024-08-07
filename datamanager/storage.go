package datamanager

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"
	"time"

	"test-assignment-abz/config"
	"test-assignment-abz/core"
)

type FileStorage struct {
	aPIClient CurrencyFetcher
	cfg       *config.Config
}

func NewFileStorage(apiClient *APIClient, cfg *config.Config) *FileStorage {
	return &FileStorage{
		aPIClient: apiClient,
		cfg:       cfg,
	}
}

func (f *FileStorage) UpdateTodaysData() error {
	data, err := f.aPIClient.FetchCurrencyRates()
	if err != nil {
		return err
	}

	filePath := "./files/currency_thismonth.json"
	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("Error opening file: %s\n", err)
		return err
	}
	defer file.Close()

	data2Json, err := io.ReadAll(file)
	if err != nil {
		log.Printf("Error reading file: %s\n", err)
		return err
	}

	var data2 []core.CurrencyDataBankAPI
	err = json.Unmarshal(data2Json, &data2)
	if err != nil {
		return err
	}

	for _, d := range data2 {
		if d.Date == data.Date {
			d.ExchangeRate = data.ExchangeRate
		}
	}

	dataJson, err := json.MarshalIndent(data2, "", "	")
	if err != nil {
		return err
	}

	filePath = "./files/currency_thismonth.json"
	file, err = os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(dataJson)
	if err != nil {
		return err
	}

	return nil
}

func (f *FileStorage) FetchAndSaveFirstDayOfMonthRates() error {
	err := os.Remove("./files/currency_lastmonth.json")
	if err != nil {
		return err
	}

	err = os.Rename("./files/currency_thismonth.json", "./files/currency_lastmonth.json")
	if err != nil {
		return err
	}

	_, err = os.Create("./files/currency_thismonth.json")
	if err != nil {
		return err
	}

	data, err := f.aPIClient.FetchCurrencyRates()
	if err != nil {
		return err
	}

	dataJson, err := json.MarshalIndent(data, "", "	")
	if err != nil {
		return err
	}

	filePath := "./files/currency_thismonth.json"
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(dataJson)
	if err != nil {
		return err
	}

	return nil
}

func (f *FileStorage) GetTodaysData(selectedCurrencies map[string]bool) (json.RawMessage, error) {
	filePath := "./files/currency_thismonth.json"
	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("Error opening file: %s\n", err)
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Printf("Error reading file: %s\n", err)
		return nil, err
	}

	var payloadThisMonth []core.CurrencyDataBankAPI
	err = json.Unmarshal(data, &payloadThisMonth)
	if err != nil {
		return nil, err
	}

	t := time.Now()
	today := t.Format("02.01.2006")

	var todaysRates core.CurrencyDataBankAPI
	for _, p := range payloadThisMonth {
		if p.Date == today {
			todaysRates = p
		}
	}

	todaysRatesFiltered := filterCurrencies([]core.CurrencyDataBankAPI{todaysRates}, selectedCurrencies)
	todaysRatesJson, err := json.MarshalIndent(todaysRatesFiltered, "", "	")
	if err != nil {
		return nil, err
	}

	return todaysRatesJson, nil
}

func (f *FileStorage) FetchAndSaveHistoricalData() error {
	dataLastMonth, dataThisMonth, err := f.aPIClient.FetchHistoricalData()
	if err != nil {
		return err
	}

	jsonLastMonth, err := json.MarshalIndent(dataLastMonth, "", "	")
	if err != nil {
		return err
	}

	filePath := "./files/currency_lastmonth.json"

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(jsonLastMonth)
	if err != nil {
		return err
	}

	jsonThisMonth, err := json.MarshalIndent(dataThisMonth, "", "	")
	if err != nil {
		return err
	}

	filePath = "./files/currency_thismonth.json"

	file, err = os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(jsonThisMonth)
	if err != nil {
		return err
	}

	return nil
}

func (f *FileStorage) VerifyDataCompletion() (bool, error) {
	if _, err := os.Stat("./files/currency_lastmonth.json"); errors.Is(err, os.ErrNotExist) {
		return false, nil
	}

	if _, err := os.Stat("./files/currency_thismonth.json"); errors.Is(err, os.ErrNotExist) {
		return false, nil
	}

	filePath := "./files/currency_thismonth.json"
	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("Error opening file: %s\n", err)
		return false, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Printf("Error reading file: %s\n", err)
		return false, err
	}

	if len(data) < 1 {
		return false, nil
	}

	var payloadThisMonth []core.CurrencyDataBankAPI
	err = json.Unmarshal(data, &payloadThisMonth)
	if err != nil {
		return false, err
	}

	filePath = "./files/currency_lastmonth.json"
	file, err = os.Open(filePath)
	if err != nil {
		log.Printf("Error opening file: %s\n", err)
		return false, err
	}
	defer file.Close()

	data, err = io.ReadAll(file)
	if err != nil {
		log.Printf("Error reading file: %s\n", err)
		return false, err
	}

	if len(data) < 1 {
		return false, nil
	}

	var payloadLastMonth []core.CurrencyDataBankAPI
	err = json.Unmarshal(data, &payloadLastMonth)
	if err != nil {
		return false, err
	}

	datesThisMonth := make([]string, 0)
	for _, p := range payloadThisMonth {
		datesThisMonth = append(datesThisMonth, p.Date)
	}
	datesLastMonth := make([]string, 0)
	for _, p := range payloadLastMonth {
		datesLastMonth = append(datesLastMonth, p.Date)
	}

	thisMonth := ThisMonthDates()
	lastMonth := LastMonthDates()

	for _, date := range thisMonth {
		if !contains(datesThisMonth, date) {
			return false, nil
		}
	}

	for _, date := range lastMonth {
		if !contains(datesLastMonth, date) {
			return false, nil
		}
	}

	return true, nil
}

func (f *FileStorage) GetThisMonthData(selectedCurrencies map[string]bool) (json.RawMessage, error) {
	filePath := "./files/currency_thismonth.json"
	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("Error opening file: %s\n", err)
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Printf("Error reading file: %s\n", err)
		return nil, err
	}

	var payload []core.CurrencyDataBankAPI
	err = json.Unmarshal(data, &payload)
	if err != nil {
		return nil, err
	}

	filteredCurrencies := filterCurrencies(payload, selectedCurrencies)
	out, err := json.MarshalIndent(filteredCurrencies, "", "	")
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (f *FileStorage) GetLastMonthData(selectedCurrencies map[string]bool) (json.RawMessage, error) {
	filePath := "./files/currency_lastmonth.json"
	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("Error opening file: %s\n", err)
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Printf("Error reading file: %s\n", err)
		return nil, err
	}

	var payload []core.CurrencyDataBankAPI
	err = json.Unmarshal(data, &payload)
	if err != nil {
		return nil, err
	}

	filteredCurrencies := filterCurrencies(payload, selectedCurrencies)
	out, err := json.MarshalIndent(filteredCurrencies, "", "	")
	if err != nil {
		return nil, err
	}

	return out, nil
}

func filterCurrencies(data []core.CurrencyDataBankAPI, selectedCurrencies map[string]bool) []core.CurrencyDataBankAPI {
	filteredData := make([]core.CurrencyDataBankAPI, 0)

	for _, item := range data {
		filteredExchangeRates := make([]core.ExchangeRateBankAPI, 0)
		for _, rate := range item.ExchangeRate {
			if checked := selectedCurrencies[rate.Currency]; checked {
				filteredExchangeRates = append(filteredExchangeRates, rate)
			}
		}
		if len(filteredExchangeRates) > 0 {
			newItem := core.CurrencyDataBankAPI{
				Date:         item.Date,
				Bank:         item.Bank,
				BaseCurrency: item.BaseCurrency,
				ExchangeRate: filteredExchangeRates,
			}
			filteredData = append(filteredData, newItem)
		}
	}

	return filteredData
}

func contains(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}
