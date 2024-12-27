package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

const (
	numRequests      = 25_000                  // Total number of requests to send
	concurrencyLevel = 100                     // Number of concurrent requests
	url              = "http://localhost:8080" // Target URL
)

type BenchmarkResult struct {
	latency time.Duration
	err     error
}

func sendRequest(client *http.Client, url string) BenchmarkResult {
	start := time.Now()
	resp, err := client.Get(url)
	latency := time.Since(start)

	if err != nil {
		return BenchmarkResult{latency, err}
	}
	resp.Body.Close()
	return BenchmarkResult{latency, nil}
}

func main() {
	// Setting up an HTTP client with default settings
	client := &http.Client{}

	// Using channels to track results and concurrency with WaitGroup
	results := make(chan BenchmarkResult, numRequests)
	var wg sync.WaitGroup

	// Start time for throughput calculation
	startTime := time.Now()

	// Launch multiple Goroutines to simulate concurrent requests
	for i := 0; i < concurrencyLevel; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < numRequests/concurrencyLevel; j++ {
				result := sendRequest(client, url)
				results <- result
			}
		}()
	}

	// Wait for all Goroutines to finish
	wg.Wait()
	close(results)

	// Calculate total latency and errors
	var totalLatency time.Duration
	var successCount, errorCount int

	for result := range results {
		if result.err != nil {
			errorCount++
			fmt.Println("Error:", result.err)
		} else {
			totalLatency += result.latency
			successCount++
		}
	}

	// End time for throughput calculation
	endTime := time.Now()
	totalTime := endTime.Sub(startTime)

	// Calculate average latency and throughput
	avgLatency := totalLatency / time.Duration(successCount)
	throughput := float64(successCount) / totalTime.Seconds()

	// Print benchmark results
	fmt.Printf("Benchmark Results:\n")
	fmt.Printf("Total Requests: %d\n", numRequests)
	fmt.Printf("Successful Requests: %d\n", successCount)
	fmt.Printf("Failed Requests: %d\n", errorCount)
	fmt.Printf("Total Time Taken: %.2f seconds\n", totalTime.Seconds())
	fmt.Printf("Average Latency: %v\n", avgLatency)
	fmt.Printf("Throughput: %.2f requests/second\n", throughput)
}
