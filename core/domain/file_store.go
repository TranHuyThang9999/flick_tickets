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
}
