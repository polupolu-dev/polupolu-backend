package usecase

import (
	"github.com/polupolu-dev/polupolu-backend/internal/domain/interfaces"
	"github.com/polupolu-dev/polupolu-backend/internal/domain/models"
)

type UsersUsecase struct {
	repo interfaces.UserRepository
}

func NewUsersUsecase(repo interfaces.UserRepository) *UsersUsecase {
	return &UsersUsecase{repo: repo}
}

// 取得 (MVP)
// 仕様: `user_id` からユーザー構造体を取得
func (u *UsersUsecase) GetUser(userID string) (*models.User, error) {
	return u.repo.Find(userID)
}

// 作成 (MVP)
// 仕様: ユーザー構造体からユーザーを作成し，作成したユーザー構造体を返す
func (u *UsersUsecase) CreateUser(user *models.User) (*models.User, error) {
	return u.repo.Create(user)
}

// 削除
// 仕様: `user_id` からユーザーを削除する
func (u *UsersUsecase) DeleteUser(userID string) error {
	return u.repo.Delete(userID)
}

// 仕様: ユーザー構造体からユーザーを更新し，更新したユーザー構造体を返す
// 名前: updateUser
func (u *UsersUsecase) UpdateUser(user *models.User) (*models.User, error) {
	return u.repo.Update(user)
}
