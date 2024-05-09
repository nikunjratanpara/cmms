package contract

import "github.com/google/uuid"

type Audit struct {
	CreatedAt  int64
	ModifiedAt int64
	CreatedBy  uuid.UUID
	ModifiedBy uuid.UUID
}

func (audit Audit) GetCreatedBy() uuid.UUID {
	return audit.CreatedBy
}

func (audit Audit) GetCreatedAt() int64 {
	return audit.CreatedAt
}

func (audit Audit) GetModifiedBy() uuid.UUID {
	return audit.ModifiedBy
}

func (audit Audit) GetModifiedAt() int64 {
	return audit.ModifiedAt
}
