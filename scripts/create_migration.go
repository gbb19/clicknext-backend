package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	var name string
	flag.StringVar(&name, "name", "", "Migration name")
	flag.Parse()

	if name == "" {
		fmt.Println("Error: Migration name is required")
		os.Exit(1)
	}

	timestamp := time.Now().Unix()
	filename := fmt.Sprintf("%d_%s", timestamp, name)

	upFilePath := filepath.Join("migrations", filename+".up.sql")
	downFilePath := filepath.Join("migrations", filename+".down.sql")

	upFile, err := os.Create(upFilePath)
	if err != nil {
		fmt.Printf("Error creating up migration file: %v\n", err)
		os.Exit(1)
	}
	defer upFile.Close()

	downFile, err := os.Create(downFilePath)
	if err != nil {
		fmt.Printf("Error creating down migration file: %v\n", err)
		os.Exit(1)
	}
	defer downFile.Close()

	fmt.Printf("Created migration files:\n%s\n%s\n", upFilePath, downFilePath)
}
