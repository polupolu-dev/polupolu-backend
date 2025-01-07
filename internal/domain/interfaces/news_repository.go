package interfaces

import (
	"context"

	"github.com/polupolu-dev/polupolu-backend/internal/domain/models"
)

// ニュースデータソースと対話するメソッドを定義
type NewsRepository interface {
	// ニュース詳細取得 (MVP)
	// 仕様: `news_id` からニュース構造体を取得
	GetByNewsID(context context.Context, newsID string) (*models.News, error)

	// 特定カテゴリのニュース取得 (MVP)
	// 仕様: `category` からカテゴリをキーに持つカテゴリ構造体の配列を取得する
	GetNewsByCategory(ctx context.Context, category string) ([]models.News, error)

	// ニュース作成 (MVP)
	// 仕様: ニュース構造体からニュースを作成し，作成したニュース構造体を返す
	Create(context context.Context, news models.News) (*models.News, error)

	// すべてのニュース取得
	// 仕様: すべてのニュース構造体を配列で取得する
	GetAllNews(context context.Context) ([]models.News, error)

	// ニュースの削除
	// 仕様: `news_id` からニュースを削除する
	Delete(context context.Context, newsID string) error

	// ニュースの更新
	// 仕様: ニュース構造体からニュースを更新し，ニュース構造体を返す
	Update(context context.Context, news models.News) (*models.News, error)
}
