package main

import (
	"time"
)

func service() {
}

func getService() {
}

func main() {
	go service()

	time.Sleep(1 * time.Second)

	getService()
}
