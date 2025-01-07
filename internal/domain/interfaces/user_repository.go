package interfaces

import (
	"context"

	"github.com/polupolu-dev/polupolu-backend/internal/domain/models"
)

type UserRepository interface {
	// 取得 (MVP)
	// 仕様: `user_id` からユーザー構造体を取得
	Get(context context.Context, userID string) (*models.User, error)

	// 作成 (MVP)
	// 仕様: ユーザー構造体からユーザーを作成し，作成したユーザー構造体を返す
	Create(context context.Context, user models.User) (*models.User, error)

	// 削除
	// 仕様: `user_id` からユーザーを削除する
	Delete(context context.Context, userID string) error

	// 仕様: ユーザー構造体からユーザーを更新し，更新したユーザー構造体を返す
	// 名前: updateUser
	Update(context context.Context, user models.User) (*models.User, error)
}
