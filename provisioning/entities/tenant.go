package entities

import (
	"github.com/google/uuid"
	"github.com/nikunjratanpara/cmms/internal/contract"
)

type Tenant struct {
	contract.Audit
	TenantId   uuid.UUID
	Name       string
	TenantCode string
	Deleted    bool
}

func (tenant Tenant) GetId() uuid.UUID {
	return tenant.TenantId
}
