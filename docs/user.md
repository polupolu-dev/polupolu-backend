# users

## ユーザーの詳細取得

- GetUser
- endpoint: `/api/v1/users/{user_id}`
- method: `GET`

### 成功

- HTTP/1.1 200 OK
  ```bash
  curl http://localhost:8080/api/v1/users/ca84473f-d3a2-4b36-b418-ddbdb2964c1b
  ```
  ```json
  {
    "id": "ca84473f-d3a2-4b36-b418-ddbdb2964c1b",
    "comment_ids": [],
    "gender": "",
    "age_group": 0,
    "occupation": "",
    "political_view": "",
    "opinion_tone": "",
    "speech_style": "",
    "comment_length": 0,
    "background_knowledge": "",
    "emotion": ""
  }
  ```

### 失敗

- HTTP/1.1 404 Not Found
  - ID に対応するユーザーが存在しない
  ```bash
  curl http://localhost:8080/api/v1/users/not-exist-user
  ```
- HTTP/1.1 500 Internal Server Error
  - 未確認
  - データの取得中にエラー

## ユーザー作成

- CreateUser
- endpoint: `/api/v1/users`
- method: `POST`

### 成功

- HTTP/1.1 201 Created
  - [TODO] レスポンスヘッダの `Location` で作成されたユーザーの URI を返す
  ```bash
  curl -X POST \
    -H "Content-Type: application/json" \
    -d '{ "id": "ca84473f-d3a2-4b36-b418-ddbdb2964c1b", "gender": "男性" }' \
    http://localhost:8080/api/v1/users
  ```
  - `Location: /api/v1/users/new_user`

### 失敗

- HTTP/1.1 400 Bad Request
  - リクエストボディが不正（id がない）
  - [TODO] エラーにする
  ```bash
    curl -X POST \
      -H "Content-Type: application/json" \
      -d '{ "gender": "男性" }' \
      http://localhost:8080/api/v1/users
  ```
  - リクエストボディが不正（大き過ぎる）
  ```bash
    curl -X POST \
      -H "Content-Type: application/json" \
      -d '{ "id": "ca84473f-d3a2-4b36-b418-ddbdb2964c1b", "age_group": "999" }' \
      http://localhost:8080/api/v1/users
  ```
  - UUID でない ID
  ```bash
    curl -X POST \
      -H "Content-Type: application/json" \
      -d '{ "id": "not_uuid", "name": "Jane Doe", "email": "jane.doe@example.com" }' \
      http://localhost:8080/api/v1/users
  ```
- HTTP/1.1 409 Conflict
  - 同じユーザー ID でユーザーが既に存在する場合
  - [TODO] 適切なステータスコードにする（500 -> 409）
  ```bash
    curl -X POST \
      -H "Content-Type: application/json" \
      -d '{ "id": "ca84473f-d3a2-4b36-b418-ddbdb2964c1b" }' \
      http://localhost:8080/api/v1/users
  ```
- HTTP/1.1 500 Internal Server Error
  - データの作成中にエラー

## ユーザー削除

- DeleteUser
- endpoint: `/api/v1/users/{user_id}`
- method: `DELETE`

### 成功

- HTTP/1.1 204 No Content
  - 未確認
  ```bash
  curl -X DELETE http://localhost:8080/api/v1/users/ca84473f-d3a2-4b36-b418-ddbdb2964c1b
  ```

### 失敗

- HTTP/1.1 400 Bad Request
  - ID に対応するユーザーが存在しない（削除済み）
  - [TODO] 404 番へ変更検討
  ```bash
  curl -X DELETE http://localhost:8080/api/v1/users/not-exist-user
  ```
- HTTP/1.1 409 Conflict
  - リソースがロック中
  - [TODO] 未確認
  ```bash
  curl -X DELETE http://localhost:8080/api/v1/users/user123
  ```
- HTTP/1.1 500 Internal Server Error
  - データの削除中にエラー

## ユーザー更新

- UpdateUser
- endpoint: `/api/v1/users/{user_id}`
- method: `PUT`
- [TODO] 更新しない値を更新しない

### 成功

- HTTP/1.1 204 No Content (更新成功)
  ```bash
  curl -X PUT \
    -H "Content-Type: application/json" \
    -d '{ "occupation": "学生", "comment_ids": ["f81e95ef-0320-4810-be74-39af2571312f"] }' \
    http://localhost:8080/api/v1/users/ca84473f-d3a2-4b36-b418-ddbdb2964c1b
  ```
- HTTP/1.1 201 Created (新規で作成成功)
  - [TODO] 中途半端な実装なので，廃止にする（エラーにする）
  ```bash
    curl -X PUT \
      -H "Content-Type: application/json" \
      -d '{ "emotion": "驚き", "speech_style": "優しい口調" }' \
      http://localhost:8080/api/v1/users/066c9108-f5ca-4087-b798-1d96dbd90b23
  ```
  - レスポンスヘッダの `Location` で作成されたユーザーの URI を返す
  - `Location: /api/v1/users/new_user`

### 失敗

- HTTP/1.1 400 Bad Request
  - リクエストボディが不正
  - [TODO] 不要な値の対応？
  ```bash
    curl -X PUT \
      -H "Content-Type: application/json" \
      -d '{ "email": "updated@example.com" }' \
      http://localhost:8080/api/v1/users/ca84473f-d3a2-4b36-b418-ddbdb2964c1b
  ```
- HTTP/1.1 404 Not Found
  - ID に対応するユーザーが存在しない
  - [TODO] 適切なステータスコードにする（204 -> 404）
  ```bash
    curl -X PUT \
      -H "Content-Type: application/json" \
      -d '{ "name": "更新" }' \
      http://localhost:8080/api/v1/users/3802d7af-2214-4da9-811f-77626d1434f8
  ```
- HTTP/1.1 409 Conflict
  - リソースがロック中
  - [TODO] 未確認
  ```bash
    curl -X PUT \
      -H "Content-Type: application/json" \
      -d '{ "name": "更新" }' \
      http://localhost:8080/api/v1/users/ca84473f-d3a2-4b36-b418-ddbdb2964c1b
  ```
- HTTP/1.1 500 Internal Server Error
  - データの更新中にエラー
