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
	cmtService.PostComment(
		context.Background(),
		comment.Comment{
			ID:     "",
			Body:   "Comment directly from Jane",
			Slug:   "comment-directly-from-jane",
			Author: "Jane Doe",
		},
	)
	cmt, _ := cmtService.GetComment(context.Background(), "01623d7d-cf41-4a37-adc7-730f0fd07b53")
	fmt.Println(cmt)

	if err := cmtService.DeleteComment(context.Background(), "eec973dd-0659-4681-95fa-4d79f03f6397"); err != nil {
		return fmt.Errorf("error deleting comment: %+v", err)
	}

	deletedCmt, _ := cmtService.GetComment(context.Background(), "eec973dd-0659-4681-95fa-4d79f03f6397")
	fmt.Println(deletedCmt)

	cmtService.UpdateComment(
		context.Background(),
		"1c19656e-b3b2-4623-9af3-19b9ffa48d3c",
		comment.Comment{
			ID:     "",
			Body:   "this is updated comment",
			Slug:   "this-is-updated-comment",
			Author: "updated-author",
		})
	updatedCmt, _ := cmtService.GetComment(context.Background(), "1c19656e-b3b2-4623-9af3-19b9ffa48d3c")
	fmt.Println(updatedCmt)
	return nil
}

func main() {
	fmt.Println("Go REST Api Course")

	if err := Run(); err != nil {
		fmt.Println(err)
	}

}
