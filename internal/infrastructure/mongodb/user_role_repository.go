package mongodb

import (
	"context"
	"fmt"
	"prk/internal/domain/user"
	"prk/internal/domain/userrole"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRoleRepo struct {
	collection     mongo.Collection
	userCollection mongo.Collection
	roleCollection mongo.Collection
}

func NewUserRolerepo(db *mongo.Database) userrole.Repository {
	return &UserRoleRepo{
		collection:     *db.Collection("user_role"),
		userCollection: *db.Collection("user"),
		roleCollection: *db.Collection("roles"),
	}
}

func (ur *UserRoleRepo) CreateUserRole(userRole *userrole.UserRole) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := ur.collection.InsertOne(ctx, userRole)
	return err
}

func (ur *UserRoleRepo) FindRoleByUser(userID string) ([]*user.Role, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var connections []userrole.UserRole
	cursor, err := ur.collection.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return nil, fmt.Errorf("failed to find user-role connections: %w", err)
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &connections); err != nil {
		return nil, fmt.Errorf("failed to decode user-role connections: %w", err)
	}

	if len(connections) == 0 {
		return []*user.Role{}, nil
	}
	roleIDSet := make(map[string]struct{})
	for _, conn := range connections {
		roleIDSet[conn.RoleID] = struct{}{}
	}
	var roleIds []string
	for id := range roleIDSet {
		roleIds = append(roleIds, id)
	}
	filter := bson.M{"_id": bson.M{"$in": roleIds}}
	roleCursor, err := ur.roleCollection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch users by ids: %w", err)
	}
	defer roleCursor.Close(ctx)
	var roles []*user.Role
	if err := roleCursor.All(ctx, &roles); err != nil {
		return nil, fmt.Errorf("failed to decode users: %w", err)
	}

	return roles, nil
}

func (ur *UserRoleRepo) FindRoleByName(roleName string) (*user.Role, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var u *user.Role
	err := ur.collection.FindOne(ctx, bson.M{"name": roleName}).Decode(&u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (ur *UserRoleRepo) FindUserByRole(roleID string) ([]*user.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var connections []userrole.UserRole
	cursor, err := ur.collection.Find(ctx, bson.M{"role_id": roleID})
	if err != nil {
		return nil, fmt.Errorf("failed to find user-role connections: %w", err)
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &connections); err != nil {
		return nil, fmt.Errorf("failed to decode user-role connections: %w", err)
	}

	if len(connections) == 0 {
		return []*user.User{}, nil
	}

	userIDSet := make(map[string]struct{})
	for _, conn := range connections {
		userIDSet[conn.UserID] = struct{}{}
	}
	var userIds []string
	for id := range userIDSet {
		userIds = append(userIds, id)
	}

	filter := bson.M{"_id": bson.M{"$in": userIds}}
	userCursor, err := ur.userCollection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch users by ids: %w", err)
	}
	defer userCursor.Close(ctx)

	var users []*user.User
	if err := userCursor.All(ctx, &users); err != nil {
		return nil, fmt.Errorf("failed to decode users: %w", err)
	}

	return users, nil
}
