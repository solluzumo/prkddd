package router

import (
	"net/http"
	"prk/internal/app"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func SetupWeb(appInstance *app.App) *chi.Mux {
	r := chi.NewRouter()
	docHandler := appInstance.Handlers.DocumentHandler
	userHandler := appInstance.Handlers.UserDocHandler
	userRoleHandler := appInstance.Handlers.UserRoleHandler
	userDocHandler := appInstance.Handlers.UserDocHandler
	docTypeHandler := appInstance.Handlers.DocTypeHandler
	journaltypeHandler := appInstance.Handlers.JournalTypeHandler

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API Gateway"))
	})
	r.Route("/documents", func(r chi.Router) {
		r.Get("/", docHandler.GetDocuments)
		r.Post("/", docHandler.AddDocument)
		r.Route("/{documentID}", func(r chi.Router) {
			r.Use(docHandler.DocumentCtx)
			r.Get("/", docHandler.GetDocument)
			r.Put("/", docHandler.UpdateDocument)
			r.Patch("/", docHandler.PartialUpdateDocument)
			r.Delete("/", docHandler.DeleteDocument)
		})
	})
	r.Route("/document-types", func(r chi.Router) {
		r.Get("/", docHandler.GetDocumentTypes)
		r.Post("/", docHandler.AddDocumentType)

		r.Route("/{typeID}", func(r chi.Router) {
			r.Get("/", docHandler.GetDocumentType)
			r.Put("/", docHandler.UpdateDocumentType)
			r.Delete("/", docHandler.DeleteDocumentType)
		})
	})

	r.Route("/users", func(r chi.Router) {
		r.Get("/", userHandler.GetUsers)
		r.Post("/", userHandler.CreateUser)

		r.Route("/{userID}", func(r chi.Router) {
			r.Get("/", userHandler.GetUser)
			r.Put("/", userHandler.UpdateUser)
			r.Delete("/", userHandler.DeleteUser)

			// Маршруты для ролей пользователя
			r.Get("/roles", userHandler.GetUserRoles)
			r.Post("/roles", userHandler.AssignRoleToUser)
			r.Delete("/roles/{roleID}", userHandler.RevokeRoleFromUser)
		})
	})

	r.Route("/roles", func(r chi.Router) {
		r.Get("/", userHandler.GetAllRoles)
		r.Post("/", userHandler.CreateRole)

		r.Route("/{roleID}", func(r chi.Router) {
			r.Get("/", userHandler.GetRole)
			r.Put("/", userHandler.UpdateRole)
			r.Delete("/", userHandler.DeleteRole)
		})
	})
	return r
}
