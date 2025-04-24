package doctype

import (
	"prk/internal/domain/doctype"
	"prk/internal/domain/idgen"
)

type DocTypeService struct {
	docTypeRepo doctype.Repository
	uuidGen     idgen.UUIDGenerator
}

func NewService(repo doctype.Repository) *DocTypeService {
	return &DocTypeService{docTypeRepo: repo}
}

func (s *DocTypeService) CreateDocType(dto CreateDocTypeDTO) error {
	newDocType := &doctype.DocType{
		DocTypeID:   s.uuidGen.New(),
		DocTypeName: dto.Name,
	}
	return s.docTypeRepo.CreateDocType(newDocType)
}

func (s *DocTypeService) DeleteDocType(docTypeID string) error {
	return s.docTypeRepo.DeleteDocType(docTypeID)
}

func (s *DocTypeService) GetOneDocType(docTypeID string) (*doctype.DocType, error) {
	return s.docTypeRepo.FindByIdDocType(docTypeID)
}

func (s *DocTypeService) GetAllDocType() ([]*doctype.DocType, error) {
	return s.docTypeRepo.FindAllDocType()
}
