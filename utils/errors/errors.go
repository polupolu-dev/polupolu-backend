package errors

import (
	"errors"
	"fmt"
)

// 汎用エラー
var (
	ErrInvalidRequest        = errors.New("リクエストが無効です。")
	ErrInternalServerError   = errors.New("エラーが発生しました。")
	ErrInvalidRequestPayload = errors.New("リクエストの形式が不正です。")
)

// ユーザー関連エラー
var (
	ErrFailedCreateUser  = errors.New("ユーザーの作成に失敗しました。")
	ErrUserNotFound      = errors.New("ユーザーが見つかりませんでした。")
	ErrUserAlreadyExists = errors.New("既に存在するユーザーです。")
)

// 認証関連エラー
var (
	ErrUnAuthorized = errors.New("認証に失敗しました。")
)

func New(msg string) error {
	return errors.New(msg)
}

func Is(err, target error) bool {
	return errors.Is(err, target)
}

func NewNotFoundError(entity, id string) error {
	return fmt.Errorf("%s (ID: %s) が見つかりませんでした。", entity, id)
}
