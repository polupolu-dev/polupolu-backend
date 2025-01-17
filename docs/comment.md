# comment

## コメントの取得 (ニュース記事 ID 指定)

- GetCommentsForNews
- endpoint: `/api/v1/news/{news_id}/comments`
- method: `GET`

### 成功

- HTTP/1.1 200 OK
  - 未確認
  - 正常
  ```bash
  curl http://localhost:8080/api/v1/news/f81e95ef-0320-4810-be74-39af2571312f/comments
  ```
- HTTP/1.1 204 No Content
  - 未確認
  - コメントがない
  ```bash
  curl http://localhost:8080/api/v1/news/f81e95ef-0320-4810-be74-39af2571312f/comments
  ```

### 失敗

- HTTP/1.1 400 Bad Request
  - 未確認
  - UUID ではない ID
  ```bash
  curl http://localhost:8080/api/v1/news/not-uuid/comments
  ```
- HTTP/1.1 404 Not Found
  - 未確認
  - ID に対応するニュースがない
  ```bash
  curl http://localhost:8080/api/v1/news/f81e95ef-0320-4810-be74-39af2571312f/comments
  ```
- HTTP/1.1 500 Internal Server Error
  - 未確認
  - データの取得中にエラー

## コメントの取得 (コメント ID 指定)

- GetComment
- endpoint: `/api/v1/comments/{comment_id}`
- method: `GET`

### 成功

- HTTP/1.1 200 OK
  ```bash
  curl http://localhost:8080/api/v1/comments/a1b2c3d4-e5f6-7890-1234-567890abcdef
  ```

### 失敗

- HTTP/1.1 400 Bad Request
  - 未確認
  - UUID ではない ID
  ```bash
  curl http://localhost:8080/api/v1/comments/not-uuid
  ```
- HTTP/1.1 404 Not Found
  - 未確認
  - ID に対応するコメントがない
  ```bash
  curl http://localhost:8080/api/v1/comments/a1b2c3d4-e5f6-7890-1234-567890abcdef
  ```
- HTTP/1.1 500 Internal Server Error
  - 未確認
  - データの取得中にエラー

## ユーザーのコメント取得

- GetUserComments
- endpoint: `/api/v1/users/{user_id}/comments`
- method: `GET`

### 成功

- HTTP/1.1 200 OK
  - 未確認
  ```bash
  curl http://localhost:8080/api/v1/users/12345678-1234-1234-1234-123456789012/comments
  ```
- HTTP/1.1 204 No Content
  - 未確認
  - コメントがない
  ```bash
  curl http://localhost:8080/api/v1/users/12345678-1234-1234-1234-123456789012/comments
  ```

### 失敗

- HTTP/1.1 400 Bad Request
  - 未確認
  - UUID ではない ID
  ```bash
  curl http://localhost:8080/api/v1/users/not-uuid/comments
  ```
- HTTP/1.1 404 Not Found
  - 未確認
  - ID に対応するユーザーがない
  ```bash
  curl http://localhost:8080/api/v1/users/12345678-1234-1234-1234-123456789012/comments
  ```
- HTTP/1.1 500 Internal Server Error
  - 未確認

## コメント作成

- CreateComment
- endpoint: `/api/v1/news/{news_id}/comments`
- method: `POST`

### 成功

- HTTP/1.1 201 Created
  - 未確認
  - レスポンスヘッダの Location で作成されたコメントの URI を返す
  - `Location: /api/v1/comments/f81e95ef-0320-4810-be74-39af2571312f`
  ```bash
    curl -X POST \
    -H "Content-Type: application/json" \
    -d '{ "user_id": "12345678-1234-1234-1234-123456789012", "content": "新しいコメントです" }' \
    http://localhost:8080/api/v1/news/f81e95ef-0320-4810-be74-39af2571312f/comments
  ```

### 失敗

