package journaltype

type Repository interface {
	CreateJournalType(journalType *JournalType) error
	DeleteJournalType(journalTypeID string) error
	FindByIdJournalType(journalTypeID string) (*JournalType, error)
	FindAllJournalType() ([]*JournalType, error)
}
