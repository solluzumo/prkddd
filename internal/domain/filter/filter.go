package filter

type DocumentFilter struct {
	TypeID             *int
	IsUpdatedRegularly *bool
	ExpertReview       *bool
	JournalCategory    string
	Status             DocumentStatus
}

type DocumentStatus string

const (
	Draft     DocumentStatus = "draft"
	Published DocumentStatus = "published"
	Archived  DocumentStatus = "archived"
)