- [TODO] **未実装**（すべて 500 で返すで対応）
- HTTP/1.1 400 Bad Request
  - 未確認
  - UUID ではない ID
  ```bash
    curl -X POST \
    -H "Content-Type: application/json" \
    -d '{ "user_id": "12345678-1234-1234-1234-123456789012", "content": "新しいコメントです" }' \
    http://localhost:8080/api/v1/news/not-uuid/comments
  ```
  - リクエストボディの user_id がない
  ```bash
    curl -X POST \
    -H "Content-Type: application/json" \
    -d '{ "content": "新しいコメントです" }' \
    http://localhost:8080/api/v1/news/f81e95ef-0320-4810-be74-39af2571312f/comments
  ```
  - リクエストボディの content がない
  ```bash
    curl -X POST \
    -H "Content-Type: application/json" \
    -d '{ "user_id": "12345678-1234-1234-1234-123456789012" }' \
    http://localhost:8080/api/v1/news/f81e95ef-0320-4810-be74-39af2571312f/comments
  ```
  - 存在しないユーザー ID
  ```bash
    curl -X POST \
    -H "Content-Type: application/json" \
    -d '{ "user_id": "12345678-1234-1234-1234-123456789012", "content": "新しいコメントです" }' \
    http://localhost:8080/api/v1/news/f81e95ef-0320-4810-be74-39af2571312f/comments
  ```
- HTTP/1.1 404 Not Found
  - 未確認
  - ID に対応するニュースがない
  ```bash
      curl -X POST \
      -H "Content-Type: application/json" \
      -d '{ "user_id": "12345678-1234-1234-1234-123456789012", "content": "新しいコメントです" }' \
      http://localhost:8080/api/v1/news/not-uuid/comments
  ```
- HTTP/1.1 409 Conflict
  - 未確認
  - 同じユーザーが同じニュース記事に複数回コメント
  ```bash
    curl -X POST \
    -H "Content-Type: application/json" \
    -d '{ "user_id": "12345678-1234-1234-1234-123456789012", "content": "新しいコメントです" }' \
    http://localhost:8080/api/v1/news/f81e95ef-0320-4810-be74-39af2571312f/comments
  ```
- HTTP/1.1 500 Internal Server Error
  - 未確認
  - データの作成中にエラー

## 返信コメント作成

- CreateReply
- endpoint: `/api/v1/comments/{comment_id}/replies`
- method: `POST`

### 成功

- HTTP/1.1 201 Created
  - 未確認
  - レスポンスヘッダの Location で作成されたコメントの URI を返す
  - `Location: /api/v1/comments/a1b2c3d4-e5f6-7890-1234-567890abcdef`
  ```bash
    curl -X POST \
    -H "Content-Type: application/json" \
    -d '{ "user_id": "12345678-1234-1234-1234-123456789012", "content": "新しいコメントです" }' \
    http://localhost:8080/api/v1/comments/a1b2c3d4-e5f6-7890-1234-567890abcdef/replies
  ```

### 失敗

- HTTP/1.1 400 Bad Request
  - 未確認
  - UUID ではない ID
  ```bash
    curl -X POST \
    -H "Content-Type: application/json" \
    -d '{ "user_id": "12345678-1234-1234-1234-123456789012", "content": "新しいコメントです" }' \
    http://localhost:8080/api/v1/comments/not-uuid/replies
  ```
  - リクエストボディの user_id がない
  ```bash
    curl -X POST \
    -H "Content-Type: application/json" \
    -d '{ "content": "新しいコメントです" }' \
    http://localhost:8080/api/v1/comments/a1b2c3d4-e5f6-7890-1234-567890abcdef/replies
  ```
  - リクエストボディの content がない
  ```bash
    curl -X POST \
    -H "Content-Type: application/json" \
    -d '{ "user_id": "12345678-1234-1234-1234-123456789012" }' \
    http://localhost:8080/api/v1/comments/a1b2c3d4-e5f6-7890-1234-567890abcdef/replies
  ```
  - 存在しないユーザー ID
  ```bash
    curl -X POST \
    -H "Content-Type: application/json" \
    -d '{ "user_id": "12345678-1234-1234-1234-123456789012", "content": "新しいコメントです" }' \
    http://localhost:8080/api/v1/comments/a1b2c3d4-e5f6-7890-1234-567890abcdef/replies
  ```
