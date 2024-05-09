package contract

import "github.com/go-chi/chi/v5"

type IModule interface {
	RegisterApis(applicationRouter *chi.Mux) // mount new router to application router for the module
}
