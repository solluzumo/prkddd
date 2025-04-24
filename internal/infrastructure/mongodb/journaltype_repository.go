package mongodb

import (
	"context"
	"prk/internal/domain/journaltype"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type JournalTypeRepo struct {
	collection *mongo.Collection
}

func NewJournalTypeRepository(db *mongo.Database) journaltype.Repository {
	return &JournalTypeRepo{collection: db.Collection("journal_type")}
}

func (rep *JournalTypeRepo) FindByIdJournalType(id string) (*journaltype.JournalType, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var jt *journaltype.JournalType
	err := rep.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&jt)
	if err != nil {
		return nil, err
	}
	return jt, nil
}

func (rep *JournalTypeRepo) FindAllJournalType() ([]*journaltype.JournalType, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var results []*journaltype.JournalType
	cur, err := rep.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	if err := cur.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}

func (rep *JournalTypeRepo) CreateJournalType(jt *journaltype.JournalType) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := rep.collection.InsertOne(ctx, jt)
	return err
}

func (rep *JournalTypeRepo) DeleteJournalType(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := rep.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
