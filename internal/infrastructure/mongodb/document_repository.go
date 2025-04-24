package mongodb

import (
	"context"
	"errors"
	"prk/internal/domain/document"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DocumentRepository struct {
	collection *mongo.Collection
}

func NewDocumentRepository(db *mongo.Database) document.Repository {
	return &DocumentRepository{
		collection: db.Collection("documents"),
	}
}

func (r *DocumentRepository) ExistsDocument(fileName string) (bool, error) {
	_, err := r.FindDocumentByName(fileName)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, nil // документ не найден — значит, не существует
		}
		return false, err // реальная ошибка
	}
	return true, nil // найден — существует
}

func (rep *DocumentRepository) CreateDocument(doc *document.Document) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := rep.collection.InsertOne(ctx, doc)
	if err != nil {
		return err
	}
	return nil
}

func (rep *DocumentRepository) FindDocumentByName(filesName string) (*document.Document, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var object *document.Document
	err := rep.collection.FindOne(ctx, bson.M{"files_name": filesName}).Decode(&object)
	if err != nil {
		return nil, err
	}
	return object, nil
}

func (rep *DocumentRepository) FindDocumentById(objectID string) (*document.Document, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var object *document.Document
	err := rep.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&object)
	if err != nil {
		return nil, err
	}
	return object, nil
}

func (rep *DocumentRepository) FindDocuments() ([]*document.Document, int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	total, err := rep.collection.CountDocuments(ctx, bson.D{})
	if err != nil {
		return nil, 0, err
	}
	var data []*document.Document
	cur, err := rep.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, 0, err
	}
	defer cur.Close(ctx)
	if err := cur.All(ctx, &data); err != nil {
		return nil, 0, err
	}

	return data, total, nil
}

func (rep *DocumentRepository) TouchExperReview(docID string, value bool) error {
	res := rep.collection.FindOneAndUpdate(context.Background(), bson.M{"_id": docID}, bson.M{"expert_review": value})
	if err := res.Err(); err != nil {
		return err
	}
	return nil
}

func (rep *DocumentRepository) TouchDate(docID string, value time.Time) error {
	res := rep.collection.FindOneAndUpdate(context.Background(), bson.M{"_id": docID}, bson.M{"updated_at": value})
	if err := res.Err(); err != nil {
		return err
	}
	return nil
}

func (rep *DocumentRepository) DeleteDocument(documentID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := rep.collection.DeleteOne(ctx, bson.M{"id": documentID})
	if err != nil {
		return err
	}
	return nil
}
