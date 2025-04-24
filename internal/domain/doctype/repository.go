package doctype

type Repository interface {
	CreateDocType(docType *DocType) error
	DeleteDocType(docTypeID string) error
	FindByIdDocType(docTypeID string) (*DocType, error)
	FindAllDocType() ([]*DocType, error)
}
