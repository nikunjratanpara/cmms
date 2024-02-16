package infrastructure

import (
	"context"

	"github.com/google/uuid"
	"github.com/nikunjratanpara/cmms/internal/infrastructure"
	"github.com/nikunjratanpara/cmms/provisioning/entities"
)

type TenantRepository struct {
	infrastructure.Repository
}

func NewTenantRepository() *TenantRepository {
	return &TenantRepository{infrastructure.Repository{DatabaseConnectionProvider: &infrastructure.DatabaseConnectionProvider{}}}
}

func (tenantRepository TenantRepository) Insert(tenant entities.Tenant) (*entities.Tenant, error) {
	conn := tenantRepository.GetConnetion()
	query := "Insert into \"Provisioning\".\"Tenant\" (\"TenantId\",\"Name\", \"TenantCode\", \"Deleted\", \"CreatedAt\", \"CreatedBy\", \"ModifiedAt\", \"ModifiedBy\" ) Values ($1, $2, $3, $4, $5, $6, $7, $8)"

	commandTag, err := conn.Exec(context.Background(), query, tenant.TenantId, tenant.Name, tenant.TenantCode, tenant.Deleted, tenant.CreatedAt, tenant.CreatedBy, tenant.ModifiedAt, tenant.ModifiedBy)
	if err != nil {
		return nil, err
	}
	if commandTag.RowsAffected() == 1 {
		return tenantRepository.GetById(tenant.GetId())
	}
	return nil, nil //Unlikely to occure
}

func (tenantRepository TenantRepository) GetById(tenantId uuid.UUID) (*entities.Tenant, error) {
	conn := tenantRepository.GetConnetion()
	query := "Select \"TenantId\", \"Name\", \"TenantCode\",\"Deleted\", \"CreatedAt\", \"CreatedBy\", \"ModifiedAt\", \"ModifiedBy\" from \"Provisioning\".\"Tenant\" where \"TenantId\" = $1 "
	row := conn.QueryRow(context.Background(), query, tenantId)
	tenant := &entities.Tenant{}
	err := row.Scan(&tenant.TenantId, &tenant.Name, &tenant.TenantCode, &tenant.Deleted, &tenant.CreatedAt, &tenant.CreatedBy, &tenant.ModifiedAt, &tenant.ModifiedBy)

	if err != nil {
		return nil, err
	}
	return tenant, err
}

func (tenantRepository TenantRepository) GetAll() ([]entities.Tenant, error) {
	conn := tenantRepository.GetConnetion()
	query := "Select \"TenantId\", \"Name\", \"TenantCode\",\"Deleted\", \"CreatedAt\", \"CreatedBy\", \"ModifiedAt\", \"ModifiedBy\" from \"Provisioning\".\"Tenant\""
	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	var tenants []entities.Tenant
	for rows.Next() {
		tenant := &entities.Tenant{}
		err := rows.Scan(&tenant.TenantId, &tenant.Name, &tenant.TenantCode, &tenant.Deleted, &tenant.CreatedAt, &tenant.CreatedBy, &tenant.ModifiedAt, &tenant.ModifiedBy)
		if err != nil {
			return nil, err
		}
		tenants = append(tenants, *tenant)
	}
	return tenants, err
}
