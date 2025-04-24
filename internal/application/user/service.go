package user

import (
	"prk/internal/domain/idgen"
	"prk/internal/domain/user"
	"prk/internal/domain/userrole"
)

type UserService struct {
	userRepo     user.Repository
	userRoleRepo userrole.Repository
	uuidGen      idgen.UUIDGenerator
}

func NewService(repo user.Repository) *UserService {
	return &UserService{userRepo: repo}
}

func (s *UserService) CreateUser(dto CreateUserDTO) error {
	newUser := &user.User{
		ID:   s.uuidGen.New(),
		Name: dto.Name,
	}
	err := s.userRepo.CreateUser(newUser)
	if err != nil {
		return err
	}
	role, err := s.userRoleRepo.FindRoleByName(dto.RoleName)
	if err != nil {
		return nil
	}
	newUserRole := &userrole.UserRole{
		ID:     s.uuidGen.New(),
		UserID: newUser.ID,
		RoleID: role.ID,
	}

	err = s.userRoleRepo.CreateUserRole(newUserRole)

	return err
}

func (s *UserService) GetUserById(userID string) (*user.User, error) {
	return s.userRepo.FindByIDUser(userID)
}
func (s *UserService) GetAllUser() ([]*user.User, error) {
	return s.userRepo.FindAllUser()
}
func (s *UserService) DeleteUser(userID string) error {
	return s.userRepo.DeleteUser(userID)
}
