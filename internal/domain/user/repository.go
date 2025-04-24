package user

type Repository interface {
	CreateUser(user *User) error
	FindByIDUser(userID string) (*User, error)
	FindAllUser() ([]*User, error)
	DeleteUser(userID string) error
}
