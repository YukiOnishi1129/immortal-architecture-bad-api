// Package transport bridges generated OpenAPI handlers to legacy services.
package transport

import (
	"errors"
	"net/http"
	"strings"
	"time"

	sqldb "immortal-architecture-bad-api/backend/internal/db/sqlc"
	openapi "immortal-architecture-bad-api/backend/internal/generated/openapi"
	"immortal-architecture-bad-api/backend/internal/service"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
)

// Handler wires generated OpenAPI interface to the legacy services.
type Handler struct {
	accountService  *service.AccountService
	templateService *service.TemplateService
	noteService     *service.NoteService
}

var _ openapi.ServerInterface = (*Handler)(nil)

// NewHandler creates an API handler backed by existing services.
func NewHandler(accountSvc *service.AccountService, templateSvc *service.TemplateService, noteSvc *service.NoteService) *Handler {
	return &Handler{
		accountService:  accountSvc,
		templateService: templateSvc,
		noteService:     noteSvc,
	}
}

// AccountsCreateOrGetAccount handles OAuth upsert.
func (h *Handler) AccountsCreateOrGetAccount(ctx echo.Context) error {
	var body openapi.AccountsCreateOrGetAccountJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return respondError(ctx, http.StatusBadRequest, "invalid payload")
	}

	account, err := h.accountService.CreateOrGetAccount(ctx.Request().Context(), service.CreateOrGetAccountInput{
		Email:             body.Email,
		FullName:          body.Name,
		Provider:          body.Provider,
		ProviderAccountID: body.ProviderAccountId,
		Thumbnail:         body.Thumbnail,
	})
	if err != nil {
		return respondError(ctx, http.StatusInternalServerError, "failed to upsert account")
	}
	return ctx.JSON(http.StatusOK, accountToResponse(account))
}

// AccountsGetCurrentAccount returns account by header or query.
func (h *Handler) AccountsGetCurrentAccount(ctx echo.Context) error {
	accountID := ctx.Request().Header.Get("X-Account-ID")
	if accountID == "" {
		accountID = ctx.QueryParam("accountId")
	}
	if strings.TrimSpace(accountID) == "" {
		return respondError(ctx, http.StatusUnauthorized, "missing X-Account-ID header")
	}

	account, err := h.accountService.GetAccountByID(ctx.Request().Context(), accountID)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrAccountNotFound):
			return respondError(ctx, http.StatusNotFound, "account not found")
		case errors.Is(err, service.ErrInvalidAccountID):
			return respondError(ctx, http.StatusBadRequest, "invalid account id")
		default:
			return respondError(ctx, http.StatusInternalServerError, "failed to fetch account")
		}
	}
	return ctx.JSON(http.StatusOK, accountToResponse(account))
}

// AccountsGetAccountById returns account by ID.
// revive:disable-next-line:var-naming // Method name fixed by generated interface.
func (h *Handler) AccountsGetAccountById(ctx echo.Context, accountID string) error {
	if strings.TrimSpace(accountID) == "" {
		return respondError(ctx, http.StatusBadRequest, "account id is required")
	}
	account, err := h.accountService.GetAccountByID(ctx.Request().Context(), accountID)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrAccountNotFound):
			return respondError(ctx, http.StatusNotFound, "account not found")
		case errors.Is(err, service.ErrInvalidAccountID):
			return respondError(ctx, http.StatusBadRequest, "invalid account id")
		default:
			return respondError(ctx, http.StatusInternalServerError, "failed to fetch account")
		}
	}
	return ctx.JSON(http.StatusOK, accountToResponse(account))
}

// NotesListNotes returns note list.
func (h *Handler) NotesListNotes(ctx echo.Context, params openapi.NotesListNotesParams) error {
	filters := service.NoteFilters{}
	if params.Q != nil {
		filters.Query = params.Q
	}
	if params.OwnerId != nil {
		filters.OwnerID = params.OwnerId
	}
	if params.TemplateId != nil {
		filters.TemplateID = params.TemplateId
	}
	if params.Status != nil {
		status := string(*params.Status)
		filters.Status = &status
	}

	notes, err := h.noteService.ListNotes(ctx.Request().Context(), filters)
	if err != nil {
		return respondError(ctx, http.StatusInternalServerError, "failed to list notes")
	}
	return ctx.JSON(http.StatusOK, notes)
}

