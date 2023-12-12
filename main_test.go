package main

import (
	"fmt"
	getip "mta-hosting/getIP"
	"net/http"
	"testing"
	"time"
)

const url = "http://0.0.0.0:5001/mta-hosting-optimizer/"
const statusUrl = "http://0.0.0.0:5001/healthCheckup/"

func TestHandler(t *testing.T) {

	go startAPIServer()
	time.Sleep(2 * time.Second)
	resp, _ := http.Get(statusUrl)

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
	}
}

func BenchmarkIPSet(b *testing.B) {
	getip.InitDataStore()
	go startAPIServer()
	time.Sleep(2 * time.Second)

	for i := 0; i < b.N; i++ {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Error making GET request: %v\n", err)
			return
		}
		defer resp.Body.Close()
	}
}
