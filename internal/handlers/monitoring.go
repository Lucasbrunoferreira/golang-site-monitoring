package handlers

import (
	"../utils"
	"fmt"
	"net/http"
	"time"
)

const MonitoingCount = 3
const MonitoringDelay = 10 * time.Second

func StartMonitoring() {
	fmt.Println("Enter the URL to monitoring: ")
	siteToMonitoring := utils.GetURLtoMonitoring()
	TestURL(siteToMonitoring)
}

func TestURL(URL string) {
	fmt.Println("Starting monitoring...")

	for i := 0; i < MonitoingCount; i++ {
		response, err := http.Get(URL)

		if err != nil {
			fmt.Println("Error on access the site")
		}

		var isOnline bool = false

		if response != nil {
			isOnline = response.StatusCode == 200

			if isOnline {
				fmt.Println("Site is living!", response.Status)
			} else {
				fmt.Println("The site is not available!")
			}
		}


		writeLog(URL, isOnline)

		time.Sleep(MonitoringDelay)
	}
}
