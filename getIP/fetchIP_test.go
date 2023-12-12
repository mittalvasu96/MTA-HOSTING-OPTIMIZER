package getip

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetServerInformation(t *testing.T) {
	InitDataStore()

	results, err := GetIP()
	assert.NoError(t, err)
	expectedResults := []IPConfig{
		{Hostname: "mta-prod-1", IP: "127.0.0.1", Active: true},
		{Hostname: "mta-prod-1", IP: "127.0.0.2", Active: false},
		{Hostname: "mta-prod-2", IP: "127.0.0.3", Active: true},
		{Hostname: "mta-prod-2", IP: "127.0.0.4", Active: true},
		{Hostname: "mta-prod-2", IP: "127.0.0.5", Active: false},
		{Hostname: "mta-prod-3", IP: "127.0.0.6", Active: false},
	}

	sort.Slice(expectedResults, func(i, j int) bool {
		return expectedResults[i].Hostname < expectedResults[j].Hostname || (expectedResults[i].Hostname == expectedResults[j].Hostname && expectedResults[i].IP < expectedResults[j].IP)
	})
	sort.Slice(results, func(i, j int) bool {
		return results[i].Hostname < results[j].Hostname || (results[i].Hostname == results[j].Hostname && results[i].IP < results[j].IP)
	})

	// Compare the actual results with the expected results
	assert.Equal(t, expectedResults, results)
}
