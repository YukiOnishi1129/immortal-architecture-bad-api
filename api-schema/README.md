# Mini Notion API Schema

TypeSpecを使用したAPI仕様の定義と、GoおよびTypeScriptの型自動生成を行うプロジェクトです。

## 概要

このリポジトリは、Mini NotionのフロントエンドとバックエンドのAPI契約を定義し、型安全性を保証します。

- **TypeSpec**: API仕様を定義
- **OpenAPI YAML**: TypeSpecから自動生成
- **Go**: OpenAPIからGoの型とクライアントを自動生成
- **TypeScript**: OpenAPIからTypeScriptの型とクライアントを自動生成

## ディレクトリ構成

```
api-schema/
├── typespec/              # TypeSpec定義
│   ├── models/           # データモデル定義
│   │   ├── note.tsp
│   │   ├── template.tsp
│   │   ├── account.tsp
│   │   └── common.tsp
│   ├── routes/           # エンドポイント定義
│   │   ├── notes.tsp
│   │   ├── templates.tsp
│   │   └── accounts.tsp
│   ├── main.tsp          # エントリーポイント
│   └── tspconfig.yaml    # TypeSpec設定
├── generated/            # 自動生成コード
│   ├── openapi.yaml      # 生成されたOpenAPI仕様
│   ├── go/               # Go生成コード
│   └── typescript/       # TypeScript生成コード
├── scripts/              # 生成スクリプト
│   ├── generate.sh
│   ├── generate-openapi.sh
│   ├── generate-go.sh
│   └── generate-ts.sh
├── package.json
└── README.md
```

## セットアップ

### 必要な環境

- Node.js 20+
- npm または pnpm

### インストール

```bash
npm install
```

## 使い方

### 全自動生成

TypeSpecからOpenAPI YAML、Go、TypeScriptのコードを一括生成します。

```bash
npm run generate
```

### 個別生成

#### OpenAPI YAMLの生成

```bash
npm run generate:openapi
```

#### Goコードの生成

```bash
npm run generate:go
```

#### TypeScriptコードの生成

```bash
npm run generate:ts
```

## TypeSpec定義の編集

### モデルの追加・編集

`typespec/models/` 配下にTypeSpecファイルを作成または編集します。

例：
```typespec
// typespec/models/note.tsp
model Note {
  id: string;
  title: string;
  status: NoteStatus;
}
```

### ルート（エンドポイント）の追加・編集

`typespec/routes/` 配下にTypeSpecファイルを作成または編集します。

例：
```typespec
// typespec/routes/notes.tsp
@route("/api/notes")
interface Notes {
  @get list(): Note[];
  @post create(@body note: CreateNoteRequest): Note;
}
```

### 生成コードの更新

TypeSpecを編集したら、以下のコマンドで生成コードを更新します。

```bash
npm run generate
```

## 生成されたコードの使用

### Go側での使用

```go
import "github.com/your-org/mini-notion-api/generated/go/models"

func CreateNote(req models.CreateNoteRequest) models.NoteResponse {
    // ...
}
```

### TypeScript側での使用

```typescript
import { NoteResponse, CreateNoteRequest } from '@mini-notion/api-schema/typescript/models';

async function createNote(req: CreateNoteRequest): Promise<NoteResponse> {
    // ...
}
```

## CI/CD

GitHub Actionsを使用して、TypeSpecの変更時に自動的にコード生成を実行します。

## ライセンス

MIT
