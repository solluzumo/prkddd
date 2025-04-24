package mongodb

import (
	"context"
	"prk/internal/domain/userdoc"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserDocRepo struct {
	collection mongo.Collection
}

func NewUserDocRepo(db *mongo.Database) userdoc.Repository {
	return &UserDocRepo{collection: *db.Collection("doc_author")}
}

func (ud *UserDocRepo) ConnectDocumentUser(docUser *userdoc.DocAuthor) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := ud.collection.InsertOne(ctx, docUser)
	return err
}
