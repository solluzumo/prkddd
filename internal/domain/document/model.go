package document

type Document struct {
	ID               string
	DocumentType     string
	Title            string
	Date             string
	CreatedAt        string
	UpdatedAt        string
	FilesName        string
	UpdatedRegularly bool
	ExpertReview     bool
	JournalCategory  string
	Source           string
	MainFile         string
	Addition         []string
	Status           string
}
