package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/nativebpm/gotenberg"
)

func main() {
	client, err := gotenberg.NewClient(&http.Client{}, "http://localhost:3000")
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	health, err := client.GetHealth(ctx)
	if err != nil {
		log.Printf("Health check failed: %v", err)
	}

	fmt.Printf("Health Status: %s\n", health.Status)
	for module, status := range health.Details {
		fmt.Printf("  %s: %v\n", module, status)
	}

	version, err := client.GetVersion(ctx)
	if err != nil {
		log.Printf("Version check failed: %v", err)
	}

	fmt.Printf("Version: %s\n", version)

	metrics, err := client.GetMetrics(ctx)
	if err != nil {
		log.Printf("Metrics check failed: %v", err)
	}

	fmt.Printf("Metrics:\n%s\n", metrics)

}
