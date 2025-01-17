# news

## ニュース記事の詳細取得

- GetNewsDetail
- endpoint: `/api/v1/news/{news_id}`
- method: `GET`

### 成功

- HTTP/1.1 200 OK
  - 未確認
  ```bash
  curl http://localhost:8080/api/v1/news/f81e95ef-0320-4810-be74-39af2571312f
  ```

```json
{
  "id": "f81e95ef-0320-4810-be74-39af2571312f",
  "title": "最新のテクノロジーニュース",
  "content": "AIの進化に関する記事です。",
  "category": "テクノロジー",
  "created_at": "2023-10-27T10:00:00Z",
  "updated_at": "2023-10-27T10:00:00Z"
}
```

### 失敗

- HTTP/1.1 400 Bad Request
  - 未確認
  - UUID ではない ID
  ```bash
  curl http://localhost:8080/api/v1/news/not-uuid
  ```
- HTTP/1.1 404 Not Found
  - 未確認
  - ID に対応するニュースがない
  ```bash
  curl http://localhost:8080/api/v1/news/f81e95ef-0320-4810-be74-39af2571312f
  ```
- HTTP/1.1 500 Internal Server Error
  - 未確認
  - データの取得中にエラー

  ## カテゴリ別ニュース記事取得

- GetCategoryNews
- endpoint: `/api/v1/news/categories/{category}`
- method: `GET`

### 成功

- HTTP/1.1 200 OK
  - 未確認
  ```bash
  curl http://localhost:8080/api/v1/news/categories/テクノロジー
  ```
  ```json
  [
    {
      "id": "f81e95ef-0320-4810-be74-39af2571312f",
      "title": "最新のテクノロジーニュース",
      "content": "AIの進化に関する記事です。",
      "category": "テクノロジー",
      "created_at": "2023-10-27T10:00:00Z",
      "updated_at": "2023-10-27T10:00:00Z"
    },
    {
      "id": "0a1b2c3d-e4f5-6789-0123-456789abcdef",
      "title": "別のテクノロジー記事",
      "content": "クラウドコンピューティングに関する記事。",
      "category": "テクノロジー",
      "created_at": "2023-10-27T12:00:00Z",
      "updated_at": "2023-10-27T12:00:00Z"
    }
  ]
  ```

### 失敗

- HTTP/1.1 404 Not Found
  - 未確認
  - 指定されたカテゴリが存在しない
  ```bash
  curl http://localhost:8080/api/v1/news/categories/存在しないカテゴリ
  ```
- HTTP/1.1 500 Internal Server Error
  - 未確認
  - データの取得中にエラー

## ニュース記事作成

- CreateNews
- endpoint: `/api/v1/news`
- method: `POST`

### 成功

- HTTP/1.1 201 Created
  - 未確認
  ```bash
  curl -X POST \
    -H "Content-Type: application/json" \
    -d '{ "title": "新しいニュース", "content": "新しい記事です。", "category": "一般" }' \
    http://localhost:8080/api/v1/news
  ```
  - レスポンスヘッダの `Location` で作成されたニュース記事の URI を返す
  - `Location: /api/v1/news/new_news_id`
  ```json
  {
    "id": "new_news_id",
    "title": "新しいニュース",
    "content": "新しい記事です。",
    "category": "一般",
    "created_at": "2023-10-27T12:00:00Z",
    "updated_at": "2023-10-27T12:00:00Z"
  }
  ```

### 失敗

- HTTP/1.1 400 Bad Request
  - 未確認
  - リクエストボディが不正
  ```bash
    curl -X POST \
      -H "Content-Type: application/json" \
      -d '{ "title": "新しいニュース" }' \
      http://localhost:8080/api/v1/news
  ```
- HTTP/1.1 409 Conflict
  - 未確認
  - 同じタイトルで記事が既に存在する場合
  ```bash
    curl -X POST \
    -H "Content-Type: application/json" \
    -d '{ "title": "既存のニュース", "content": "新しい記事です。", "category": "一般" }' \
    http://localhost:8080/api/v1/news
  ```
- HTTP/1.1 500 Internal Server Error
  - 未確認
  - データの作成中にエラー

## 全てのニュース記事取得

- GetAllNews
- endpoint: `/api/v1/news`
- method: `GET`

### 成功

