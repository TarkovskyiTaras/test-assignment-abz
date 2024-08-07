package datamanager

import "test-assignment-abz/core"

type CurrencyFetcher interface {
	FetchCurrencyRates() (*core.CurrencyDataBankAPI, error)
	FetchHistoricalData() ([]core.CurrencyDataBankAPI, []core.CurrencyDataBankAPI, error)
}
