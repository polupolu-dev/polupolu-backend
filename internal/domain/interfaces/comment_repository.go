package interfaces

import (
	"github.com/polupolu-dev/polupolu-backend/internal/domain/models"
)

// コメントデータソースと対話するメソッドを定義
type CommentRepository interface {
	// 作成
	Create(comment *models.Comment) (*models.Comment, error)

	// 探索（読み込み）
	Find(id string) (*models.Comment, error)
	FindList(id string) ([]models.Comment, error)

	// 更新
	Update(comment *models.Comment) (*models.Comment, error)

	// 削除
	Delete(id string) error
}
