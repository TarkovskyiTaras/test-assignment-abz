package main

import (
	"fmt"
	"golang.org/x/sys/windows/registry"
	"log"
	"os"
	"test-assignment-abz/config"
	"test-assignment-abz/datamanager"
	"test-assignment-abz/guimanager"

	"github.com/robfig/cron/v3"
)

func main() {
	err := handleArgs()

	if err != nil {
		log.Fatal(err)
	}
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}

	apiClient := datamanager.NewAPIClient()
	fileStorage := datamanager.NewFileStorage(apiClient, cfg)

	isComplete, err := fileStorage.VerifyDataCompletion()
	if err != nil {
		log.Fatal(err)
	}
	if !isComplete {
		log.Println("Data is incomplete, full fetching is started")
		err = fileStorage.FetchAndSaveHistoricalData()
		if err != nil {
			log.Fatal(err)
		}
	}

	//Reorganize all data every first day of a month at 00:00
	c := cron.New()
	_, err = c.AddFunc("0 0 1 * *", func() {
		log.Println("fetching first day of a month's rates")
		err2 := fileStorage.FetchAndSaveFirstDayOfMonthRates()
		if err2 != nil {
			log.Fatal(err2)
		}
	})

	//Update today's rates each x interval
	_, err = c.AddFunc(cfg.FetchInterval, func() {
		log.Println("updating today's rates")
		err2 := fileStorage.UpdateTodaysData()
		if err2 != nil {
			log.Fatal(err2)
		}
	})
	if err != nil {
		log.Fatal(err)
	}

	c.Start()
	defer c.Stop()

	guiManager := guimanager.NewGUIManager(fileStorage, cfg)
	guiManager.Run()
}

func handleArgs() error {
	appName := "currency_app"
	exePath, err := os.Executable()
	if err != nil {
		return err
	}
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "-install":
			err := setAutorun(appName, exePath)
			if err != nil {
				return err
			}
			fmt.Println("Application set to run on startup.")
		case "-uninstall":
			err := removeAutorun(appName)
			if err != nil {
				return err
			}
			fmt.Println("Application removed from startup.")
		default:
			return fmt.Errorf("unknown option: %s", os.Args[1])
		}
	}

	return nil
}

func setAutorun(keyName, exePath string) error {
	key, exists, err := registry.CreateKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Run`, registry.SET_VALUE)
	if err != nil {
		return fmt.Errorf("could not access registry: %v", err)
	}
	defer key.Close()

	if exists {
		log.Println("Autorun key already exists")
	}

	err = key.SetStringValue(keyName, exePath)
	if err != nil {
		return fmt.Errorf("could not set autorun: %v", err)
	}

	return nil
}

func removeAutorun(keyName string) error {
	key, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Run`, registry.SET_VALUE)
	if err != nil {
		return fmt.Errorf("could not access registry: %v", err)
	}
	defer key.Close()

	err = key.DeleteValue(keyName)
	if err != nil {
		return fmt.Errorf("could not remove autorun: %v", err)
	}

	return nil
}
