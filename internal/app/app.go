package app

import (
	//Сервисы
	doctypeS "prk/internal/application/doctype"
	documentS "prk/internal/application/document"
	idgenS "prk/internal/application/idgen"
	journaltypeS "prk/internal/application/journaltype"
	userS "prk/internal/application/user"
	userdocS "prk/internal/application/userdoc"
	userroleS "prk/internal/application/userrole"
	"prk/internal/interfaces/http/handlers"

	//Репозитории
	doctypeR "prk/internal/domain/doctype"
	documentR "prk/internal/domain/document"
	idgenR "prk/internal/domain/idgen"
	journaltypeR "prk/internal/domain/journaltype"
	userR "prk/internal/domain/user"
	userdocR "prk/internal/domain/userdoc"
	userroleR "prk/internal/domain/userrole"

	//Инфраструктура
	"prk/internal/infrastructure/mongodb"

	//Остальное
	"prk/internal/config"

	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	Config   *config.Config
	DB       *mongo.Database
	Repos    *Repositories
	Services *Services
	Handlers *Handlers
	IDGen    idgenR.UUIDGenerator
}

type Handlers struct {
	DocumentHandler    handlers.DocumentHandler
	UserHandler        handlers.UserHandler
	DocTypeHandler     handlers.DocTypeHandler
	JournalTypeHandler handlers.JournalTypeHandler
	UserDocHandler     handlers.UserDocHandler
	UserRoleHandler    handlers.UserRoleHandler
}

type Repositories struct {
	UserRepo        userR.Repository
	DocumentRepo    documentR.Repository
	DocTypeRepo     doctypeR.Repository
	JournalTypeRepo journaltypeR.Repository
	UserDocRepo     userdocR.Repository
	UserRoleRepo    userroleR.Repository
	FSRepo          documentR.FileStorage
}

type Services struct {
	UserService        *userS.UserService
	DocumentService    *documentS.DocumentService
	DocTypeService     *doctypeS.DocTypeService
	JournalTypeService *journaltypeS.JournalTypeService
	UserDocService     *userdocS.UserDocService
	UserRoleService    *userroleS.UserRoleService
}

func New(cfg *config.Config, db *mongo.Database) *App {
	uuidGen := &idgenS.DefaultUUIDGenerator{}
	repos := &Repositories{
		UserRepo:        mongodb.NewUserRepository(db),
		DocumentRepo:    mongodb.NewDocumentRepository(db),
		DocTypeRepo:     mongodb.NewDocTypeRepository(db),
		JournalTypeRepo: mongodb.NewJournalTypeRepository(db),
		UserDocRepo:     mongodb.NewUserDocRepo(db),
		UserRoleRepo:    mongodb.NewUserRolerepo(db),
		FSRepo:          mongodb.NewFSRepository(db),
	}
	services := &Services{
		UserService:        userS.NewService(repos.UserRepo),
		DocumentService:    documentS.NewService(repos.DocumentRepo, repos.DocTypeRepo, repos.JournalTypeRepo, repos.UserRepo, repos.FSRepo),
		DocTypeService:     doctypeS.NewService(repos.DocTypeRepo),
		JournalTypeService: journaltypeS.NewService(repos.JournalTypeRepo),
		UserDocService:     userdocS.NewUserDocService(repos.UserDocRepo),
		UserRoleService:    userroleS.NewUserRoleService(repos.UserRoleRepo),
	}
	handlers := &Handlers{
		DocumentHandler:    *handlers.NewDocumentHandler(services.DocumentService),
		UserHandler:        *handlers.NewUserHandler(services.UserService),
		DocTypeHandler:     *handlers.NewDocTypeHandler(services.DocTypeService),
		JournalTypeHandler: *handlers.NewJournalTypeHandler(services.JournalTypeService),
		UserDocHandler:     *handlers.NewUserDocHandler(services.UserDocService),
		UserRoleHandler:    *handlers.NewUserRoleHandler(services.UserRoleService),
	}
	return &App{
		Config:   cfg,
		Repos:    repos,
		Services: services,
		IDGen:    uuidGen,
		Handlers: handlers,
	}
}
