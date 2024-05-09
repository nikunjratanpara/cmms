package apis

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/nikunjratanpara/cmms/provisioning/requests"
	"github.com/nikunjratanpara/cmms/provisioning/services"
)

type ITenantController interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	GetById(w http.ResponseWriter, r *http.Request)
	/*Update(w http.ResponseWriter, r http.Request)
	Delete(w http.ResponseWriter, r http.Request)
	Get(w http.ResponseWriter, r http.Request)*/
}

type TenantController struct {
	services.ITenantService
}

func NewTenantController() TenantController {
	return TenantController{services.NewTenantService()}
}

func (controller TenantController) Create(w http.ResponseWriter, r *http.Request) {
	var tenantCreateRequest requests.TenantCreateRequest
	err := json.NewDecoder(r.Body).Decode(&tenantCreateRequest)

	if err != nil {
		writeErrorInResponse(w, err)
		return
	}

	createdTenant, err := controller.ITenantService.Create(tenantCreateRequest)
	if err != nil {
		writeErrorInResponse(w, err)
		return
	}
	json.NewEncoder(w).Encode(createdTenant)
}

func (controller TenantController) GetAll(w http.ResponseWriter, r *http.Request) {
	tenants, err := controller.ITenantService.GetAll()
	if err != nil {
		writeErrorInResponse(w, err)
		return
	}
	json.NewEncoder(w).Encode(tenants)
}

func (controller TenantController) GetById(w http.ResponseWriter, r *http.Request) {
	tenantId, err := uuid.Parse(chi.URLParam(r, "tenantId"))
	if err != nil {
		writeErrorInResponse(w, err)
		return
	}

	tenant, err := controller.ITenantService.GetById(tenantId)
	if err != nil {
		writeErrorInResponse(w, err)
		return
	}
	json.NewEncoder(w).Encode(tenant)
}

func writeErrorInResponse(w http.ResponseWriter, err error) {
	log.Fatalln(err)
	w.Write([]byte(err.Error()))
}
