# アンチパターン実装ウォークスルー

この Go バックエンドは「やってはいけない設計」をわざと詰め込んだ教材です。OpenAPI の I/O だけは守っていますが、それ以外は現場で遭遇しがちな地雷を並べています。バックエンド設計に慣れていない人でも一発で「これはマズい」と感じられるよう、症状と実害をシンプルに並べました。

---

## 0. まず全体像

| どこで壊れているか | 具体的な症状 | すぐ確認できるファイル |
| --- | --- | --- |
| Controller が太りすぎ | 1 メソッドで HTTP 入出力・バリデーション・DB 詰め替えを全部処理 | `internal/controller/*_controller.go` |
| Service が外部依存まみれ | OpenAPI の型と `sqlc` の型、さらには `echo.Context` まで引きずり込む | `internal/service/*.go` |
| トランザクションを素手で書く | どのユースケースでも `Begin/Commit/Rollback` をコピペ | `internal/service/note_service.go` など |
| 条件分岐だらけ | ステータス判定やフィルターがひたすら if/else で散らばる | `internal/service/note_service.go` |
| エラーが全部同じ顔 | `failed to ...` しか返らず原因が分からない | `internal/controller/helpers.go` |
| テストはほぼ無理 | Service をモックだけで検証できず、コメントで「ここまで書きたいのに無理」と嘆いている | `internal/service/template_service_test.go` |

---

## 1. Controller が何でも屋

- 例: `internal/controller/template_controller.go` の `TemplatesCreateTemplate`
- 1 つの HTTP ハンドラの中で JSON パース → 入力チェック → DB 用構造体作成 → Service 呼び出し → レスポンス整形を全部担当。
- コードが 50 行超えは当たり前で、仕様追加のたびに Controller だけで大量修正が発生。
- Echo の `Context` や OpenAPI 型、`sqldb.Field` などの型がごちゃ混ぜで登場。
- **実害**: どこに何のロジックがあるか覚えられない。テストを書くには Echo の Context をモックし、API 用と DB 用のオブジェクトを両方用意する必要があり、誰も着手しなくなる。

---

## 2. Service が Echo と DB にべったり

- 例: `internal/service/template_service.go`
- 関数の定義からして `CreateTemplate(ctx echo.Context, ...)` のように Web フレームワーク依存。将来、別のフレームワークの Fiber/chi などに変えたくても Service から全部書き換えが必要。
- OpenAPI で生成した `openapi.Models*` と、`sqlc` が出す `sqldb.*` をそのまま受け渡し。ドメインの独自型が存在しない。
- **実害**: API スキーマや DB スキーマの変更が一気に全部の層へ波及する。Service の単体テストを書くには Echo のモックと `sqlc` のモックを両方用意しなければならず、テストコードだけで数百行に膨れ上がる。

---

## 3. トランザクションを素手で管理

- 例: `internal/service/note_service.go:118-175`
- ほぼすべてのユースケースで `tx, err := s.pool.Begin(ctx)` から始まり、`Commit` と `Rollback` を手で書いている。
- どこかで `Rollback` を書き忘れるだけで簡単にリークする。`make lint` でも `Rollback` のエラーチェック漏れが大量に指摘される。
- **実害**: 些細な修正でトランザクション制御を壊しやすい。Repository を挟んでいないので、テストでも本物の DB を立てない限り動作を保証できない。

---

## 4. if 分岐だらけのユースケース

- 例: `internal/service/note_service.go` の `ListNotes`, `ChangeStatus`
- `if filters.OwnerID != nil { ... }` のような分岐がメソッド全体に散らばり、ステータスや検索条件が 1 つ増えるだけで複数の if を追加するハメになる。
- **実害**: 「Draft のときだけ特別なチェック」などの要件を追加すると、一体どこを直せばいいか分からない。レビューで読み切るのも困難。

---

## 5. 役に立たないエラーレスポンス

- 例: `internal/controller/helpers.go`
- `respondError` は `Code = HTTP ステータス文字列`, `Message = 固定文言` でしか返せない。
- 404 でも 500 でも `"failed to ..."` としか返さないので、クライアントはログを掘らないと原因が分からない。
- **実害**: エラーハンドリングを改善したくなったとき、全 Controller を横断的に書き換えないといけない。API クライアントからはバグにしか見えない。

---

## 6. テストで詰む様子をそのまま残してある

- `internal/service/template_service_test.go` に、書ききれないテストケースをコメントで列挙している。
- たとえば以下のような正常系/異常系は、本来 Service 層の責務として検証したいが、DB と `sqlc` に直結しているためモックだけでは再現できない。
- コメント例: 「フィールド生成が途中で失敗したらトランザクションがロールバックされるか」「`pool.Begin` が失敗したときの扱い」「ownerId が存在しないときの外部キー制約エラー」など。
- **実害**: ユニットテストでは何も保証できず、結局 DB を立ち上げた統合テストを毎回走らせるしかない。開発速度が大幅に落ちる。

---

## 7. まとめ（この設計がどれだけ辛いか）

| アンチパターン | 目印 | 実際に困ること |
| --- | --- | --- |
| Controller 何でも屋 | `*_controller.go` が巨大 | 仕様変更時に Controller も Service も二度書き。レビューもテストも地獄 |
| Service が Echo/DB 依存 | `CreateTemplate(ctx echo.Context, ...)` | フレームワーク変更・DB 変更で全層を書き換え。テストでモック地獄 |
| 手書きトランザクション | コピペされた `tx.Begin/Commit/Rollback` | Rollback 忘れの温床。Repository がないので差し替え不能 |
| if だらけのビジネスロジック | `note_service.go` の大量 if | 新条件を足すたびに既存ロジックを壊すリスク増大 |
| 使えないエラーレスポンス | `respondError` | クライアントから原因が見えず、デバッグが辛い |
| テスト不能 | `template_service_test.go` のコメント | 「本来書くべきユニットテスト」が全部コメントに追いやられている |

このドキュメントと実際のコードを照らし合わせれば、アンチパターンがどうやって開発体験を壊すのかを簡単に観察できます。「フレームワーク/DB に依存した Service」「Controller が肥大化する構造」など、どれか 1 つでも当てはまるとテストや保守が一気に難しくなることが分かるはずです。
