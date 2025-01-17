package models

import (
	"errors"
	"time"
)

type Comment struct {
	// 必須
	ID             string         `json:"id"`              // コメントID（識別用）
	ReplyToID      string         `json:"reply_to_id"`     // 返信先のコメントIDかニュースIDの文字列
	UserID         string         `json:"user_id"`         // ユーザーのID
	Content        string         `json:"content"`         // コメント内容
	CreatedAt      time.Time      `json:"created_at"`      // 生成時刻
	FeedbackScores FeedbackScores `json:"feedback_scores"` // コメントの共感・なるほど・いまいちスコア

	// オプショナル
	RepliedIDs []string `json:"replied_ids"` // 返信として付けられたコメントIDの文字列配列
}

// string で必須の項目のバリデーション
func (c *Comment) CommentValidate() error {
	if c.ID == "" {
		return errors.New("id is required")
	}
	if c.ReplyToID != "" {
		return errors.New("reply_to_id is required")
	}
	if c.UserID != "" {
		return errors.New("user_id is required")
	}
	if c.Content == "" {
		return errors.New("content is required")
	}

	return nil
}
