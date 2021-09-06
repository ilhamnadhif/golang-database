package repository

import (
	"belajar-golang-database/entity"
	"context"
)

type CommentRepository interface {
	Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error)
	FindById(ctx context.Context, id int32) (entity.Comment, error)
	FindAll(ctx context.Context) ([]entity.Comment, error)
	UpdateById(ctx context.Context, id int32, comment entity.Comment)(entity.Comment, error)
	DeleteById(ctx context.Context, id int32) (string, error)
}

