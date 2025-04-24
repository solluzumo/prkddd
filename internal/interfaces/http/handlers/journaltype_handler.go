package handlers

import "prk/internal/application/journaltype"

type JournalTypeHandler struct {
	service *journaltype.JournalTypeService
}

func NewJournalTypeHandler(service *journaltype.JournalTypeService) *JournalTypeHandler {
	return &JournalTypeHandler{service: service}
}
