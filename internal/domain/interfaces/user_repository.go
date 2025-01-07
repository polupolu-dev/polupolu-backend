package interfaces

import (
	"github.com/polupolu-dev/polupolu-backend/internal/domain/models"
)

type UserRepository interface {
	// 作成
	Create(user *models.User) (*models.User, error)

	// 探索（読み込み）
	Find(id string) (*models.User, error)

	// 更新
	Update(user *models.User) (*models.User, error)

	// 削除
	Delete(id string) error
}
