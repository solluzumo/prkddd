package document

import (
	"context"
	"errors"
	"fmt"
	"path/filepath"
	"prk/internal/application/auth"
	"prk/internal/domain/doctype"
	"prk/internal/domain/document"
	"prk/internal/domain/idgen"
	"prk/internal/domain/journaltype"
	"prk/internal/domain/user"
	"prk/internal/domain/userdoc"
	"prk/pkg/utils"
	"strings"
	"time"
)

type DocumentService struct {
	docRepo     document.Repository
	docTypeRepo doctype.Repository
	journalRepo journaltype.Repository
	userRepo    user.Repository
	fileRepo    document.FileStorage
	userDocRepo userdoc.Repository
	uuidGen     idgen.UUIDGenerator
}

func NewService(repo document.Repository,
	dt doctype.Repository,
	jt journaltype.Repository,
	us user.Repository,
	fs document.FileStorage) *DocumentService {
	return &DocumentService{
		docRepo:     repo,
		docTypeRepo: dt,
		journalRepo: jt,
		userRepo:    us,
		fileRepo:    fs,
	}
}

func (s *DocumentService) CreateDocument(ctx context.Context, token string, dto CreateDocumentDTO) error {

	docType, err := s.docTypeRepo.FindByIdDocType(dto.DocumentTypeID)
	if err != nil {
		return err
	}

	journalType, err := s.journalRepo.FindByIdJournalType(dto.JournalCategoryID)
	if err != nil {
		return err
	}

	userID, err := auth.ExtractUserID(token)
	if err != nil {
		return fmt.Errorf("token invalid: %w", err)
	}

	user, err := s.userRepo.FindByIDUser(userID)
	if err != nil {
		return fmt.Errorf("user not found: %w", err)
	}
	authorName := user.Name
	authorNameParsed := utils.SlugifyName(authorName)
	titleParsed := utils.SlugifyName(dto.Title)
	filesname := authorNameParsed + "-" + titleParsed

	exists, err := s.docRepo.ExistsDocument(filesname)
	if err != nil {
		return fmt.Errorf("failed to check document existence: %w", err)
	}
	if exists {
		return errors.New("document already exists")
	}

	mainFileExt := strings.TrimPrefix(filepath.Ext(dto.MainFileName), ".")
	fileId, err := s.fileRepo.UploadFile(dto.MainFile, filesname+"-main-"+mainFileExt)
	if err != nil {
		return err
	}
	dto.MainFile.Close()

	var additionFileIDs []string
	count := 1
	for _, fileHeader := range dto.AdditionalFiles {
		additionalFile, err := fileHeader.Open()
		if err != nil {
			return err
		}

		additionalFileName := fileHeader.Filename
		additionalFileExt := strings.TrimPrefix(filepath.Ext(additionalFileName), ".")

		fullName := fmt.Sprintf("%s-ad%d-%s", filesname, count, additionalFileExt)
		additionalFileID, err := s.fileRepo.UploadFile(additionalFile, fullName)
		additionalFile.Close()
		if err != nil {
			return err
		}
		count++
		additionFileIDs = append(additionFileIDs, additionalFileID)
	}

	t, err := time.Parse("2006-01-02", dto.Date)
	if err != nil {
		return err
	}
	formattedDate := t.Format(time.RFC3339)
	now := time.Now().Format(time.RFC3339)
	document := document.Document{
		ID:               s.uuidGen.New(),
		DocumentType:     docType.DocTypeID,
		Title:            dto.Title,
		Date:             formattedDate,
		CreatedAt:        now,
		UpdatedAt:        now,
		FilesName:        filesname,
		UpdatedRegularly: dto.UpdatedRegularly,
		ExpertReview:     dto.ExpertReview,
		JournalCategory:  journalType.ID,
		Source:           dto.Source,
		MainFile:         fileId,
		Addition:         additionFileIDs,
	}

	err = s.docRepo.CreateDocument(&document)
	if err != nil {
		return err
	}
	docUser := &userdoc.DocAuthor{
		ID:         s.uuidGen.New(),
		UserID:     userID,
		DocumentID: document.ID,
	}
	err = s.userDocRepo.ConnectDocumentUser(docUser)
	return err
}

func (s *DocumentService) DeleteDocument(docID string) error {
	return s.docRepo.DeleteDocument(docID)
}

func (s *DocumentService) TouchDate(docID string, newDate time.Time) error {
	return s.docRepo.TouchDate(docID, newDate)
}

func (s *DocumentService) TouchExperReview(docID string, newReview bool) error {
	return s.docRepo.TouchExperReview(docID, newReview)
}

func (s *DocumentService) FindDocumentById(docID string) (*document.Document, error) {
	return s.docRepo.FindDocumentById(docID)
}

func (s *DocumentService) FindDocumentByName(filesName string) (*document.Document, error) {
	return s.docRepo.FindDocumentByName(filesName)
}

func (s *DocumentService) FindDocuments(ctx context.Context, dto ListDoucmentDTO) ([]*document.Document, int64, error) {
	//проверять не пусты ли фильтры и сортировки и вызывать разные методы бд
	fmt.Println(dto.Filters, dto.Limit, dto.SortField, dto.SortOrder)
	return s.docRepo.FindDocuments()
}

// func (s *DocumentService) GetAllDocumentsByUser(userID string) ([]*document.Document, error) {
// 	return s.docRepo.GetAllByUser(userID)
// }
