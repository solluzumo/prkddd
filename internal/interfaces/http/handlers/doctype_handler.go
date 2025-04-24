package handlers

import (
	"prk/internal/application/doctype"
)

type DocTypeHandler struct {
	service *doctype.DocTypeService
}

func NewDocTypeHandler(service *doctype.DocTypeService) *DocTypeHandler {
	return &DocTypeHandler{service: service}
}
