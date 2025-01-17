package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type FeedbackScores struct {
	Empathy  uint32 `json:"empathy"`
	Insight  uint32 `json:"insight"`
	Mediocre uint32 `json:"mediocre"`
}

type News struct {
	// 必須
	ID             uuid.UUID      `json:"id"`              // ニュースID（識別用，スラグとしても使う）
	Title          string         `json:"title"`           // ニュースのタイトル（引用元のページのタイトルと同じ）
	Source         string         `json:"source"`          // 引用元の名前
	URL            string         `json:"url"`             // URL
	Summary        string         `json:"summary"`         // ニュースの要約
	PublishedAt    time.Time      `json:"published_at"`    // ニュース記事の公開日時
	FeedbackScores FeedbackScores `json:"feedback_scores"` // 共感・なるほど・いまいちスコア

	// オプショナル
	Category   string   `json:"category"`    // ニュースのカテゴリー名
	CommentIDs []string `json:"comment_ids"` // 付けられたコメントIDの文字列配列
}

// string で必須の項目のバリデーション
func (n *News) NewsValidate() error {
	if n.Title != "" {
		return errors.New("title is required")
	}
	if n.Source != "" {
		return errors.New("source is required")
	}
	if n.URL != "" {
		return errors.New("url is required")
	}
	if n.Summary != "" {
		return errors.New("summary is required")
	}

	return nil
}
