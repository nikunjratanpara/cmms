package provisioning

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nikunjratanpara/cmms/provisioning/apis"
)

type ProvisioningModule struct {
}

func (module ProvisioningModule) RegisterApis(applicationRouter *chi.Mux) {
	apiRouter := chi.NewRouter()
	controller := apis.NewTenantController()

	apiRouter.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Provisioning Home"))
	})

	apiRouter.Post("/", controller.Create)
	apiRouter.Get("/{tenantId}", controller.GetById)
	apiRouter.Get("/all", controller.GetAll)

	applicationRouter.Mount("/provisioning", apiRouter)
}
