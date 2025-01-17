package usecase

import (
	"context"

	"github.com/polupolu-dev/polupolu-backend/internal/domain/interfaces"
	"github.com/polupolu-dev/polupolu-backend/internal/domain/models"
)

type UsersUsecase struct {
	userRepo interfaces.UserRepository
}

func NewUsersUsecase(ur interfaces.UserRepository) *UsersUsecase {
	return &UsersUsecase{
		userRepo: ur,
	}
}

// 取得 (MVP)
// 仕様: `user_id` からユーザー構造体を取得
func (uc *UsersUsecase) GetUser(ctx context.Context, userID string) (*models.User, error) {
	return uc.userRepo.Get(ctx, userID)
}

// 作成 (MVP)
// 仕様: ユーザー構造体からユーザーを作成し，作成したユーザー構造体を返す
func (uc *UsersUsecase) CreateUser(ctx context.Context, user *models.User) error {
	return uc.userRepo.Create(ctx, user)
}

// 削除
// 仕様: `user_id` からユーザーを削除する
func (uc *UsersUsecase) DeleteUser(ctx context.Context, userID string) error {
	return uc.userRepo.Delete(ctx, userID)
}

// 仕様: ユーザー構造体からユーザーを更新し，更新したユーザー構造体を返す
// 名前: updateUser
func (uc *UsersUsecase) UpdateUser(ctx context.Context, user *models.User) error {
	return uc.userRepo.Update(ctx, user)
}
