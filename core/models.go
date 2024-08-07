package core

type CurrencyDataBankAPI struct {
	Date         string                `json:"date"`
	Bank         string                `json:"bank"`
	BaseCurrency string                `json:"baseCurrencyLit"`
	ExchangeRate []ExchangeRateBankAPI `json:"exchangeRate"`
}

type ExchangeRateBankAPI struct {
	Currency       string  `json:"currency"`
	SaleRate       float64 `json:"saleRate,omitempty"`
	PurchaseRate   float64 `json:"purchaseRate,omitempty"`
	SaleRateNB     float64 `json:"saleRateNB,omitempty"`
	PurchaseRateNB float64 `json:"purchaseRateNB,omitempty"`
}
