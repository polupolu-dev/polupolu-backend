package interfaces

import (
	"github.com/polupolu-dev/polupolu-backend/internal/domain/models"
)

// ニュースデータソースと対話するメソッドを定義
type NewsRepository interface {
	// 作成
	Create(news models.News) (*models.News, error)

	// 探索（読み込み）
	Find(id string) (*models.News, error)
	FindAll() ([]models.News, error)
	FindList(category string) ([]models.News, error)

	// 更新
	Update(news models.News) (*models.News, error)

	// 削除
	Delete(id string) error
}
