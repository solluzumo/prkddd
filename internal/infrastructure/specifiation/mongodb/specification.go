package mongodb

import (
	"prk/internal/domain/filter"

	"go.mongodb.org/mongo-driver/bson"
)

type DocumentSpecification interface {
	ToQuery() interface{}
}

type BasicDocumentSpec struct {
	DocTypeID          *int
	IsUpdatedRegularly *bool
	ExpertReview       *bool
	JournalCategory    string
	Status             filter.DocumentStatus
}



func (s *BasicDocumentSpec) ToQuery() bson.D {
	var query bson.D
	if s.DocTypeID != nil {
		query = append(query, bson.E{Key: "document_type", Value: *s.DocTypeID})
	}
	if s.IsUpdatedRegularly != nil {
		query = append(query, bson.E{Key: "updated_regularly", Value: *s.DocTypeID})
	}
	if s.ExpertReview != nil {
		query = append(query, bson.E{Key: "expert_review", Value: *s.DocTypeID})
	}
	if s.JournalCategory != "" {
		query = append(query, bson.E{Key: "journal_category", Value: *s.DocTypeID})
	}
	if s.Status != "" {
		query = append(query, bson.E{Key: "status", Value: *s.DocTypeID})
	}
	return query
}
