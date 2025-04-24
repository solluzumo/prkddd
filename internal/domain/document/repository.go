package document

import (
	"mime/multipart"
	"time"
)

type Repository interface {
	CreateDocument(doc *Document) error
	DeleteDocument(docID string) error
	FindDocumentById(docID string) (*Document, error)
	FindDocumentByName(docID string) (*Document, error)
	FindDocuments() ([]*Document, int64, error)
	// FindAllByUser(userId string) ([]*Document, error)
	TouchDate(docID string, newDate time.Time) error
	TouchExperReview(docID string, newReview bool) error
	ExistsDocument(fileName string) (bool, error)
}

type FileStorage interface {
	UploadFile(file multipart.File, filename string) (string, error)
}
