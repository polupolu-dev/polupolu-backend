package usecase

import (
	"context"

	"github.com/polupolu-dev/polupolu-backend/internal/domain/interfaces"
	"github.com/polupolu-dev/polupolu-backend/internal/domain/models"
)

type NewsUsecase struct {
	newsRepo   interfaces.NewsRepository
	llmService interfaces.LLMService
}

func NewNewsUsecase(nr interfaces.NewsRepository, ls interfaces.LLMService) *NewsUsecase {
	return &NewsUsecase{
		newsRepo:   nr,
		llmService: ls,
	}
}

// ニュース詳細取得 (MVP)
// 仕様: `news_id` からニュース構造体を取得
func (uc *NewsUsecase) GetNewsDetail(ctx context.Context, id string) (*models.News, error) {
	// ewsUseCase) GetNews(ctx context.Context, id string) (*models.News, error) {
	return uc.newsRepo.GetByID(ctx, id)
}

// 特定カテゴリのニュース取得 (MVP)
// 仕様: `category` からカテゴリをキーに持つカテゴリ構造体の配列を取 得する
func (uc *NewsUsecase) GetCategoryNews(ctx context.Context, category string) ([]models.News, error) {
	return uc.newsRepo.GetByCategory(ctx, category)
}

// ニュース作成 (MVP)
// 仕様: ニュース構造体からニュースを作成し，作成したニュース構造体を返す
func (un *NewsUsecase) CreateNews(ctx context.Context, news *models.News) error {
	return un.newsRepo.Create(ctx, news)
}

// すべてのニュース取得
// 仕様: すべてのニュース構造体を配列で取得する
func (un *NewsUsecase) GetAllNews(ctx context.Context) ([]models.News, error) {
	return un.newsRepo.GetAll(ctx)
}

// ニュースの削除
// 仕様: `news_id` からニュースを削除する
func (un *NewsUsecase) DeleteNews(ctx context.Context, id string) error {
	return un.newsRepo.Delete(ctx, id)
}

// ニュースの更新
// 仕様: ニュース構造体からニュースを更新し，ニュース構造体を返す
func (un *NewsUsecase) UpdateNews(ctx context.Context, news *models.News) error {
	return un.newsRepo.Update(ctx, news)
}
