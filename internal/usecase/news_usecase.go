package usecase

import (
	"github.com/polupolu-dev/polupolu-backend/internal/domain/interfaces"
	"github.com/polupolu-dev/polupolu-backend/internal/domain/models"
)

type NewsUsecase struct {
	repo interfaces.NewsRepository
}

func NewNewsUsecase(repo interfaces.NewsRepository) *NewsUsecase {
	return &NewsUsecase{repo: repo}
}

// ニュース詳細取得 (MVP)
// 仕様: `news_id` からニュース構造体を取得
func (n *NewsUsecase) GetNewsDetail(newsID string) (*models.News, error) {
	return n.repo.Find(newsID)
}

// 特定カテゴリのニュース取得 (MVP)
// 仕様: `category` からカテゴリをキーに持つカテゴリ構造体の配列を取 得する
func (n *NewsUsecase) GetCategoryNews(category string) ([]models.News, error) {
	return n.repo.FindList(category)
}

// ニュース作成 (MVP)
// 仕様: ニュース構造体からニュースを作成し，作成したニュース構造体を返す
func (n *NewsUsecase) CreateNews(news *models.News) (*models.News, error) {
	return n.repo.Create(news)
}

// すべてのニュース取得
// 仕様: すべてのニュース構造体を配列で取得する
func (n *NewsUsecase) GetAllNews() ([]models.News, error) {
	return n.repo.FindAll()
}

// ニュースの削除
// 仕様: `news_id` からニュースを削除する
func (n *NewsUsecase) DeleteNews(newsID string) error {
	return n.repo.Delete(newsID)
}

// ニュースの更新
// 仕様: ニュース構造体からニュースを更新し，ニュース構造体を返す
func (n *NewsUsecase) UpdateNews(news *models.News) (*models.News, error) {
	return n.repo.Update(news)
}
