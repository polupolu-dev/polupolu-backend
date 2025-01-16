package interfaces

import (
	"context"

	"github.com/polupolu-dev/polupolu-backend/internal/domain/models"
)

type UserRepository interface {
	// 作成
	Create(ctx context.Context, user *models.User) (*models.User, error)

	// 探索（読み込み）
	Find(ctx context.Context, id string) (*models.User, error)

	// 更新
	Update(ctx context.Context, user *models.User) (*models.User, error)

	// 削除
	Delete(ctx context.Context, id string) error
}
