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
func (c *CollectionFileStore) GetAll(ctx context.Context) ([]*domain.FileStorages, error) {
	var files []*domain.FileStorages
	result := c.fs.Find(&files)
	return files, result.Error
}
func (c *CollectionFileStore) DeleteFileByAnyIdObject(ctx context.Context, tx *gorm.DB, anyId int64) error {
	result := tx.Where("ticket_id = ?", anyId).Delete(&domain.FileStorages{}) //
	return result.Error
}

func (c *CollectionFileStore) AddListInformationFileStorages(ctx context.Context, tx *gorm.DB, req []*domain.FileStorages) error {
	result := tx.Create(req)
	return result.Error
}
func (c *CollectionFileStore) DeleteFileByIdNotTransaction(ctx context.Context, fileId int64) error {
	result := c.fs.Where("id=?", fileId).Delete(&domain.FileStorages{}) //
	return result.Error
}
