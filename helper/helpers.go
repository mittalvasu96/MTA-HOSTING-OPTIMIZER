package helper

import (
	"fmt"
	"log"
	"mta-hosting/env"
	getip "mta-hosting/getIP"
	"strconv"
)

func FetchActiveMTA() ([]string, error) {

	envThreshold := env.Env["MTA_THRESHOLD"]
	threshold, _ := strconv.Atoi(envThreshold)
	if envThreshold == "" {
		threshold = 1
	}

	fmt.Println(threshold)
	configIps, err := getip.GetIP()
	if err != nil {
		log.Println("Unable to get ip information")
	}

	activeIPs := make(map[string][]string)
	var efficientHost []string
	for _, configIp := range configIps {
		_, ok := activeIPs[configIp.Hostname]
		if !ok {
			activeIPs[configIp.Hostname] = []string{}
		}
		if configIp.Active {
			activeIPs[configIp.Hostname] = append(activeIPs[configIp.Hostname], configIp.IP)
		}

	}

	for hostname, activeIP := range activeIPs {
		if len(activeIP) <= threshold {
			efficientHost = append(efficientHost, hostname)
		}
	}
	return efficientHost, nil
}
