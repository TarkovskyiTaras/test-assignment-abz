# Currency Provider App

Currency Provider App is a dynamic, user-friendly application designed to provide real-time currency exchange rate information for a diverse set of global currencies. This tool is essential for anyone needing up-to-date currency data, from financial analysts to travelers wanting to get the best exchange rate for their next trip.
## Installation

### Prerequisites

1. **Go Installation**: Ensure that you have Go installed on your system. You can download it from [https://golang.org/dl/](https://golang.org/dl/).
2. **C Compiler**: For Windows users, install one of the supported C compilers needed for the Fyne library:
    - **TDM-GCC**: Download and install from [tdm-gcc.tdragon.net](http://tdm-gcc.tdragon.net/).

### Getting Started

Follow these steps to install the application:

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/TarkovskyiTaras/test-assignment-abz
   cd test-assignment-abz
2. **Download Dependencies**:
   ```bash
   go mod download
3. **Build the Application**:
   ```bash
   go build -o currency-provider-app.exe
4. **Build the Application**:
   ```bash
   ./currency-provider-app.exe
### Uninstallation
   ```bash
   rm -rf test-assignment-abz
   ```
## Configuration
The application can be configured using a YAML file located in the `./config` directory. Below is the available configuration option and how to customize it.

## Configuration

Configure the application using the `config.yaml` file located in the `./config` directory. This configuration file allows you to set options related to how frequently currency rates are fetched and which currencies are displayed.

### Configuring with YAML

Modify the `config.yaml` file located at `./config/config.yaml` according to your needs:
```yaml
# How often to fetch today's currency rates
fetch_interval: "@every 10s"
```

Include ISO Currency Codes (ISO 4217) for selection via checkboxes in the GUI interface under the 'currencies' configuration:
```yaml
currencies:
```
  - USD
  - EUR
  - UAH
  - GBP
  - ...

## Dependencies

This application relies on several external libraries that need to be included for proper functionality. Below are the dependencies required:

- **Fyne**: A modern UI toolkit for building desktop applications in Go. This is used to create and manage the graphical user interface.
   - Version: `fyne.io/fyne/v2 v2.5.0`
   - [Fyne Documentation](https://fyne.io/)


- **Cron**: A robust scheduler in Go, allowing for the timing of execution of tasks, such as fetching the latest currency rates at predefined intervals.
   - Version: `github.com/robfig/cron/v3 v3.0.1`
   - [Cron V3 Documentation](https://github.com/robfig/cron)


- **Viper**: A complete configuration solution for Go applications including 12-factor apps. It is designed to work within an application to handle configuration needs and formats.
   - Version: `github.com/spf13/viper v1.19.0`
   - [Viper GitHub Repository](https://github.com/spf13/viper)

## Troubleshooting Tips and Common Issues

Here are some troubleshooting tips and common issues that users might encounter when using the application, along with guidance on how to address them.

### Initial Data Fetch Duration

**Issue**: The first time you run the application, it initiates a full data fetch for the past two months from the PrivatBank API. This process may take several minutes.

**Cause**: The application is designed to retrieve historical currency data by making multiple API calls to PrivatBank, each fetching data for a single day. To avoid overwhelming the API and to comply with rate limits, there is a deliberate delay of 10 seconds between each call.

**Resolution**: Please be patient during this initial setup phase. Allow the application to complete this initial data fetch uninterrupted. Subsequent data updates will be faster, involving only daily updates rather than a full historical fetch.

### Delayed Response or Timeouts

**Issue**: Users may experience delays or timeouts during the initial or subsequent data fetches.

**Cause**: Delays or timeouts can occur due to network issues, server load at PrivatBank, or your local internet connection speed.

**Resolution**: If you experience a timeout, wait a few moments and retry. Ensure your internet connection is stable. If the problem persists, consider increasing the delay between API calls in the configuration settings if the API's rate limit policies allow.

### Configuring API Call Delays

**Issue**: Need to adjust the delay between API calls to handle different rate limit policies or to improve response times.

**Resolution**: You can adjust the delay between API calls by changing the `fetch_interval` in the `config.yaml` file. Be cautious with reducing the delay as it might lead to hitting the API rate limits more quickly.

### Additional Support

If you encounter issues that are not resolved by the above tips, please check the API documentation or contact support for further assistance.