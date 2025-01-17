package usecase

import (
	"context"

	"github.com/polupolu-dev/polupolu-backend/consts"
	"github.com/polupolu-dev/polupolu-backend/internal/domain/interfaces"
	"github.com/polupolu-dev/polupolu-backend/internal/domain/models"
)

type CommentUsecase struct {
	commentRepo interfaces.CommentRepository
	newsRepo    interfaces.NewsRepository
	userRepo    interfaces.UserRepository
	llmService  interfaces.LLMService
}

func NewCommentUsecase(cr interfaces.CommentRepository, nr interfaces.NewsRepository, ur interfaces.UserRepository, ls interfaces.LLMService) *CommentUsecase {
	return &CommentUsecase{
		commentRepo: cr,
		newsRepo:    nr,
		userRepo:    ur,
		llmService:  ls,
	}
}

// ニュースへのコメント一覧取得 (MVP)
// 仕様: `news_id` からコメント構造体の配列を取得する
func (uc *CommentUsecase) GetCommentsForNews(ctx context.Context, newsID string) ([]models.Comment, error) {
	return uc.commentRepo.GetByID(ctx, newsID)
}

// 特定コメント取得 (MVP)
// 仕様: `comment_id` からコメント構造体を取得する
func (uc *CommentUsecase) GetComment(ctx context.Context, commentID string) (*models.Comment, error) {
	return uc.commentRepo.GetByCommentID(ctx, commentID)
}

// 特定ユーザーのコメント一覧取得 (MVP)
// 仕様: `user_id` からコメント構造体の配列を取得する
func (uc *CommentUsecase) GetUserComments(ctx context.Context, userID string) ([]models.Comment, error) {
	return uc.commentRepo.GetByID(ctx, userID)
}

// ニュースへのコメント作成 (MVP)
// 仕様: コメント構造体からコメントを作成し，コメント構造体を返す
func (uc *CommentUsecase) CreateComment(ctx context.Context, comment *models.Comment) error {
	// ニュースが存在することを検証
	news, err := uc.newsRepo.GetByID(ctx, comment.ReplyToID)
	if err != nil {
		return err
	}

	// コンテンツが空の場合、LLM を使用してコメントを生成します
	if comment.Content == "" {
		content, err := uc.llmService.GenerateComment(ctx, news.Summary, consts.PromptComment)
		if err != nil {
			return err
		}
		comment.Content = content
	}

	return uc.commentRepo.Create(ctx, comment)
}

// コメントへの返信作成 (MVP)
// 仕様: コメント構造体からコメントを作成し，コメント構造体を返す
func (uc *CommentUsecase) CreateReply(ctx context.Context, reply *models.Comment) error {
	// 親コメントが存在することを検証
	parentComment, err := uc.commentRepo.GetByCommentID(ctx, reply.ReplyToID)
	if err != nil {
		return err
	}

	if reply.Content == "" {
		content, err := uc.llmService.GenerateComment(
			ctx, parentComment.Content, consts.PromptComment)
		if err != nil {
			return err
		}
		reply.Content = content
	}

	return uc.commentRepo.Create(ctx, reply)
}

// 削除
// 仕様: `comment_id` からコメントを削除する
func (uc *CommentUsecase) DeleteComment(ctx context.Context, commentID string) error {
	return uc.commentRepo.Delete(ctx, commentID)
}

// 更新
// 仕様: コメント構造体からからコメントを更新し，コメント構造体を返す
func (uc *CommentUsecase) UpdateComment(ctx context.Context, comment *models.Comment) error {
	return uc.commentRepo.Update(ctx, comment)
}
