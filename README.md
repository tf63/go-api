## GOでREST, gRPC, GraphQLのAPIを設計する
<!-- ロゴとラベルの色はここから https://simpleicons.org -->
[![CI](https://github.com/tf63/go_api/actions/workflows/go.yml/badge.svg)](https://github.com/tf63/go_api/actions/workflows/go.yml)
![Go](https://img.shields.io/badge/Go-1.19-00ADD8?logo=go)
![Docker](https://img.shields.io/badge/Docker-20.10.23-2496ED?logo=docker)
![Postgres](https://img.shields.io/badge/Postgres-15.2-4169E1?logo=postgresql)

**概要**

家計簿を管理するAPIを作ってみる
- `Expense`: 支払いを管理する (title, price, user_idなどを持つ)
- `User`: ユーザーを管理 (nameなどを持つ)


**Todo**
- [x] OpenAPIでスキーマを作成する https://github.com/tf63/go_api/issues/1
- [x] Repositoryを作成する https://github.com/tf63/go_api/issues/4 https://github.com/tf63/go_api/issues/5
- [x] Entityを作成する https://github.com/tf63/go_api/issues/3
- [x] RESTサーバーを実装する https://github.com/tf63/go_api/issues/2
- [x] RESTサーバーをテストする https://github.com/tf63/go_api/issues/9 https://github.com/tf63/go_api/issues/13
- [x] gqlgenでスキーマを作成する https://github.com/tf63/go_api/issues/11
- [x] GraphQLサーバーを実装する https://github.com/tf63/go_api/issues/10
- [ ] GraphQLサーバーをテストする https://github.com/tf63/go_api/issues/12
- [ ] gRPCサーバーを実装する
- [ ] gRPCサーバーをテストする
- [ ] Add delight to the experience when all tasks are complete :tada:

**動作確認**
- REST https://github.com/tf63/go_api/issues/2
- GraphQL https://github.com/tf63/go_api/issues/10

**技術選定**
| 技術 | 役割 |
| - | - |
| net/http | サーバー (標準ライブラリ) |
| GORM | ORM |
| OpenAPI | REST用ライブラリ |
| oapi-codegen | OpenAPIのコード生成 |
| gqlgen | GraphQLライブラリ |
| testify | Goのテストライブラリ |
| godoc | ドキュメント生成 |
| Postman | APIの動作確認 |
| PostgreSQL | DB |
| pgAdmin | DBの監視 |
| Docker | 開発環境 |
| Github Actions | CI |

### 設計など

**API設計**

```mermaid
---
title: API (ver1)
---
graph TB
    External --> |:8080|RESTHandler
    External --> |:9090|GraphQLHandler
    RESTHandler --> ExpenseRepository
    RESTHandler --> UserRepository
    GraphQLHandler --> ExpenseRepository
    GraphQLHandler --> UserRepository
    ExpenseRepository --> ExpenseEntity
    UserRepository --> UserEntity
    ExpenseRepository --> PostgreSQL
    UserRepository --> PostgreSQL

```

**テーブル設計**

```mermaid
---
title: DB (ver1)
---
erDiagram
    users ||--o{ expenses: ""
    users {
        bigint id "PK (index)"
        string name "ユーザー名, not null"
        timestamp created_at
        timestamp updated_at
        timestamp deleted_at
    }

    expenses {
        bigint id "PK (index)"
        bigint user_id "FK index"
        string title "not null"
        bigint price "default:0"
        timestamp created_at
        timestamp updated_at
        timestamp deleted_at
    }
```

