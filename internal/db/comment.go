package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/by-sabbir/go-rest/internal/comment"
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
		where id=$1`,
		uuid,
	)
	err := row.Scan(&cmtRow.ID, &cmtRow.Body, &cmtRow.Slug, &cmtRow.Author)
	if err != nil {
		return comment.Comment{}, fmt.Errorf("error fetching comments from uuid: %+v", err)
	}
	return convertComment(cmtRow), nil
}
