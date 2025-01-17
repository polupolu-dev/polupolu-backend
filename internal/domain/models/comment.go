package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	// 必須
	ID             uuid.UUID      `json:"id"`              // コメントID（識別用）
	ReplyToID      uuid.UUID      `json:"reply_to_id"`     // 返信先のコメントIDかニュースIDの文字列
	UserID         uuid.UUID      `json:"user_id"`         // ユーザーのID
	Content        string         `json:"content"`         // コメント内容
	CreatedAt      time.Time      `json:"created_at"`      // 生成時刻
	FeedbackScores FeedbackScores `json:"feedback_scores"` // コメントの共感・なるほど・いまいちスコア

	// オプショナル
	RepliedIDs []uuid.UUID `json:"replied_ids"` // 返信として付けられたコメントIDの文字列配列
}

// string で必須の項目のバリデーション
func (c *Comment) CommentValidate() error {
	if c.Content == "" {
		return errors.New("content is required")
	}

	return nil
}
