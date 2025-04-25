package main

import (
	"clicknext-backend/internal/config"
	"clicknext-backend/pkg/migration"
	"flag"
	"log"
	"os"
)

func main() {
	var command string
	var step int

	flag.StringVar(&command, "command", "", "Migration command (up/down/step)")
	flag.IntVar(&step, "step", 0, "Number of migration steps")
	flag.Parse()

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	migrator, err := migration.NewMigrator(&cfg.Database)
	if err != nil {
		log.Fatalf("Error creating migrator: %v", err)
	}

	switch command {
	case "up":
		log.Println("Running migrations up...")
		if err := migrator.Up(); err != nil {
			log.Fatalf("Error: %v", err)
		}
	case "down":
		log.Println("Running migrations down...")
		if err := migrator.Down(); err != nil {
			log.Fatalf("Error: %v", err)
		}
	case "step":
		log.Printf("Running %d migration steps...\n", step)
		if err := migrator.Steps(step); err != nil {
			log.Fatalf("Error: %v", err)
		}
	case "drop":
		log.Println("Dropping the database...")
		if err := migrator.DropDatabase(); err != nil {
			log.Fatalf("Error: %v", err)
		}
	default:
		log.Println("Invalid command. Use -command=up|down|step")
		os.Exit(1)
	}

	log.Println("Migration completed successfully")
}
