package getip

import "errors"

var (
	results []IPConfig
)

func GetIP() ([]IPConfig, error) {
	if results == nil {
		return results, errors.New("no ip found")
	}
	return results, nil
}

func InitDataStore() {
	results = append(results, IPConfig{IP: "127.0.0.1", Hostname: "mta-prod-1", Active: true})
	results = append(results, IPConfig{IP: "127.0.0.2", Hostname: "mta-prod-1", Active: false})
	results = append(results, IPConfig{IP: "127.0.0.3", Hostname: "mta-prod-2", Active: true})
	results = append(results, IPConfig{IP: "127.0.0.4", Hostname: "mta-prod-2", Active: true})
	results = append(results, IPConfig{IP: "127.0.0.5", Hostname: "mta-prod-2", Active: false})
	results = append(results, IPConfig{IP: "127.0.0.6", Hostname: "mta-prod-3", Active: false})
}
