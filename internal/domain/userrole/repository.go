package userrole

import "prk/internal/domain/user"

type Repository interface {
	CreateUserRole(userRole *UserRole) error
	FindRoleByName(roleName string) (*user.Role, error)
	FindUserByRole(roleID string) ([]*user.User, error)
	FindRoleByUser(userID string) ([]*user.Role, error)
}
