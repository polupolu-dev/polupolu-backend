# users

## ユーザーの詳細取得

- GetUser
- endpoint: `/api/v1/users/{user_id}`
- method: `GET`

### 成功

- HTTP/1.1 200 OK
  - 未確認
  ```bash
  curl http://localhost:8080/api/v1/users/ca84473f-d3a2-4b36-b418-ddbdb2964c1b
  ```
  ```json
  {
    "id": "user123",
    "name": "John Doe",
    "email": "john.doe@example.com",
    "created_at": "2023-10-27T10:00:00Z",
    "updated_at": "2023-10-27T10:00:00Z"
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
  - 未確認
  ```bash
  curl -X POST \
    -H "Content-Type: application/json" \
    -d '{ "id": "new_user", "name": "Jane Doe", "email": "jane.doe@example.com" }' \
    http://localhost:8080/api/v1/users
  ```
  - レスポンスヘッダの `Location` で作成されたユーザーの URI を返す
  - `Location: /api/v1/users/new_user`
  ```json
  {
    "id": "new_user",
    "name": "Jane Doe",
    "email": "jane.doe@example.com",
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
      -d '{ "name": "Jane Doe" }' \
      http://localhost:8080/api/v1/users
  ```
- HTTP/1.1 409 Conflict
  - 未確認
  - 同じユーザー ID でユーザーが既に存在する場合
  ```bash
    curl -X POST \
      -H "Content-Type: application/json" \
      -d '{ "id": "user123", "name": "Jane Doe", "email": "jane.doe@example.com" }' \
      http://localhost:8080/api/v1/users
  ```
- HTTP/1.1 500 Internal Server Error
  - 未確認
  - データの作成中にエラー

## ユーザー削除

- DeleteUser
- endpoint: `/api/v1/users/{user_id}`
- method: `DELETE`

### 成功

- HTTP/1.1 204 No Content
  - 未確認
  ```bash
  curl -X DELETE http://localhost:8080/api/v1/users/user123
  ```

### 失敗

- HTTP/1.1 404 Not Found
  - 未確認
  - ID に対応するユーザーが存在しない（削除済み）
  ```bash
  curl -X DELETE http://localhost:8080/api/v1/users/not-exist-user
  ```
- HTTP/1.1 409 Conflict
  - 未確認
  - リソースがロック中
  ```bash
  curl -X DELETE http://localhost:8080/api/v1/users/user123
  ```
- HTTP/1.1 500 Internal Server Error
  - 未確認
  - データの削除中にエラー

## ユーザー更新

- UpdateUser
- endpoint: `/api/v1/users/{user_id}`
- method: `PUT`

### 成功

- HTTP/1.1 204 No Content (更新成功)
  - 未確認
  ```bash
  curl -X PUT \
    -H "Content-Type: application/json" \
    -d '{ "name": "Updated Name", "email": "updated@example.com" }' \
    http://localhost:8080/api/v1/users/user123
  ```
- HTTP/1.1 201 Created (新規で作成成功)
  - 未確認
  ```bash
    curl -X PUT \
      -H "Content-Type: application/json" \
      -d '{ "id": "new_user", "name": "New User", "email": "new@example.com" }' \
      http://localhost:8080/api/v1/users/new_user
  ```
  - レスポンスヘッダの `Location` で作成されたユーザーの URI を返す
  - `Location: /api/v1/users/new_user`
  ```json
  {
    "id": "new_user",
    "name": "New User",
    "email": "new@example.com",
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
      -d '{ "email": "updated@example.com" }' \
      http://localhost:8080/api/v1/users/user123
  ```
- HTTP/1.1 404 Not Found
  - 未確認
  - ID に対応するユーザーが存在しない
  ```bash
    curl -X PUT \
      -H "Content-Type: application/json" \
      -d '{ "name": "更新" }' \
      http://localhost:8080/api/v1/users/not-exist-user
  ```
- HTTP/1.1 409 Conflict
  - 未確認
  - リソースが既にある
  ```bash
    curl -X PUT \
      -H "Content-Type: application/json" \
      -d '{ "name": "更新" }' \
      http://localhost:8080/api/v1/users/user123
  ```
  - リソースがロック中
  ```bash
    curl -X PUT \
      -H "Content-Type: application/json" \
      -d '{ "name": "更新" }' \
      http://localhost:8080/api/v1/users/user123
  ```
- HTTP/1.1 500 Internal Server Error
  - 未確認
  - データの更新中にエラー
