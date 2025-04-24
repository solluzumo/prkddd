package userdoc

import (
	"prk/internal/domain/idgen"
	"prk/internal/domain/userdoc"
)

type UserDocService struct {
	udRepo  userdoc.Repository
	uuidGen idgen.UUIDGenerator
}

func NewUserDocService(repo userdoc.Repository) *UserDocService {
	return &UserDocService{udRepo: repo}
}

func (ud *UserDocService) ConnectUserToRole(dto UserDocDTO) error {
	newAuthor := &userdoc.DocAuthor{
		ID:         ud.uuidGen.New(),
		UserID:     dto.UserID,
		DocumentID: dto.DocumentID,
	}
	return ud.udRepo.ConnectDocumentUser(newAuthor)
}
