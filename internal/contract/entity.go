package contract

import "github.com/google/uuid"

type IRepository[T IEntity] interface {
}

type IEntity interface {
	GetId() uuid.UUID
	GetCreatedBy() uuid.UUID
	GetCreatedAt() int64
	GetModifiedBy() uuid.UUID
	GetModifiedAt() int64
}

type Entity struct {
	Audit
	Id uuid.UUID `json:"id"`
}

func (entity Entity) GetId() uuid.UUID {
	return entity.Id
}

func NewEntity(id uuid.UUID, createdBy uuid.UUID, createdAt int64) Entity {
	return Entity{
		Id:    id,
		Audit: NewAudit(createdBy, createdAt, createdBy, createdAt),
	}
}

func NewAudit(createdBy uuid.UUID, createdAt int64, modifiedBy uuid.UUID, modifiedAt int64) Audit {
	return Audit{
		CreatedBy:  createdBy,
		CreatedAt:  createdAt,
		ModifiedBy: modifiedBy,
		ModifiedAt: modifiedAt,
	}
}
