/*
Package utils provides utility functions for common tasks in the application.

Copyright © 2024 Guilherme Thomazi Bonicontro <thomazi@linux.com>
*/
package utils

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
)

// GetCapacitiesAuth retrieves Capacities authentication credentials from environment variables.
func GetCapacitiesAuth() (string, string) {
	// Retrieve Capacities Space ID
	capacitiesSpaceId := os.Getenv("CAPACITIES_SPACE_ID")
	if capacitiesSpaceId == "" {
		log.Fatalf("CAPACITIES_SPACE_ID environment variable must be defined")
	}

	// Retrieve Capacities API Token
	capacitiesApiToken := os.Getenv("CAPACITIES_API_TOKEN")
	if capacitiesApiToken == "" {
		log.Fatalf("CAPACITIES_API_TOKEN environment variable must be defined")
	}

	return capacitiesSpaceId, capacitiesApiToken
}

// CheckError checks for errors and logs them if present.
func CheckError(err error) {
	if err != nil {
		log.Fatalf(err.Error())
	}
}

// ReadFromStdin reads input from standard input.
func ReadFromStdin() (string, error) {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return "", err
	}

	var stdin []byte
	if stat.Mode()&os.ModeCharDevice == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			stdin = append(stdin, scanner.Bytes()...)
		}
		if err := scanner.Err(); err != nil {
			return "", err
		}
		return string(stdin), nil
	}

	return "", errors.New("unable to read from stdin")
}

// PrettyPrintJSON prints JSON data in a human-readable format.
func PrettyPrintJSON(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}
