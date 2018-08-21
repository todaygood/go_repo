package main

import (
    "fmt"
    "github.com/opsnull/kping"
    logger "log"
    "time"
    )

func main() {
	// Create a new Pinger
	pinger, err := kping.NewPinger("192.168.60.236", 100, 10, 1*time.Minute, 100*time.Millisecond)
	if err != nil {
		logger.Fatalln(err)
	}

	// Add IP addresses to Pinger
	if err := pinger.AddIPs([]string{"114.114.114.114", "8.8.8.8"}); err != nil {
		logger.Fatalln(err)
	}

	// Run !
	statistics, err := pinger.Run()
	if err != nil {
		logger.Fatalln(err)
	}

	// Print result
	for ip, statistic := range statistics {
		fmt.Printf("%s: %v\n", ip, statistic)
	}
}
