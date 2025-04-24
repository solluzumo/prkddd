package mongodb

import (
	"context"
	"prk/internal/domain/doctype"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DocTypeRepo struct {
	collection *mongo.Collection
}

func NewDocTypeRepository(db *mongo.Database) doctype.Repository {
	return &DocTypeRepo{collection: db.Collection("doc_type")}
}

func (rep *DocTypeRepo) FindByIdDocType(docTypeID string) (*doctype.DocType, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var docType *doctype.DocType
	err := rep.collection.FindOne(ctx, bson.M{"_id": docTypeID}).Decode(&docType)
	if err != nil {
		return nil, err
	}
	return docType, nil
}

func (rep *DocTypeRepo) FindAllDocType() ([]*doctype.DocType, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var data []*doctype.DocType
	cur, err := rep.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	if err := cur.All(ctx, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func (rep *DocTypeRepo) CreateDocType(data *doctype.DocType) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := rep.collection.InsertOne(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

func (rep *DocTypeRepo) DeleteDocType(docTypeID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.M{"id": docTypeID}

	_, err := rep.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}
