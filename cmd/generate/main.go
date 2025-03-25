package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/4chain-ag/universal-test-vectors/vectors"
)

func main() {
	for k, v := range vectors.All {
		generateJSON(k, v)
	}
}

func generateJSON(key string, s any) {
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling JSON: %v\n", err)
		return
	}

	fileName := fmt.Sprintf("%s.json", key)
	filePath := filepath.Join("generated", fileName)

	err = os.MkdirAll("generated", 0750)
	if err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		return
	}

	err = os.WriteFile(filePath, data, 0600)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		return
	}

	fmt.Printf("Generated JSON file: %s\n", filePath)
}