- HTTP/1.1 200 OK
  - 未確認
  ```bash
  curl http://localhost:8080/api/v1/news
  ```
  ```json
  [
    {
      "id": "f81e95ef-0320-4810-be74-39af2571312f",
      "title": "最新のテクノロジーニュース",
      "content": "AIの進化に関する記事です。",
      "category": "テクノロジー",
      "created_at": "2023-10-27T10:00:00Z",
      "updated_at": "2023-10-27T10:00:00Z"
    },
    {
      "id": "0a1b2c3d-e4f5-6789-0123-456789abcdef",
      "title": "別のテクノロジー記事",
      "content": "クラウドコンピューティングに関する記事。",
      "category": "テクノロジー",
      "created_at": "2023-10-27T12:00:00Z",
      "updated_at": "2023-10-27T12:00:00Z"
    },
    {
      "id": "1b2c3d4e-f567-8901-2345-6789abcdef12",
      "title": "経済ニュース速報",
      "content": "最新の経済指標について。",
      "category": "経済",
      "created_at": "2023-10-27T14:00:00Z",
      "updated_at": "2023-10-27T14:00:00Z"
    }
  ]
  ```
- HTTP/1.1 500 Internal Server Error
  - 未確認
  - データの取得中にエラー

  ## ニュース記事更新

- UpdateNews
- endpoint: `/api/v1/news/{news_id}`
- method: `PUT`

### 成功

- HTTP/1.1 204 No Content (更新成功)
  - 未確認
  ```bash
  curl -X PUT \
    -H "Content-Type: application/json" \
    -d '{ "content": "更新された記事です。" }' \
    http://localhost:8080/api/v1/news/f81e95ef-0320-4810-be74-39af2571312f
  ```
- HTTP/1.1 201 Created (新規で作成成功)
  - 未確認
  ```bash
  curl -X PUT \
    -H "Content-Type: application/json" \
    -d '{ "title": "新規作成記事", "content": "新規作成記事です。", "category": "スポーツ" }' \
    http://localhost:8080/api/v1/news/new_news_id
  ```
  - レスポンスヘッダの `Location` で作成されたニュース記事の URI を返す
  - `Location: /api/v1/news/new_news_id`
  ```json
  {
    "id": "new_news_id",
    "title": "新規作成記事",
    "content": "新規作成記事です。",
    "category": "スポーツ",
    "created_at": "2023-10-27T16:00:00Z",
    "updated_at": "2023-10-27T16:00:00Z"
  }
  ```

### 失敗

- HTTP/1.1 400 Bad Request
  - 未確認
  - リクエストボディが不正
  ```bash
    curl -X PUT \
    -H "Content-Type: application/json" \
    -d '{ "content": 123 }' \
    http://localhost:8080/api/v1/news/f81e95ef-0320-4810-be74-39af2571312f
  ```
- HTTP/1.1 404 Not Found
  - 未確認
  - ID に対応するニュース記事がない
  ```bash
    curl -X PUT \
      -H "Content-Type: application/json" \
      -d '{ "content": "更新" }' \
      http://localhost:8080/api/v1/news/not-exist-news
  ```
- HTTP/1.1 409 Conflict
  - 未確認
  - リソースが既にある
  ```bash
    curl -X PUT \
      -H "Content-Type: application/json" \
      -d '{ "content": "更新" }' \
      http://localhost:8080/api/v1/news/f81e95ef-0320-4810-be74-39af2571312f
  ```
  - リソースがロック中
  ```bash
    curl -X PUT \
      -H "Content-Type: application/json" \
      -d '{ "content": "更新" }' \
      http://localhost:8080/api/v1/news/f81e95ef-0320-4810-be74-39af2571312f
  ```
- HTTP/1.1 500 Internal Server Error
  - 未確認
  - データの更新中にエラー

## ニュース記事削除

- DeleteNews
- endpoint: `/api/v1/news/{news_id}`
- method: `DELETE`

### 成功

- HTTP/1.1 204 No Content
  - 未確認
  ```bash
  curl -X DELETE http://localhost:8080/api/v1/news/f81e95ef-0320-4810-be74-39af2571312f
  ```

### 失敗

- HTTP/1.1 400 Bad Request
  - 未確認
  - UUID ではない ID
  ```bash
  curl -X DELETE http://localhost:8080/api/v1/news/not-uuid
  ```
- HTTP/1.1 404 Not Found
  - 未確認
  - ID に対応するニュース記事がない（削除済み）
  ```bash
  curl -X DELETE http://localhost:8080/api/v1/news/f81e95ef-0320-4810-be74-39af2571312f
  ```
- HTTP/1.1 409 Conflict
  - 未確認
  - リソースがロック中
  ```bash
  curl -X DELETE http://localhost:8080/api/v1/news/f81e95ef-0320-4810-be74-39af2571312f
  ```
- HTTP/1.1 500 Internal Server Error
  - 未確認
  - データの削除中にエラー
