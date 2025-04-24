package userrole

import (
	"prk/internal/domain/idgen"
	"prk/internal/domain/userrole"
)

type UserRoleService struct {
	urRepo  userrole.Repository
	uuidGen idgen.UUIDGenerator
}

func NewUserRoleService(repo userrole.Repository) *UserRoleService {
	return &UserRoleService{urRepo: repo}
}

func (ur *UserRoleService) ConnectUserToRole(dto UserRoleDTO) error {
	newConnection := &userrole.UserRole{
		ID:     ur.uuidGen.New(),
		UserID: dto.UserID,
		RoleID: dto.RoleID,
	}
	return ur.urRepo.CreateUserRole(newConnection)
}
