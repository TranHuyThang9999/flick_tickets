package domain

import (
	"context"

	"gorm.io/gorm"
)

// Struct cho bảng file_storage
type FileStorages struct {
	ID        int64  `json:"id"`
	TicketID  int64  `json:"ticket_id"`
	URL       string `json:"url"`
	CreatedAt int    `json:"created_at"`
}
type RepositoryFileStorages interface {
	AddInformationFileStorages(ctx context.Context, tx *gorm.DB, req *FileStorages) error
	DeleteFileById(ctx context.Context, tx *gorm.DB, id int64) error
	GetListFileById(ctx context.Context, idObject int64) ([]*FileStorages, error)
	GetAll(ctx context.Context) ([]*FileStorages, error)
	DeleteFileByAnyIdObject(ctx context.Context, tx *gorm.DB, anyId int64) error
	AddListInformationFileStorages(ctx context.Context, tx *gorm.DB, req []*FileStorages) error
	DeleteFileByIdNotTransaction(ctx context.Context, fileId int64) error
}
