package utils

import "fmt"

func GetOption() int {
	var option int
	_, err := fmt.Scan(&option)

	if err != nil {
		return 0
	} else {
		return option
	}
}

func GetURLtoMonitoring() string {
	var url string
	_, err := fmt.Scan(&url)

	if err != nil {
		return ""
	} else {
		return url
	}
}
