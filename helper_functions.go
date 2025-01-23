package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func promptUser(prompt string, scanner *bufio.Scanner) (string, string) {
	fmt.Print(prompt)
	scanner.Scan()
	userInput := scanner.Text()
	command := cleanInput(userInput)[0]

	return userInput, command
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

// Helper function keeps border the same for formatting throughout app
func addBorder() string {
	return "------------------------------------"
}

func callApi(client http.Client, ioReader io.Reader, callType string, url string, outputData interface{}) (error){
	// Form api call
	req, err := http.NewRequest(callType, url, ioReader)
	if err != nil {
		return fmt.Errorf("error creating api call: %w", err)
	}

	// Make api call
	resp, err := client.Do(req)

	if err != nil {
		return fmt.Errorf("error performing api call: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return fmt.Errorf("received non-ok status code: %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading body: %w", err)
	}

	err = json.Unmarshal(data, outputData)
	if err != nil {
		return fmt.Errorf("error unmarhsaling output data: %w", err)
	}

	return nil
}