package interfaces

import (
	"context"

	"github.com/google/uuid"
	"github.com/polupolu-dev/polupolu-backend/internal/domain/models"
)

// ニュースデータソースと対話するメソッドを定義
type NewsRepository interface {
	// 作成
	Create(ctx context.Context, news *models.News) error

	// 探索（読み込み）
	GetAll(ctx context.Context) ([]models.News, error)
	GetByID(ctx context.Context, id uuid.UUID) (*models.News, error)
	GetByCategory(ctx context.Context, category string) ([]models.News, error)

	// 更新
	Update(ctx context.Context, news *models.News) error

	// 削除
	Delete(ctx context.Context, id uuid.UUID) error
}
