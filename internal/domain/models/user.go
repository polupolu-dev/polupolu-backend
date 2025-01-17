package models

import (
	"errors"
)

type User struct {
	// 必須
	ID string `json:"id"` // ユーザーID（識別用）

	// オプショナル
	CommentIDs          []string `json:"comment_ids"`          // 付けたコメントIDの文字列配列
	Gender              string   `json:"gender"`               // ユーザーの性別
	AgeGroup            uint32   `json:"age_group"`            // ユーザーの年齢層
	Occupation          string   `json:"occupation"`           // ユーザーの職業
	PoliticalView       string   `json:"political_view"`       // ユーザーの政治的観点
	OpinionTone         string   `json:"opinion_tone"`         // 意見のトーン
	SpeechStyle         string   `json:"speech_style"`         // 話し方
	CommentLength       uint32   `json:"comment_length"`       // 生成予定のコメントの長さ
	BackgroundKnowledge string   `json:"background_knowledge"` // 話題の背景知識
	Emotion             string   `json:"emotion"`              // 感情
}

// string で必須の項目のバリデーション
func (u *User) UserValidate() error {
	if u.ID == "" {
		return errors.New("ID is required")
	}

	return nil
}
