package helper

import (
	"fmt"
	getip "mta-hosting/getIP"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchActiveMTA(t *testing.T) {
	getip.InitDataStore()

	fmt.Println(FetchActiveMTA())
	results, err := FetchActiveMTA()
	assert.NoError(t, err)
	expectedResults := []string{"mta-prod-1", "mta-prod-3"}
	sort.StringSlice(results).Sort()
	sort.StringSlice(expectedResults).Sort()
	// Compare the actual results with the expected results
	assert.Equal(t, expectedResults, results)
}
