package mongodb

import (
	"context"
	"time"

	"prk/internal/domain/user"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) user.Repository {
	return &UserRepo{
		collection: db.Collection("users"),
	}
}

func (repo *UserRepo) FindByIDUser(id string) (*user.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var u *user.User
	err := repo.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (repo *UserRepo) FindAllUser() ([]*user.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var users []*user.User
	cursor, err := repo.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *UserRepo) CreateUser(u *user.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := repo.collection.InsertOne(ctx, u)
	return err
}

func (repo *UserRepo) DeleteUser(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := repo.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