// NotesCreateNote creates note.
func (h *Handler) NotesCreateNote(ctx echo.Context) error {
	var body openapi.NotesCreateNoteJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return respondError(ctx, http.StatusBadRequest, "invalid payload")
	}

	sections := map[string]string{}
	if body.Sections != nil {
		for _, section := range *body.Sections {
			sections[section.FieldId] = section.Content
		}
	}

	note, err := h.noteService.CreateNote(
		ctx.Request().Context(),
		body.OwnerId.String(),
		body.TemplateId.String(),
		body.Title,
		sections,
	)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrInvalidAccountID), errors.Is(err, service.ErrInvalidTemplateID):
			return respondError(ctx, http.StatusBadRequest, err.Error())
		default:
			return respondError(ctx, http.StatusInternalServerError, "failed to create note")
		}
	}
	return ctx.JSON(http.StatusCreated, note)
}

// NotesDeleteNote deletes note.
func (h *Handler) NotesDeleteNote(ctx echo.Context, noteID string) error {
	if err := h.noteService.DeleteNote(ctx.Request().Context(), noteID); err != nil {
		switch {
		case errors.Is(err, service.ErrNoteNotFound):
			return respondError(ctx, http.StatusNotFound, "note not found")
		case errors.Is(err, service.ErrInvalidNoteID):
			return respondError(ctx, http.StatusBadRequest, "invalid note id")
		default:
			return respondError(ctx, http.StatusInternalServerError, "failed to delete note")
		}
	}
	return ctx.JSON(http.StatusOK, openapi.ModelsSuccessResponse{
		Success: true,
	})
}

// NotesGetNoteById returns note detail.
// revive:disable-next-line:var-naming // Method name fixed by generated interface.
func (h *Handler) NotesGetNoteById(ctx echo.Context, noteID string) error {
	note, err := h.noteService.GetNote(ctx.Request().Context(), noteID)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrNoteNotFound):
			return respondError(ctx, http.StatusNotFound, "note not found")
		case errors.Is(err, service.ErrInvalidNoteID):
			return respondError(ctx, http.StatusBadRequest, "invalid note id")
		default:
			return respondError(ctx, http.StatusInternalServerError, "failed to fetch note")
		}
	}
	return ctx.JSON(http.StatusOK, note)
}

// NotesUpdateNote updates note.
func (h *Handler) NotesUpdateNote(ctx echo.Context, noteID string) error {
	var body openapi.NotesUpdateNoteJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return respondError(ctx, http.StatusBadRequest, "invalid payload")
	}
	if strings.TrimSpace(body.Title) == "" {
		return respondError(ctx, http.StatusBadRequest, "title is required")
	}
	sections := make(map[string]string, len(body.Sections))
	for _, section := range body.Sections {
		sections[section.Id] = section.Content
	}

	note, err := h.noteService.UpdateNote(ctx.Request().Context(), noteID, body.Title, sections)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrNoteNotFound):
			return respondError(ctx, http.StatusNotFound, "note not found")
		case errors.Is(err, service.ErrInvalidNoteID):
			return respondError(ctx, http.StatusBadRequest, "invalid note id")
		default:
			return respondError(ctx, http.StatusInternalServerError, "failed to update note")
		}
	}
	return ctx.JSON(http.StatusOK, note)
}

// NotesPublishNote publishes note.
func (h *Handler) NotesPublishNote(ctx echo.Context, noteID string) error {
	return h.changeNoteStatus(ctx, noteID, "Publish")
}

// NotesUnpublishNote unpublishes note.
func (h *Handler) NotesUnpublishNote(ctx echo.Context, noteID string) error {
	return h.changeNoteStatus(ctx, noteID, "Draft")
}

func (h *Handler) changeNoteStatus(ctx echo.Context, noteID, status string) error {
	note, err := h.noteService.ChangeStatus(ctx.Request().Context(), noteID, status)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrNoteNotFound):
			return respondError(ctx, http.StatusNotFound, "note not found")
		case errors.Is(err, service.ErrInvalidNoteID):
			return respondError(ctx, http.StatusBadRequest, "invalid note id")
		default:
			return respondError(ctx, http.StatusInternalServerError, "failed to change status")
		}
	}
	return ctx.JSON(http.StatusOK, note)
}

// TemplatesListTemplates returns template list.
func (h *Handler) TemplatesListTemplates(ctx echo.Context, params openapi.TemplatesListTemplatesParams) error {
	filters := service.TemplateFilters{}
	if params.OwnerId != nil {
		filters.OwnerID = params.OwnerId
	}
	if params.Q != nil {
		filters.Query = params.Q
	}

	templates, err := h.templateService.ListTemplates(ctx.Request().Context(), filters)
	if err != nil {
		return respondError(ctx, http.StatusInternalServerError, "failed to list templates")
	}
	return ctx.JSON(http.StatusOK, templates)
}

