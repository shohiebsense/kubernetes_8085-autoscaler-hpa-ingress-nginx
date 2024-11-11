package main

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// TrafficMetrics stores network traffic data
type TrafficMetrics struct {
	RequestCount    int           `json:"request_count"`
	TotalResponseTime time.Duration `json:"total_response_time"`
	AverageResponseTime float64    `json:"average_response_time"`
}

var metrics TrafficMetrics
var mutex sync.Mutex // Mutex to prevent data race

// MetricsMiddleware tracks the request count and response time
func MetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now() // Track start time

		// Process the request
		c.Next()

		// Calculate request duration
		duration := time.Since(startTime)

		// Update metrics with a mutex for thread safety
		mutex.Lock()
		metrics.RequestCount++
		metrics.TotalResponseTime += duration
		metrics.AverageResponseTime = metrics.TotalResponseTime.Seconds() / float64(metrics.RequestCount)
		mutex.Unlock()
	}
}

func main() {
	// Initialize Gin router
	r := gin.Default()

	// Add the metrics middleware
	r.Use(MetricsMiddleware())

	// Define the /network-traffic endpoint
	r.GET("/network-traffic", func(c *gin.Context) {
		// Lock mutex to safely read metrics
		mutex.Lock()
		defer mutex.Unlock()

		// Return the metrics as JSON
		c.JSON(http.StatusOK, metrics)
	})

	// Define a sample endpoint to simulate traffic
	r.GET("/sample-endpoint", func(c *gin.Context) {
		time.Sleep(100 * time.Millisecond) // Simulate some processing time
		c.JSON(http.StatusOK, gin.H{"message": "This is a sample endpoint"})
	})

	// Start the server
	r.Run(":8085")
}
