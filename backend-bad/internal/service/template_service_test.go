package service

import (
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"

	sqldb "immortal-architecture-bad-api/backend/internal/db/sqlc"
	openapi "immortal-architecture-bad-api/backend/internal/generated/openapi"
)

func newEchoContext() echo.Context {
	e := echo.New()
	req := httptest.NewRequest("POST", "/", nil)
	rec := httptest.NewRecorder()
	return e.NewContext(req, rec)
}

func TestTemplateService_CreateTemplate_Validation(t *testing.T) {
	t.Parallel()

	// 本当はこういうことをテストしたい:
	//   1. ownerID が不正な形式ならエラー
	//   2. fields の order が 0 なら自動採番される
	//   3. DB挿入が失敗したらロールバックされる
	//   4. 存在しない ownerID なら外部キー違反エラー
	//
	// でも今の設計だと 1 しかテストできない。
	// なぜなら Service が pgxpool/sqlc に直結していて、
	// Repository インターフェースがないからモックを差し込めない。

	tests := []struct {
		name        string
		ownerID     string
		fields      []sqldb.Field
		expectedErr error
	}{
		// これだけがテストできる（UUIDパースは DB 不要）
		{
			name:        "invalid owner id",
			ownerID:     "not-a-uuid",
			expectedErr: ErrInvalidAccountID,
		},

		// ↓ 以下、書きたいけど書けないテストたち ↓

		// order=0 のとき自動採番されることを確認したい
		// → でも CreateField が DB を叩くので動かない
		// {
		// 	name:    "success with auto order",
		// 	ownerID: "11111111-1111-1111-1111-111111111111",
		// 	fields: []sqldb.Field{
		// 		{Label: "Title", Order: 0, IsRequired: true},
		// 		{Label: "Body", Order: 0, IsRequired: false},
		// 	},
		// 	expectedErr: nil,
		// },

		// CreateField が失敗したらロールバックされることを確認したい
		// → でも「CreateField を失敗させる」手段がない
		// {
		// 	name:        "create field fails -> transaction rolls back",
		// 	ownerID:     "11111111-1111-1111-1111-111111111111",
		// 	fields:      []sqldb.Field{{Label: "Title", Order: 1}},
		// 	expectedErr: errors.New("tx rollback should run"),
		// },

		// 存在しない ownerID で外部キー違反になることを確認したい
		// → でも実際に DB がないと FK 違反は起きない
		// {
		// 	name:        "owner violates FK",
		// 	ownerID:     "22222222-2222-2222-2222-222222222222",
		// 	fields:      []sqldb.Field{{Label: "Title", Order: 1}},
		// 	expectedErr: errors.New("foreign key violation"),
		// },
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx := newEchoContext()
			svc := &TemplateService{}

			_, err := svc.CreateTemplate(ctx, tt.ownerID, "Template", tt.fields)
			if tt.expectedErr == nil {
				if err != nil {
					t.Fatalf("expected no error, got %v", err)
				}
				return
			}
			if err == nil || !errors.Is(err, tt.expectedErr) {
				t.Fatalf("expected %v, got %v", tt.expectedErr, err)
			}
		})
	}
}