// TemplatesCreateTemplate creates template.
func (h *Handler) TemplatesCreateTemplate(ctx echo.Context) error {
	var body openapi.TemplatesCreateTemplateJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return respondError(ctx, http.StatusBadRequest, "invalid payload")
	}

	fields := make([]sqldb.Field, len(body.Fields))
	for i, field := range body.Fields {
		fields[i] = sqldb.Field{
			Label:      field.Label,
			Order:      field.Order,
			IsRequired: field.IsRequired,
		}
	}

	template, err := h.templateService.CreateTemplate(ctx.Request().Context(), body.OwnerId.String(), body.Name, fields)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrInvalidAccountID):
			return respondError(ctx, http.StatusBadRequest, "invalid owner id")
		default:
			return respondError(ctx, http.StatusInternalServerError, "failed to create template")
		}
	}
	return ctx.JSON(http.StatusCreated, template)
}

// TemplatesDeleteTemplate deletes template.
func (h *Handler) TemplatesDeleteTemplate(ctx echo.Context, templateID string) error {
	if err := h.templateService.DeleteTemplate(ctx.Request().Context(), templateID); err != nil {
		switch {
		case errors.Is(err, service.ErrTemplateNotFound):
			return respondError(ctx, http.StatusNotFound, "template not found")
		case errors.Is(err, service.ErrTemplateInUse):
			return respondError(ctx, http.StatusConflict, "template in use")
		case errors.Is(err, service.ErrInvalidTemplateID):
			return respondError(ctx, http.StatusBadRequest, "invalid template id")
		default:
			return respondError(ctx, http.StatusInternalServerError, "failed to delete template")
		}
	}
	return ctx.JSON(http.StatusOK, openapi.ModelsSuccessResponse{
		Success: true,
	})
}

// TemplatesGetTemplateById returns template detail.
// revive:disable-next-line:var-naming // Method name fixed by generated interface.
func (h *Handler) TemplatesGetTemplateById(ctx echo.Context, templateID string) error {
	template, err := h.templateService.GetTemplate(ctx.Request().Context(), templateID)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrTemplateNotFound):
			return respondError(ctx, http.StatusNotFound, "template not found")
		case errors.Is(err, service.ErrInvalidTemplateID):
			return respondError(ctx, http.StatusBadRequest, "invalid template id")
		default:
			return respondError(ctx, http.StatusInternalServerError, "failed to fetch template")
		}
	}
	return ctx.JSON(http.StatusOK, template)
}

// TemplatesUpdateTemplate updates template name.
func (h *Handler) TemplatesUpdateTemplate(ctx echo.Context, templateID string) error {
	var body openapi.TemplatesUpdateTemplateJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return respondError(ctx, http.StatusBadRequest, "invalid payload")
	}
	if strings.TrimSpace(body.Name) == "" {
		return respondError(ctx, http.StatusBadRequest, "name is required")
	}
	if len(body.Fields) == 0 {
		return respondError(ctx, http.StatusBadRequest, "at least one field is required")
	}

	template, err := h.templateService.UpdateTemplate(ctx.Request().Context(), templateID, body.Name, body.Fields)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrTemplateNotFound):
			return respondError(ctx, http.StatusNotFound, "template not found")
		case errors.Is(err, service.ErrInvalidTemplateID):
			return respondError(ctx, http.StatusBadRequest, "invalid template id")
		case errors.Is(err, service.ErrTemplateInUse):
			return respondError(ctx, http.StatusConflict, "template is used by existing notes")
		default:
			return respondError(ctx, http.StatusInternalServerError, "failed to update template")
		}
	}
	return ctx.JSON(http.StatusOK, template)
}

func respondError(ctx echo.Context, status int, message string) error {
	return ctx.JSON(status, openapi.ModelsErrorResponse{
		Code:    http.StatusText(status),
		Message: message,
	})
}

func accountToResponse(acc *sqldb.Account) openapi.ModelsAccountResponse {
	return openapi.ModelsAccountResponse{
		Id:          uuidToString(acc.ID),
		Email:       acc.Email,
		FirstName:   acc.FirstName,
		LastName:    acc.LastName,
		FullName:    strings.TrimSpace(acc.FirstName + " " + acc.LastName),
		Thumbnail:   textToString(acc.Thumbnail),
		LastLoginAt: timestamptzToTime(acc.LastLoginAt),
		CreatedAt:   timestamptzToTime(acc.CreatedAt),
		UpdatedAt:   timestamptzToTime(acc.UpdatedAt),
	}
}

func uuidToString(id pgtype.UUID) string {
	if !id.Valid {
		return ""
	}
	u, err := uuid.FromBytes(id.Bytes[:])
	if err != nil {
		return ""
	}
	return u.String()
}

func timestamptzToTime(ts pgtype.Timestamptz) time.Time {
	if !ts.Valid {
		return time.Time{}
	}
	return ts.Time.UTC()
}

func textToString(val pgtype.Text) *string {
	if !val.Valid {
		return nil
	}
	s := strings.TrimSpace(val.String)
	if s == "" {
		return nil
	}
	return &s
}
