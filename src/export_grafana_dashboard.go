package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	os.Exit(run())
}

func run() int {
	host := "grafana.example.com"
	auth := os.Getenv("GF_AUTH")

    file,err := os.Create("export_grafana_dashboard.log")
	if err != nil {
		panic("Failed to create log file")
	}


	logger := log.New(file,"",log.LstdFlags | log.Llongfile) 
	logger.Println("1.Println log with log.LstdFlags ...")
	logger.Println("2.Println log with log.LstdFlags ...")

	if len(auth) == 0 {
		//log.Printf("[ERROR] Set GF_AUTH")
		//return 1
		log.Fatal("[ERROR] Set GF_AUTH")
	}

	client := http.DefaultClient
	
	// Get all dashboards URI via search API
	urlStr := fmt.Sprintf("https://%s/%s", host, "api/search")
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		log.Printf("[ERROR] %s", err)
		return 1
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+auth)

	res, err := client.Do(req)
	if err != nil {
		log.Printf("[ERROR] %s", err)
		return 1
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Printf("[ERROR] Invalid status: %s", res.Status)
		return 1
	}

	search := []struct {
		URI string `json:"uri"`
	}{}
	decorder := json.NewDecoder(res.Body)
	if err := decorder.Decode(&search); err != nil {
		log.Printf("[ERROR] %s", err)
		return 1
	}

	log.Printf("[INFO] Number of dashbord: %d", len(search))
	for _, s := range search {
		urlStr := fmt.Sprintf("https://%s/api/dashboards/%s", host, s.URI)
		req, err := http.NewRequest("GET", urlStr, nil)
		if err != nil {
			log.Printf("[ERROR] %s", err)
			return 1
		}

		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", "Bearer "+auth)

		res, err := client.Do(req)
		if err != nil {
			log.Printf("[ERROR] %s", err)
			return 1
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			log.Printf("[ERROR] Invalid status: %s", res.Status)
			return 1
		}

		var db map[string]interface{}
		decorder := json.NewDecoder(res.Body)
		if err := decorder.Decode(&db); err != nil {
			log.Printf("[ERROR] %s", err)
			return 1
		}

		f, err := os.Create(s.URI + ".json")
		if err != nil {
			log.Printf("[ERROR] %s", err)
			return 1
		}

		encoder := json.NewEncoder(f)
		if err := encoder.Encode(db["dashboard"]); err != nil {
			log.Printf("[ERROR] %s", err)
			return 1

		}
	}

	return 0
}
