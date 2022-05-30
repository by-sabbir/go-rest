package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/by-sabbir/go-rest/internal/comment"
	"github.com/gorilla/mux"
)

type CommentService interface {
	GetComment(context.Context, string) (comment.Comment, error)
	PostComment(context.Context, comment.Comment) (comment.Comment, error)
	DeleteComment(context.Context, string) error
	UpdateComment(context.Context, string, comment.Comment) (comment.Comment, error)
}
type Details struct {
	Message string
}

func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	var cmt comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&cmt); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	cmt, err := h.Service.PostComment(r.Context(), cmt)
	log.Println("Got Comment: ", cmt)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(cmt); err != nil {
		panic(err)
	}
}

func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	log.Println("ID: ", id, vars)
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(Details{Message: "id is required"}); err != nil {
			panic(err)
		}
	}
	var cmt comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&cmt); err != nil {
		log.Println(err)
	}

	cmt, err := h.Service.GetComment(r.Context(), id)
	if err != nil {
		log.Println(err)
		return
	}

	if err := json.NewEncoder(w).Encode(cmt); err != nil {
		panic(err)
	}

}

func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(Details{Message: "id is required"}); err != nil {
			panic(err)
		}
	}

	var updateComment comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&updateComment); err != nil {
		log.Println(err)
	}

	updateComment, err := h.Service.UpdateComment(r.Context(), id, updateComment)
	if err != nil {
		log.Println(err)
	}
	if err := json.NewEncoder(w).Encode(updateComment); err != nil {
		log.Println(err)
		panic(err)
	}
}

func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(Details{Message: "id is required"}); err != nil {
			panic(err)
		}
	}
	err := h.Service.DeleteComment(r.Context(), id)
	if err != nil {
		log.Println(err)
		return
	}
	if err := json.NewEncoder(w).Encode(Details{Message: "successfully deleted"}); err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusNoContent)
}
