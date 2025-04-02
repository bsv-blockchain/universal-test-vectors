package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/4chain-ag/universal-test-vectors/vectors"
)

//go:generate go run main.go --destination=../../generated
//go:generate npm install && npx tsx ../../brc100frames/index.ts ../../generated

func main() {
	destination := flag.String("destination", "generated", "Destination directory for generated files")
	flag.Parse()

	for k, v := range vectors.All {
		generateJSON(*destination, k, v)
	}
}

func generateJSON(destination string, key string, s any) {
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling JSON: %v\n", err)
		return
	}

	fileName := fmt.Sprintf("%s.json", key)
	filePath := filepath.Join(destination, fileName)

	err = os.MkdirAll(destination, 0750)
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
