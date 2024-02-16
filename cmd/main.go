package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nikunjratanpara/cmms/internal/contract"
	"github.com/nikunjratanpara/cmms/provisioning"
)

func main() {
	router := chi.NewRouter()
	mountApiRoutes(router)
	http.ListenAndServe(":8081", router)
}

func mountApiRoutes(router *chi.Mux) {
	modules := [...]contract.IModule{
		&provisioning.ProvisioningModule{},
	}

	for _, module := range modules {
		module.RegisterApis(router)
	}
}
