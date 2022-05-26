package main

import (
	"fmt"

	"github.com/by-sabbir/go-rest/internal/comment"
	"github.com/by-sabbir/go-rest/internal/db"
	transportHttp "github.com/by-sabbir/go-rest/internal/transport/http"
)

func Run() error {
	fmt.Println("Starting up the application")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to database")
		return err
	}
	fmt.Println("Successfully Connected to the DB!")

	if err := db.MigrateDB(); err != nil {
		fmt.Println("Failed to migrate database")
		return fmt.Errorf("migrations failed because of: %w", err)
	}

	cmtService := comment.NewService(db)
	httpHandler := transportHttp.NewHandler(cmtService)
	if err := httpHandler.Serve(); err != nil {
		return err
	}
	return nil
}

func main() {
	fmt.Println("Go REST Api Course")

	if err := Run(); err != nil {
		fmt.Println(err)
	}

}
