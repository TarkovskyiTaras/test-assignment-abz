# Detailed Documentation for Currency Provider App

## External APIs and Libraries

### Privat Bank API
The application uses the Privat Bank API for reliable currency data. Example JSON response:
```json
{
  "date": "01.12.2014",
  "bank": "PB",
  "baseCurrency": 980,
  "baseCurrencyLit": "UAH",
  "exchangeRate": [
    {
      "baseCurrency": "UAH",
      "currency": "USD",
      "saleRateNB": 15.056413,
      "purchaseRateNB": 15.056413,
      "saleRate": 15.7,
      "purchaseRate": 15.35
    },
    {
      "baseCurrency": "UAH",
      "currency": "EUR",
      "saleRateNB": 18.79492,
      "purchaseRateNB": 18.79492,
      "saleRate": 20,
      "purchaseRate": 19.2
    }
    // Additional currencies...
  ]
}
```

## Code Structure Overview

### GUI Module
Handles user interactions and displays data. This module includes the interface for selecting currencies via checkboxes and displaying the fetched currency rates.

### Fetcher Module
Fetches data from the API. This module is responsible for making API calls to the Privat Bank API to retrieve the latest currency exchange rates.

### Scheduler Module
Schedules periodic tasks. Using the Cron V3 library, this module sets up tasks to fetch currency data at regular intervals as specified in the configuration file.

### Config Module
Manages configuration settings using Viper. This module reads and writes configuration settings from the `config.yaml` file, allowing users to customize the application's behavior.

### Storage Module
Stores data in a file-based storage system. This module handles the saving of fetched currency data into JSON files, ensuring that the data is persistently stored and can be accessed for historical reference and analysis.