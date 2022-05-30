package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/by-sabbir/go-rest/internal/comment"
	"github.com/google/uuid"
)

type CommentRow struct {
	ID     string
	Slug   sql.NullString
	Body   sql.NullString
	Author sql.NullString
}

func convertComment(c CommentRow) comment.Comment {
	return comment.Comment{
		ID:     c.ID,
		Body:   c.Body.String,
		Slug:   c.Slug.String,
		Author: c.Author.String,
	}
}

func (d *DataBase) GetComment(ctx context.Context, uuid string) (comment.Comment, error) {
	var cmtRow CommentRow
	row := d.Client.QueryRowContext(
		ctx,
		`select id, body, slug, author
		from comments
		where id::text=$1`,
		uuid,
	)
	err := row.Scan(&cmtRow.ID, &cmtRow.Body, &cmtRow.Slug, &cmtRow.Author)
	if err != nil {
		return comment.Comment{}, fmt.Errorf("error fetching comments from uuid: %+v", err)
	}
	return convertComment(cmtRow), nil
}

func (d *DataBase) PostComment(ctx context.Context, cmt comment.Comment) (comment.Comment, error) {
	cmt.ID = uuid.NewString()
	postRow := CommentRow{
		ID:     cmt.ID,
		Slug:   sql.NullString{String: cmt.Slug, Valid: true},
		Body:   sql.NullString{String: cmt.Body, Valid: true},
		Author: sql.NullString{String: cmt.Author, Valid: true},
	}
	row, err := d.Client.NamedQueryContext(
		ctx,
		`insert into comments
		(id, body, slug, author)
		values
		(:id, :body, :slug, :author)`,
		postRow,
	)
	if err != nil {
		return comment.Comment{}, fmt.Errorf("error posting comment: %+v", err)
	}
	if err := row.Close(); err != nil {
		return comment.Comment{}, fmt.Errorf("could not close the row, %+v", err)
	}

	return convertComment(postRow), nil
}

func (d *DataBase) DeleteComment(ctx context.Context, uuid string) error {
	row, err := d.Client.ExecContext(
		ctx,
		`delete from comments
		where id::text=$1`,
		uuid,
	)
	fmt.Println(row)
	if err != nil {
		return err
	}

	return nil
}

func (d *DataBase) UpdateComment(
	ctx context.Context, id string, cmt comment.Comment,
) (comment.Comment, error) {
	cmtRow := CommentRow{
		ID:     id,
		Slug:   sql.NullString{String: cmt.Slug, Valid: true},
		Body:   sql.NullString{String: cmt.Body, Valid: true},
		Author: sql.NullString{String: cmt.Author, Valid: true},
	}
	row, err := d.Client.NamedQueryContext(
		ctx,
		`update comments set
		body=:body, slug=:slug, author=:author
		where id=:id`,
		cmtRow,
	)
	if err != nil {
		return comment.Comment{}, fmt.Errorf("error updating comment: %w", err)
	}
	if err := row.Close(); err != nil {
		return comment.Comment{}, fmt.Errorf("error updating row: %w", err)
	}
	return convertComment(cmtRow), nil
}
