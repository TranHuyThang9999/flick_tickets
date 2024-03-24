package repository

import (
	"context"
	"flick_tickets/configs"
	"flick_tickets/core/adapter"
	"flick_tickets/core/domain"

	"gorm.io/gorm"
)

type CollectionFileStore struct {
	fs *gorm.DB
}

func NewConllectionFileStore(cf *configs.Configs, fileStore *adapter.PostGresql) domain.RepositoryFileStorages {
	return &CollectionFileStore{
		fs: fileStore.CreateCollection(),
	}
}
func (c *CollectionFileStore) AddInformationFileStorages(ctx context.Context, tx *gorm.DB, req *domain.FileStorages) error {
	result := tx.Create(req)
	return result.Error
}
func (c *CollectionFileStore) DeleteFileById(ctx context.Context, tx *gorm.DB, id int64) error {
	result := tx.Where("id=?", id).Delete(&domain.FileStorages{}) //
	return result.Error
}
func (c *CollectionFileStore) GetListFileById(ctx context.Context, id int64) ([]*domain.FileStorages, error) {
	var files []*domain.FileStorages
	result := c.fs.Where("ticket_id = ?", id).Find(&files)
	return files, result.Error
}