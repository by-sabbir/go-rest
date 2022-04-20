package main

import (
	"context"
	"fmt"

	"github.com/by-sabbir/go-rest/internal/db"
)

func Run() error {
	fmt.Println("Starting up the application")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to database")
		return err
	}

	if err := db.Ping(context.Background()); err != nil {
		return err
	}
	fmt.Println("Successfully Connected to the DB!")
	return nil
}

func main() {
	fmt.Println("Go REST Api Course")

	if err := Run(); err != nil {
		fmt.Println(err)
	}

}
