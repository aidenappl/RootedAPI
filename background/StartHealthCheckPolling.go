package background

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/aidenappl/rootedapi/env"
)

func StartHealthCheckPolling(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			log.Println("🔍 Pinging health check...")

			// Create HTTP client with timeout
			client := &http.Client{
				Timeout: 10 * time.Second,
			}

			// Make GET request to health check endpoint
			resp, err := client.Get(env.HealthCheckURL)
			if err != nil {
				log.Printf("❌ Health check failed: %v", err)
				continue
			}

			resp.Body.Close()

			if resp.StatusCode == http.StatusOK {
				log.Println("✅ Health check successful")
			} else {
				log.Printf("⚠️ Health check returned status: %d", resp.StatusCode)
			}

		case <-ctx.Done():
			log.Println("🚦 Stopping health check polling...")
			return
		}
	}
}