func TestTemplateService_UpdateTemplate_Validation(t *testing.T) {
	t.Parallel()

	// 本当はこういうことをテストしたい:
	//   1. templateID が不正な形式ならエラー
	//   2. 正常に更新できること
	//   3. 使用中のテンプレートは更新できない（ErrTemplateInUse）
	//   4. fields が空なら事前にエラー
	//
	// でも今の設計だと 1 しかテストできない。
	// 2-4 は全部 DB を叩く処理の後にあるので、モックがないと到達できない。

	tests := []struct {
		name        string
		templateID  string
		fields      []openapi.ModelsUpdateFieldRequest
		expectedErr error
	}{
		// これだけがテストできる（UUIDパースは DB 不要）
		{
			name:        "invalid template id",
			templateID:  "bad-id",
			fields:      nil,
			expectedErr: ErrInvalidTemplateID,
		},

		// ↓ 以下、書きたいけど書けないテストたち ↓

		// 正常に更新できることを確認したい
		// → でも UpdateTemplate が DB を叩くので動かない
		// {
		// 	name:       "success update with fields",
		// 	templateID: "11111111-1111-1111-1111-111111111111",
		// 	fields: []openapi.ModelsUpdateFieldRequest{
		// 		{Id: nil, Label: "Title", Order: 1, IsRequired: true},
		// 	},
		// 	expectedErr: nil,
		// },

		// 使用中のテンプレートが更新拒否されることを確認したい
		// → でも CheckTemplateInUse が DB を叩くので動かない
		// {
		// 	name:       "template in use",
		// 	templateID: "11111111-1111-1111-1111-111111111111",
		// 	fields: []openapi.ModelsUpdateFieldRequest{
		// 		{Id: nil, Label: "Title", Order: 1, IsRequired: true},
		// 	},
		// 	expectedErr: ErrTemplateInUse,
		// },

		// fields が空のときにエラーになることを確認したい
		// → これは DB 前のバリデーションだけど、その前に UpdateTemplate が DB を叩く
		// {
		// 	name:        "empty fields should fail before sync",
		// 	templateID:  "11111111-1111-1111-1111-111111111111",
		// 	fields:      []openapi.ModelsUpdateFieldRequest{},
		// 	expectedErr: errors.New("at least one field is required"),
		// },
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx := newEchoContext()
			svc := &TemplateService{}

			_, err := svc.UpdateTemplate(ctx, tt.templateID, "Template", tt.fields)
			if tt.expectedErr == nil {
				if err != nil {
					t.Fatalf("expected no error, got %v", err)
				}
				return
			}
			if err == nil || !errors.Is(err, tt.expectedErr) {
				t.Fatalf("expected %v, got %v", tt.expectedErr, err)
			}
		})
	}
}

func TestTemplateService_DeleteTemplate_Validation(t *testing.T) {
	t.Parallel()

	// 本当はこういうことをテストしたい:
	//   1. templateID が不正な形式ならエラー
	//   2. 正常に削除できること
	//   3. 使用中のテンプレートは削除できない（ErrTemplateInUse）
	//   4. 存在しないテンプレートなら ErrTemplateNotFound
	//
	// でも今の設計だと 1 しかテストできない。
	// 2-4 は全部 CheckTemplateInUse や DeleteTemplate が DB を叩くので到達できない。

	tests := []struct {
		name        string
		templateID  string
		expectedErr error
	}{
		// これだけがテストできる（UUIDパースは DB 不要）
		{
			name:        "invalid template id",
			templateID:  "bad-id",
			expectedErr: ErrInvalidTemplateID,
		},

		// ↓ 以下、書きたいけど書けないテストたち ↓

		// 正常に削除できることを確認したい
		// → でも DeleteTemplate が DB を叩くので動かない
		// {
		// 	name:        "success deletion",
		// 	templateID:  "11111111-1111-1111-1111-111111111111",
		// 	expectedErr: nil,
		// },

		// 使用中のテンプレートが削除拒否されることを確認したい
		// → でも CheckTemplateInUse が DB を叩くので動かない
		// {
		// 	name:        "template in use blocks deletion",
		// 	templateID:  "11111111-1111-1111-1111-111111111111",
		// 	expectedErr: ErrTemplateInUse,
		// },

		// 存在しないテンプレートで NotFound になることを確認したい
		// → でも実際に DB がないと「存在しない」を再現できない
		// {
		// 	name:        "template not found",
		// 	templateID:  "22222222-2222-2222-2222-222222222222",
		// 	expectedErr: ErrTemplateNotFound,
		// },
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx := newEchoContext()
			svc := &TemplateService{}

			err := svc.DeleteTemplate(ctx, tt.templateID)
			if tt.expectedErr == nil {
				if err != nil {
					t.Fatalf("expected no error, got %v", err)
				}
				return
			}
			if err == nil || !errors.Is(err, tt.expectedErr) {
				t.Fatalf("expected %v, got %v", tt.expectedErr, err)
			}
		})
	}
}
