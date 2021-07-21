package handlers

import (
	"../utils"
	"fmt"
)

func ShowMenu () {
	for {
		fmt.Println("Select one option:")
		fmt.Println("(1) - Start monitoring")
		fmt.Println("(2) - Show logs")
		fmt.Println("(0) - Exit")

		handleOptions()
	}
}

func handleOptions() {
	option := utils.GetOption()

	switch option {
	case 1:
		StartMonitoring()
	case 2:
		ShowLogs()
	case 0:
		utils.HandleExit()
	default:
		utils.ShowErrorMessage()
	}
}
