package services

import (
	"time"

	"github.com/google/uuid"
	"github.com/nikunjratanpara/cmms/provisioning/entities"
	"github.com/nikunjratanpara/cmms/provisioning/requests"
	"github.com/nikunjratanpara/core"

	repositories "github.com/nikunjratanpara/cmms/provisioning/repositories"
)

type ITenantService interface {
	Create(createRequest requests.TenantCreateRequest) (*entities.Tenant, error)
	GetAll() ([]entities.Tenant, error)
	GetById(tenantId uuid.UUID) (*entities.Tenant, error)
	/*Update(id string, updateRequest TenantCreateRequest) infrastructure.Tenant
	Delete(id string) infrastructure.Tenant
	*/
}

type TenantService struct {
	*repositories.TenantRepository
}

func (service TenantService) Create(createRequest requests.TenantCreateRequest) (*entities.Tenant, error) {
	createdBy := uuid.New()
	createdAt := time.Now().UnixMilli()
	tenant := entities.Tenant{
		TenantCode: createRequest.TenantCode,
		Name:       createRequest.Name,
		TenantId:   uuid.New(),
		Deleted:    false,
		Audit:      core.NewAudit(createdBy, createdAt, createdBy, createdAt),
	}
	createdTenant, err := service.TenantRepository.Insert(tenant)
	if err != nil {
		return nil, err
	}
	return createdTenant, err
}

func (service TenantService) GetAll() ([]entities.Tenant, error) {
	return service.TenantRepository.GetAll()
}

func (service TenantService) GetById(tenantId uuid.UUID) (*entities.Tenant, error) {
	tenant, err := service.TenantRepository.GetById(tenantId)
	return tenant, err
}

func NewTenantService() ITenantService {
	return TenantService{repositories.NewTenantRepository()}
}
