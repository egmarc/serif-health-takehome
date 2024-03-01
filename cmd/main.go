package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type File struct {
	Description string `json:"description"`
	Location    string `json:"location"`
}

func main() {
	resp, err := http.Get("https://antm-pt-prod-dataz-nogbd-nophi-us-east1.s3.amazonaws.com/anthem/2024-03-01_anthem_index.json.gz")
	if err != nil {
		fmt.Printf("HTTP Request failed with error %s\n", err)
		return
	}
	defer resp.Body.Close()

	gz, err := gzip.NewReader(resp.Body)
	if err != nil {
		fmt.Printf("Gzip decompression failed with error %s\n", err)
		return
	}
	defer gz.Close()

	decoder := json.NewDecoder(gz)

	outputFile, err := os.Create("filtered_locations.txt")
	if err != nil {
		fmt.Printf("Failed to create output file: %v", err)
		return
	}
	defer outputFile.Close()

	for {
		token, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Error getting token: %v\n", err)
			return
		}
		if s, ok := token.(string); ok && s == "in_network_files" {
			_, err = decoder.Token() // read beginning of array
			if err != nil {
				fmt.Printf("Error reading array: %v\n", err)
				return
			}
			for decoder.More() {
				var file File
				err = decoder.Decode(&file)
				if err != nil {
					fmt.Printf("Error decoding file: %v\n", err)
					return
				}
				if strings.Contains(file.Description, "PPO") && strings.Contains(file.Description, "New York") {
					fmt.Fprintln(outputFile, file.Location)
				}
			}
			_, err = decoder.Token() // read end of array
			if err != nil {
				fmt.Printf("Error reading array: %v\n", err)
				return
			}
		}
	}
	fmt.Println("Done writing to filtered_locations.txt")
}
