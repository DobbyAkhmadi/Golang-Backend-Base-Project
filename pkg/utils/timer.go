package utils

import (
	"time"
)

func GetCurrentLatency() string {
	// Measure the time taken for the request
	startTime := time.Now()

	// Simulate some processing time
	time.Sleep(time.Second)

	// Calculate the elapsed time
	elapsedTime := time.Since(startTime).String()

	return elapsedTime
}
