package repository

import (
	"context"

	"github.com/WahidinAji/golang-database-mysql/entity"
)

type CommentRepository interface {
	//jika butuh transactional, tambahkan tx di parameternya.
	// Insert(tx , ctx context.Context, comment entity.Comment) (entity.Comment, error)
	Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error)
	FindById(ctx context.Context, id int32) (entity.Comment, error)
	FindByAll(ctx context.Context) ([]entity.Comment, error)
}