package document

import "mime/multipart"

type CreateDocumentDTO struct {
	DocumentTypeID    string
	Title             string
	Date              string
	FilesName         string
	UpdatedRegularly  bool
	ExpertReview      bool
	JournalCategoryID string
	Source            string
	UploadedBy        string
	MainFile          multipart.File
	MainFileName      string
	AdditionalFiles   []*multipart.FileHeader
}

type ListDoucmentDTO struct {
	Page      int
	Limit     int
	SortField string
	SortOrder string
	Filters   map[string]interface{}
}
