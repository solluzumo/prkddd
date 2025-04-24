package journaltype

import (
	"prk/internal/domain/idgen"
	"prk/internal/domain/journaltype"
)

type JournalTypeService struct {
	journalTypeRepo journaltype.Repository
	uuidGen         idgen.UUIDGenerator
}

func NewService(repo journaltype.Repository) *JournalTypeService {
	return &JournalTypeService{journalTypeRepo: repo}
}

func (s *JournalTypeService) CreateJournalType(dto CreateJournalTypeDTO) error {
	newJournalType := &journaltype.JournalType{
		ID:   s.uuidGen.New(),
		Name: dto.Name,
	}
	return s.journalTypeRepo.CreateJournalType(newJournalType)
}

func (s *JournalTypeService) DeleteJournalType(journalTypeID string) error {
	return s.journalTypeRepo.DeleteJournalType(journalTypeID)
}

func (s *JournalTypeService) GetOneJournalType(journalTypeID string) (*journaltype.JournalType, error) {
	return s.journalTypeRepo.FindByIdJournalType(journalTypeID)
}

func (s *JournalTypeService) GetAllJournalType() ([]*journaltype.JournalType, error) {
	return s.journalTypeRepo.FindAllJournalType()
}
