package infrastructure

import (
	"github.com/nikunjratanpara/core/internal/contract"
)

type IMetadataProvider[T contract.IEntity] interface {
	GetSelectColumns() []string
	GetInsertColumns() []string
	GetInsertQuery(entity T, index int)
	GetUpdatePart(entity T, index int)
	GetUpdateWhereClauses(entity T)
	GetDeleteConditions(entity T)
	GetParameters(entity T, index int32) []map[string]any
}

type MetadataProvider[T contract.IEntity] struct {
	Table            string
	PrimaryKeyFields []string
	Fields           []string
}

func (meatadataProvider MetadataProvider[T]) GetSelectColumns() []string {
	return meatadataProvider.Fields
}

func (metadataProvider MetadataProvider[T]) GetInsertColumns() []string {
	return metadataProvider.Fields
}

func (metadataProvider MetadataProvider[T]) GetInsertQuery(entity T, index int) {

}
