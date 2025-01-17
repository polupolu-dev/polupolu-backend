package interfaces

import (
	"context"

	"github.com/google/uuid"
	"github.com/polupolu-dev/polupolu-backend/internal/domain/models"
)

type UserRepository interface {
	// 作成
	Create(ctx context.Context, user *models.User) error

	// 探索（読み込み）
	Get(ctx context.Context, id uuid.UUID) (*models.User, error)

	// 更新
	Update(ctx context.Context, user *models.User) error

	// 削除
	Delete(ctx context.Context, id uuid.UUID) error
}
