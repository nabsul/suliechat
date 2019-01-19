package main

import (
	"log"
	"time"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func convertTime(timestamp string) string {
	utc, err := time.Parse(time.RFC3339, timestamp)
	checkError(err)
	return utc.Local().Format("02 Jan 2006 15:04:05 MST")
}
