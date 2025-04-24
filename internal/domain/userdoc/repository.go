package userdoc

type Repository interface {
	ConnectDocumentUser(docUser *DocAuthor) error
}
