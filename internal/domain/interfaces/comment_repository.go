package interfaces

import (
	"context"

	"github.com/polupolu-dev/polupolu-backend/internal/domain/models"
)

// コメントデータソースと対話するメソッドを定義
type CommentRepository interface {
	// ニュースへのコメント一覧取得 (MVP)
	// 仕様: `news_id` からコメント構造体の配列を取得する
	GetByNewsID(ctx context.Context, newsID string) ([]models.Comment, error)

	// 特定コメント取得 (MVP)
	// 仕様: `comment_id` からコメント構造体を取得する
	GetByCommentID(ctx context.Context, commentID string) (*models.Comment, error)

	// 特定ユーザーのコメント一覧取得 (MVP)
	// 仕様: `user_id` からコメント構造体の配列を取得する
	GetByUserID(ctx context.Context, userID string) ([]models.Comment, error)

	// ニュースへのコメント作成 (MVP)
	// コメントへの返信作成 (MVP)
	// 仕様: コメント構造体からコメントを作成し，コメント構造体を返す
	Create(ctx context.Context, comment models.Comment) (*models.Comment, error)

	// 削除
	// 仕様: `comment_id` からコメントを削除する
	Delete(ctx context.Context, commentID string) error

	// 更新
	// 仕様: コメント構造体からからコメントを更新し，コメント構造体を返す
	Update(ctx context.Context, comment models.Comment) (*models.Comment, error)
}
