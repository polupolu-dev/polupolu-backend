package interfaces

import (
	"context"

	"github.com/google/uuid"
	"github.com/polupolu-dev/polupolu-backend/internal/domain/models"
)

// コメントデータソースと対話するメソッドを定義
type CommentRepository interface {
	// 作成
	Create(ctx context.Context, comment *models.Comment) error

	// 探索（読み込み）
	GetByID(ctx context.Context, id uuid.UUID) ([]models.Comment, error)
	GetByCommentID(ctx context.Context, commentID uuid.UUID) (*models.Comment, error)

	// 更新
	Update(ctx context.Context, comment *models.Comment) error

	// 削除
	Delete(ctx context.Context, id uuid.UUID) error
}
