package main

import (
	"log"
	"test-assignment-abz/config"
	"test-assignment-abz/datamanager"
	"test-assignment-abz/guimanager"

	"github.com/robfig/cron/v3"
)

func main() {
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
