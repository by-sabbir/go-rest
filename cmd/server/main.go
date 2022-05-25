package main

import (
	"context"
	"fmt"

	"github.com/by-sabbir/go-rest/internal/comment"
	"github.com/by-sabbir/go-rest/internal/db"
)

func Run() error {
	fmt.Println("Starting up the application")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to database")
		return err
	}

	if err := db.MigrateDB(); err != nil {
		fmt.Println("Failed to migrate database")
		return fmt.Errorf("migrations failed because of: %w", err)
	}
	fmt.Println("Successfully Connected to the DB!")

	cmtService := comment.NewService(db)
	cmt, _ := cmtService.GetComment(context.Background(), "DFFC3516-6C4A-4B81-81B7-8E6298A410A4")
	fmt.Println(cmt)
	return nil
}

func main() {
	fmt.Println("Go REST Api Course")

	if err := Run(); err != nil {
		fmt.Println(err)
	}

}
