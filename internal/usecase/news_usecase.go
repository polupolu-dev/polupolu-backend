package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/polupolu-dev/polupolu-backend/internal/domain/interfaces"
	"github.com/polupolu-dev/polupolu-backend/internal/domain/models"
	"github.com/polupolu-dev/polupolu-backend/utils/consts"
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
func (uc *NewsUsecase) GetNewsDetail(ctx context.Context, newsID uuid.UUID) (*models.News, error) {
	return uc.newsRepo.GetByID(ctx, newsID)
}

// 特定カテゴリのニュース取得 (MVP)
// 仕様: `category` からカテゴリをキーに持つカテゴリ構造体の配列を取 得する
func (uc *NewsUsecase) GetCategoryNews(ctx context.Context, category string) ([]models.News, error) {
	return uc.newsRepo.GetByCategory(ctx, category)
}

// ニュース作成 (MVP)
// 仕様: ニュース構造体からニュースを作成し，作成したニュース構造体を返す
func (uc *NewsUsecase) CreateNews(ctx context.Context, news *models.News) error {
	if news.Summary == "" && uc.llmService != nil {
		summary, err := uc.llmService.GenerateComment(
			ctx, news.Title, consts.PromptSummary)
		if err != nil {
			return err
		}
		news.Summary = summary
	}

	return uc.newsRepo.Create(ctx, news)
}

// すべてのニュース取得
// 仕様: すべてのニュース構造体を配列で取得する
func (uc *NewsUsecase) GetAllNews(ctx context.Context) ([]models.News, error) {
	return uc.newsRepo.GetAll(ctx)
}

// ニュースの削除
// 仕様: `news_id` からニュースを削除する
func (uc *NewsUsecase) DeleteNews(ctx context.Context, newsID uuid.UUID) error {
	return uc.newsRepo.Delete(ctx, newsID)
}

// ニュースの更新
// 仕様: ニュース構造体からニュースを更新し，ニュース構造体を返す
func (uc *NewsUsecase) UpdateNews(ctx context.Context, news *models.News) error {
	return uc.newsRepo.Update(ctx, news)
}
