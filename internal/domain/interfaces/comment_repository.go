package interfaces

import (
	"context"

	"github.com/polupolu-dev/polupolu-backend/internal/domain/models"
)

// コメントデータソースと対話するメソッドを定義
type CommentRepository interface {
	// 作成
	Create(ctx context.Context, comment *models.Comment) (*models.Comment, error)

	// 探索（読み込み）
	Find(ctx context.Context, id string) (*models.Comment, error)
	FindList(ctx context.Context, id string) ([]models.Comment, error)

	// 更新
	Update(ctx context.Context, comment *models.Comment) (*models.Comment, error)

	// 削除
	Delete(ctx context.Context, id string) error
}