- HTTP/1.1 404 Not Found
  - 未確認
  - ID に対応するコメントがない
  ```bash
    curl -X POST \
    -H "Content-Type: application/json" \
    -d '{ "user_id": "12345678-1234-1234-1234-123456789012", "content": "新しいコメントです" }' \
    http://localhost:8080/api/v1/comments/not-uuid/replies
  ```
- HTTP/1.1 500 Internal Server Error
  - 未確認
  - データの作成中にエラー

## コメント削除

- DeleteComment
- endpoint: `/api/v1/comments/{comment_id}`
- method: `DELETE`

### 成功

- HTTP/1.1 204 No Content
  - 未確認
  ```bash
  curl -X DELETE http://localhost:8080/api/v1/comments/a1b2c3d4-e5f6-7890-1234-567890abcdef
  ```

### 失敗

- HTTP/1.1 400 Bad Request
  - 未確認
  - UUID ではない ID
  ```bash
  curl -X DELETE http://localhost:8080/api/v1/comments/not-uuid
  ```
  - ID に対応するコメントがない（削除済み）
  ```bash
  curl -X DELETE http://localhost:8080/api/v1/comments/a1b2c3d4-e5f6-7890-1234-567890abcdef
  ```
- HTTP/1.1 404 Not Found
  - 未確認
  - ID に対応するコメントがない
  ```bash
  curl -X DELETE http://localhost:8080/api/v1/comments/a1b2c3d4-e5f6-7890-1234-567890abcdef
  ```
- HTTP/1.1 409 Conflict
  - 未確認
  - リソースがロックされている
  ```bash
  curl -X DELETE http://localhost:8080/api/v1/comments/a1b2c3d4-e5f6-7890-1234-567890abcdef
  ```
- HTTP/1.1 500 Internal Server Error
  - 未確認
  - データの削除中にエラー

## コメント更新

- UpdateComment
- endpoint: `/api/v1/comments/{comment_id}`
- method: `PUT`

### 成功

- HTTP/1.1 204 No Content
  - 未確認
  - レスポンスヘッダの Location で作成されたコメントの URI を返す
  - `Location: /api/v1/comments/a1b2c3d4-e5f6-7890-1234-567890abcdef`
  ```bash
  curl -X PUT \
  -H "Content-Type: application/json" \
  -d '{ "content": "更新されたコメントです" }' \
  http://localhost:8080/api/v1/comments/a1b2c3d4-e5f6-7890-1234-567890abcdef
  ```
- HTTP/1.1 201 Created （新規作成）
  - 未確認
  - レスポンスヘッダの Location で作成されたコメントの URI を返す
  - `Location: /api/v1/comments/a1b2c3d4-e5f6-7890-1234-567890abcdef`
  ```bash
  curl -X PUT \
  -H "Content-Type: application/json" \
  -d '{ "content": "更新されたコメントです" }' \
  http://localhost:8080/api/v1/comments/a1b2c3d4-e5f6-7890-1234-567890abcdef
  ```

### 失敗

- HTTP/1.1 400 Bad Request
  - 未確認
  - UUID ではない ID
  ```bash
  curl -X PUT \
  -H "Content-Type: application/json" \
  -d '{ "content": "更新されたコメントです" }' \
  http://localhost:8080/api/v1/comments/not-uuid
  ```
  - リクエストボディの content がない
  ```bash
  curl -X PUT \
  -H "Content-Type: application/json" \
  -d '{}' \
  http://localhost:8080/api/v1/comments/a1b2c3d4-e5f6-7890-1234-567890abcdef
  ```
- HTTP/1.1 404 Not Found
  - 未確認
  - ID に対応するコメントがない
  ```bash
  curl -X PUT \
  -H "Content-Type: application/json" \
  -d '{ "content": "更新されたコメントです" }' \
  http://localhost:8080/api/v1/comments/a1b2c3d4-e5f6-7890-1234-567890abcdef
  ```
- HTTP/1.1 500 Internal Server Error
  - 未確認
  - データの更新中にエラー
