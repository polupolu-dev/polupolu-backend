package usecase

import (
	"github.com/polupolu-dev/polupolu-backend/internal/domain/interfaces"
	"github.com/polupolu-dev/polupolu-backend/internal/domain/models"
)

type CommentsUsecase struct {
	repo interfaces.CommentRepository
}

func NewCommentsUsecase(repo interfaces.CommentRepository) *CommentsUsecase {
	return &CommentsUsecase{repo: repo}
}

// ニュースへのコメント一覧取得 (MVP)
// 仕様: `news_id` からコメント構造体の配列を取得する
func (u *CommentsUsecase) GetCommentsForNews(newsID string) ([]models.Comment, error) {
	return u.repo.FindList(newsID)
}

// 特定コメント取得 (MVP)
// 仕様: `comment_id` からコメント構造体を取得する
func (u *CommentsUsecase) GetComment(commentID string) (*models.Comment, error) {
	return u.repo.Find(commentID)
}

// 特定ユーザーのコメント一覧取得 (MVP)
// 仕様: `user_id` からコメント構造体の配列を取得する
func (u *CommentsUsecase) GetUserComments(userID string) ([]models.Comment, error) {
	return u.repo.FindList(userID)
}

// ニュースへのコメント作成 (MVP)
// 仕様: コメント構造体からコメントを作成し，コメント構造体を返す
func (u *CommentsUsecase) CreateComment(comment *models.Comment) (*models.Comment, error) {
	return u.repo.Create(comment)
}

// コメントへの返信作成 (MVP)
// 仕様: コメント構造体からコメントを作成し，コメント構造体を返す
func (u *CommentsUsecase) CreateReply(comment *models.Comment) (*models.Comment, error) {
	return u.repo.Create(comment)
}

// 削除
// 仕様: `comment_id` からコメントを削除する
func (u *CommentsUsecase) DeleteComment(commentID string) error {
	return u.repo.Delete(commentID)
}

// 更新
// 仕様: コメント構造体からからコメントを更新し，コメント構造体を返す
func (u *CommentsUsecase) UpdateComment(comment *models.Comment) (*models.Comment, error) {
	return u.repo.Update(comment)
}
