package handlers

import "prk/internal/application/userrole"

type UserRoleHandler struct {
	service *userrole.UserRoleService
}

func NewUserRoleHandler(service *userrole.UserRoleService) *UserRoleHandler {
	return &UserRoleHandler{service: service}
}
