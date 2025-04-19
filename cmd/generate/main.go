package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/bsv-blockchain/universal-test-vectors/vectors"
)

//go:generate go run main.go --destination=../../generated
//go:generate npm install
//go:generate npx --yes tsx ../../brc100frames/index.ts ../../generated/brc100

func main() {
	destination := flag.String("destination", "generated", "Destination directory for generated files")
	flag.Parse()

	for k, v := range vectors.All {
		generateJSON(*destination, k, v)
	}
}

func generateJSON(destination string, key string, vec vectors.TestVector) {
	data, err := json.MarshalIndent(vec.Vector, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling JSON: %v\n", err)
		return
	}

	fileName := fmt.Sprintf("%s.json", key)
	dir := filepath.Join(destination, vec.Category)
	filePath := filepath.Join(dir, fileName)

	err = os.MkdirAll(dir, 0750)
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
