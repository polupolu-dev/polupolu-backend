package usecase

import (
	"context"

	"github.com/polupolu-dev/polupolu-backend/internal/domain/interfaces"
	"github.com/polupolu-dev/polupolu-backend/internal/domain/models"
)

const (
	prompt = "ニュースに対するコメントを生成してください"
)

type CommentsUsecase struct {
	commentRepo interfaces.CommentRepository
	newsRepo    interfaces.NewsRepository
	llmService  interfaces.LLMService
}

func NewCommentsUsecase(cr interfaces.CommentRepository, nr interfaces.NewsRepository, ls interfaces.LLMService) *CommentsUsecase {
	return &CommentsUsecase{
		commentRepo: cr,
		newsRepo:    nr,
		llmService:  ls,
	}
}

// ニュースへのコメント一覧取得 (MVP)
// 仕様: `news_id` からコメント構造体の配列を取得する
func (uc *CommentsUsecase) GetCommentsForNews(ctx context.Context, newsID string) ([]models.Comment, error) {
	return uc.commentRepo.GetByID(ctx, newsID)
}

// 特定コメント取得 (MVP)
// 仕様: `comment_id` からコメント構造体を取得する
func (uc *CommentsUsecase) GetComment(ctx context.Context, commentID string) (*models.Comment, error) {
	return uc.commentRepo.GetByCommentID(ctx, commentID)
}

// 特定ユーザーのコメント一覧取得 (MVP)
// 仕様: `user_id` からコメント構造体の配列を取得する
func (uc *CommentsUsecase) GetUserComments(ctx context.Context, userID string) ([]models.Comment, error) {
	return uc.commentRepo.GetByID(ctx, userID)
}

// ニュースへのコメント作成 (MVP)
// 仕様: コメント構造体からコメントを作成し，コメント構造体を返す
func (uc *CommentsUsecase) CreateComment(ctx context.Context, comment *models.Comment, newsID string) error {
	if comment.Content != "" {
		return uc.commentRepo.Create(ctx, comment)
	}

	// ニュースの取得
	news, err := uc.newsRepo.GetByID(ctx, newsID)
	if err != nil {
		return err
	}

	// コメントの生成
	content, err := uc.llmService.GenerateComment(ctx, news.Summary, prompt)
	if err != nil {
		return err
	}
	comment.Content = content

	return uc.commentRepo.Create(ctx, comment)
}

// コメントへの返信作成 (MVP)
// 仕様: コメント構造体からコメントを作成し，コメント構造体を返す
func (uc *CommentsUsecase) CreateReply(ctx context.Context, comment *models.Comment) error {
	return uc.commentRepo.Create(ctx, comment)
}

// 削除
// 仕様: `comment_id` からコメントを削除する
func (uc *CommentsUsecase) DeleteComment(ctx context.Context, commentID string) error {
	return uc.commentRepo.Delete(ctx, commentID)
}

// 更新
// 仕様: コメント構造体からからコメントを更新し，コメント構造体を返す
func (uc *CommentsUsecase) UpdateComment(ctx context.Context, comment *models.Comment) error {
	return uc.commentRepo.Update(ctx, comment)
}
