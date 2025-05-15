package main

import (
	"log"
	"routecore/configs"
	"routecore/internal/db"
	"routecore/internal/seed"
)

func main() {
	conf := configs.LoadConfig()
	database := db.NewDb(&conf.Db)


	if err := seed.Run(database.DB); err != nil {
		log.Fatalf("Seed error: %v", err)
	}

	log.Println("Seeding completed successfully!")
}
