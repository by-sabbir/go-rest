package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrorFetchingComment = errors.New("could not find the comment")
	ErrorNotImplemented  = errors.New("not implemented")
)

// Store - this interface defines all the methods to operate
type Store interface {
	GetComment(context.Context, string) (Comment, error)
}

// Comment - representation of a comment structure
type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

// Service - is the struct containing business logics
type Service struct {
	Store Store
}

// NewService - returns to a pointer to a new service
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("Retreiving the comment")

	cmt, err := s.Store.GetComment(ctx, id)

	if err != nil {
		fmt.Println(err)
		return Comment{}, ErrorFetchingComment
	}
	return cmt, nil
}

func (s *Service) UpdateComment(ctx context.Context, cmt Comment) error {
	return ErrorNotImplemented
}

func (s *Service) DeleteComment(ctx context.Context, id string) (Comment, error) {
	return Comment{}, ErrorNotImplemented
}
