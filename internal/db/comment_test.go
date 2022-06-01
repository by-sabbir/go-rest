//go:build integration
// +build integration

package db

import (
	"context"
	"testing"

	"github.com/by-sabbir/go-rest/internal/comment"
	"github.com/stretchr/testify/assert"
)

func TestCommentDB(t *testing.T) {
	db, err := NewDatabase()

	assert.NoError(t, err)
	cmt, err := db.PostComment(context.Background(), comment.Comment{
		Body:   "body",
		Slug:   "slug",
		Author: "author",
	})

	assert.NoError(t, err)

	gotCmt, err := db.GetComment(context.Background(), cmt.ID)
	assert.NoError(t, err)
	assert.Equal(t, "body", gotCmt.Body)
	assert.Equal(t, "slug", gotCmt.Slug)
	assert.Equal(t, "author", gotCmt.Author)

	updatedComment, errUpdate := db.UpdateComment(context.Background(), gotCmt.ID, comment.Comment{
		Body:   "updated",
		Slug:   "updatedSlug",
		Author: "updated author",
	})
	assert.NoError(t, errUpdate)
	assert.Equal(t, gotCmt.ID, updatedComment.ID)
	assert.NotEqual(t, gotCmt.Body, updatedComment.Body)
	assert.Equal(t, updatedComment.Slug, "updatedSlug")

}
