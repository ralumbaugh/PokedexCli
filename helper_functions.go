package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func promptUser(prompt string, scanner *bufio.Scanner) (string, []string) {
	fmt.Print(prompt)
	scanner.Scan()
	userInput := scanner.Text()
	fullCommands := cleanInput(userInput)
	command := fullCommands[0]

	return command, fullCommands
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

// Helper function keeps border the same for formatting throughout app
func addBorder() string {
	return "------------------------------------"
}

func callApi(client http.Client, ioReader io.Reader, callType string, url string) ([]byte, error) {
	// Form api call
	req, err := http.NewRequest(callType, url, ioReader)
	if err != nil {
		return []byte{}, fmt.Errorf("error creating api call: %w", err)
	}

	// Make api call
	resp, err := client.Do(req)

	if err != nil {
		return []byte{}, fmt.Errorf("error performing api call: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return []byte{}, fmt.Errorf("received non-ok status code: %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("error reading body: %w", err)
	}

	return data, nil
}
