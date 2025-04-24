package handlers

import "prk/internal/application/userdoc"

type UserDocHandler struct {
	service *userdoc.UserDocService
}

func NewUserDocHandler(service *userdoc.UserDocService) *UserDocHandler {
	return &UserDocHandler{service: service}
}
