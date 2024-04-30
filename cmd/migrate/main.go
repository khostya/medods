package main

import (
	"context"
	"fmt"
	"io"
	"log"
	mongo2 "medods/pkg/mongo"
	"os"
	"time"

	migrate "github.com/xakep666/mongo-migrate"
	"medods/config"
	_ "medods/migrations"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatal("Missing options: up, down, new")
	}
	option := os.Args[1]

	cfg := config.MustConfig().MONGO
	mongo, err := mongo2.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer mongo.Disconnect()

	migrate.SetDatabase(mongo.DB)
	migrate.SetMigrationsCollection("migrations")
	migrate.SetLogger(log.New(os.Stdout, "INFO: ", 0))
	switch option {
	case "new":
		if len(os.Args) != 3 {
			log.Fatal("Should be: new description-of-migration")
		}
		fName := fmt.Sprintf("./migrations/%s_%s.go", time.Now().Format("20060102150405"), os.Args[2])
		from, err := os.Open("./migrations/main.go")
		if err != nil {
			log.Fatal("Should be: new description-of-migration")
		}
		defer from.Close()

		to, err := os.OpenFile(fName, os.O_RDWR|os.O_CREATE, 0666)
		if err != nil {
			log.Fatal(err.Error())
		}
		defer to.Close()

		_, err = io.Copy(to, from)
		if err != nil {
			log.Fatal(err.Error())
		}
		log.Printf("New migration created: %s\n", fName)
	case "up":
		err = migrate.Up(context.Background(), migrate.AllAvailable)
	case "down":
		err = migrate.Down(context.Background(), migrate.AllAvailable)
	}

	if err != nil {
		log.Fatal(err.Error())
	}
}
