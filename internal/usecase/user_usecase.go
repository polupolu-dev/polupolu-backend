package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/polupolu-dev/polupolu-backend/internal/domain/interfaces"
	"github.com/polupolu-dev/polupolu-backend/internal/domain/models"
)

type UserUsecase struct {
	userRepo interfaces.UserRepository
}

func NewUserUsecase(ur interfaces.UserRepository) *UserUsecase {
	return &UserUsecase{
		userRepo: ur,
	}
}

// 取得 (MVP)
// 仕様: `user_id` からユーザー構造体を取得
func (uc *UserUsecase) GetUser(ctx context.Context, userID uuid.UUID) (*models.User, error) {
	return uc.userRepo.Get(ctx, userID)
}

// 作成 (MVP)
// 仕様: ユーザー構造体からユーザーを作成し，作成したユーザー構造体を返す
func (uc *UserUsecase) CreateUser(ctx context.Context, user *models.User) error {
	return uc.userRepo.Create(ctx, user)
}

// 削除
// 仕様: `user_id` からユーザーを削除する
func (uc *UserUsecase) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	return uc.userRepo.Delete(ctx, userID)
}

// 仕様: ユーザー構造体からユーザーを更新し，更新したユーザー構造体を返す
// 名前: updateUser
func (uc *UserUsecase) UpdateUser(ctx context.Context, user *models.User) error {
	return uc.userRepo.Update(ctx, user)
}
