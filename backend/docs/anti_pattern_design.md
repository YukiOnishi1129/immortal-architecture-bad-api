# アンチパターンAPI設計方針

Mini Notion の Go API は、**設計教材として「ダメなバックエンド」を体験してもらう**ことが目的です。API の I/O 契約やエラー JSON は TypeSpec で厳密に定義しつつ、内部実装は現場でよくあるスパゲッティ構成に寄せます。ここでは採用するアンチパターンの特徴と使用技術をまとめます。

## 全体像

- API 契約: `typespec/` で定義（OpenAPI/TS/Go 生成）。この層だけは整えておく。
- 実装: `backend/` のみで完結。Next.js からは HTTP 経由で呼び出す。
- フレームワーク / ツール
  - Web: Echo
  - DB アクセス: sqlc（生成コードをそのまま利用）
  - マイグレーション: golang-migrate（今後はここで実行。Drizzle は廃止）
  - 生成コード配置: `backend/internal/generated`
- 目的: 「見た目の API は普通だけど、内部だけ地獄」な構成を意図的に作る。

## 採用するアンチパターン

| # | 具体例 | ありがちな問題 |
|---|--------|----------------|
| 1 | **ドメイン層を持たない**。API 型（TypeSpec生成）と DB 型（sqlc生成）の直結。 | DB/契約変更が全層に波及。ビジネスルールの置き場が無い。 |
| 2 | **Controller は薄く、Service が神オブジェクト化**。 | Validation / ルール / SQL / TX / ロギング / Config などが 1 メソッドに集結。 |
| 3 | **Repository interface を作らず sqlc に直依存**。 | 差し替え不可、テストもモックしづらい。 |
| 4 | **ユースケース単位のメソッドになっていない**。 `PublishNote` などが `UpdateNote` の if で処理される。 | 仕様追加のたびに 1 メソッドが肥大化し、ルールが点在。 |
| 5 | **ビジネスルールが if/switch で散在**。 | 「システムのルール」は Service の if を grep しないと分からない。 |
| 6 | **トランザクション制御を Service ごとに手書き**。 | Begin/Commit/Rollback が重複し、境界も統一されない。 |
| 7 | **エラー整形だけ後付け**。内部エラーは雑に `errors.Is`、最後に API エラーへマッピング。 | 表向きは統一された JSON でも、分類の根拠がコードに埋没。 |
| 8 | **Logger/Config などの依存を Service に直刺し**。 | 挙動が環境変数やフラグに引きずられ、テストも困難。 |

## 想定ディレクトリ例

```
backend/
├─ cmd/api/main.go           # Echo のエントリーポイント
├─ internal/
│  ├─ handler/               # Echo ハンドラー（薄い）
│  ├─ service/               # 神オブジェクト化する層
│  ├─ generated/             # TypeSpec/SQLC 生成コード
│  └─ db/                    # sqlc のクエリ定義
└─ migrations/               # golang-migrate 用
```

## ルールまとめ

1. **API の顔つき(TypeSpec)は崩さない。** URL・Request/Response・エラー JSON はクリーン版と同一。
2. **内部はあえて崩す。** 上記アンチパターン表に従い、現場のスパゲッティ感を再現する。
3. **比較教材として使う。** 別リポジトリでクリーンアーキ版を実装し、メリットを学べるようにする。

このドキュメントを前提に、backend 実装時は「アンチパターンの写経」になるよう注意してください。***
