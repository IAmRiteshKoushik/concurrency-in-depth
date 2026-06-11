package main

import (
	"errors"
	"log"
	"log/slog"
	"time"
)

func unreliableService() (any, error) {
	if time.Now().Unix()%2 == 0 {
		return nil, errors.New("service failed")
	}
	return "Success!", nil
}

func main() {
	cb := NewCircuitBreaker(
		2,             // Failure threshold
		2*time.Second, // Recovery time
		2,             // Half-open max requests
		2*time.Second, // Half-open max time
	)

	for range 5 {
		result, err := cb.Call(unreliableService)
		if err != nil {
			slog.Error("Service request failed", "error", err)
		} else {
			slog.Info("Service request succeeded", "result", result)
		}

		time.Sleep(1 * time.Second)
		log.Println("------------------------------")
	}
}
