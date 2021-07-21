package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)

const LogsFileName = "logs.txt"

func ShowLogs() {
	file, err := ioutil.ReadFile(LogsFileName)

	if err != nil {
		log.Panic(err)
	}

	fmt.Println(string(file))
}

func writeLog(URL string, isOnline bool) {
	file, err := os.OpenFile(LogsFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Panic(err)
	}

	logs := time.Now().Format("02/01/2006 15:04:05") + URL + " - online: " + strconv.FormatBool(isOnline) + "\n"

	_, err = file.WriteString(logs)

	if err != nil {
		log.Panic(err)
	}
}

