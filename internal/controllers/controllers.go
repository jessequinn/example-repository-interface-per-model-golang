package controllers

import (
	"fmt"
	"net/http"

	"example-repository-interface-per-model-golang/internal/models"
)

// BaseHandler will hold everything that controller needs
type BaseHandler struct {
	userRepo models.UserRepository
}

// NewBaseHandler returns a new BaseHandler
func NewBaseHandler(userRepo models.UserRepository) *BaseHandler {
	return &BaseHandler{
		userRepo: userRepo,
	}
}

// HelloWorld returns Hello, World
func (h *BaseHandler) HelloWorld(w http.ResponseWriter, r *http.Request) {
	if user, err := h.userRepo.FindByID(1); err != nil {
		fmt.Println("Error", user)
	}

	w.Write([]byte("Hello, World"))
}
