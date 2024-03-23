package domain

import (
	"context"

	"gorm.io/gorm"
)

type RepositoryTransaction interface {
	BeginTransaction(ctx context.Context) (*gorm.DB, error)
}
