package service

import (
	"context"
	"errors"
	"math"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	openapi "immortal-architecture-bad-api/backend/internal/generated/openapi"
	sqldb "immortal-architecture-bad-api/backend/internal/db/sqlc"
)

var (
	// ErrTemplateNotFound indicates the template does not exist.
	ErrTemplateNotFound = errors.New("template not found")
	// ErrTemplateInUse indicates the template is referenced by notes.
	ErrTemplateInUse = errors.New("template is in use")
	// ErrInvalidTemplateID invalid UUID.
	ErrInvalidTemplateID = errors.New("invalid template id")
)

// TemplateService handles template CRUD.
type TemplateService struct {
	queries *sqldb.Queries
}

// NewTemplateService creates a new service.
func NewTemplateService(queries *sqldb.Queries) *TemplateService {
	return &TemplateService{queries: queries}
}

// TemplateFilters filters template list query.
type TemplateFilters struct {
	OwnerID *string
	Query   *string
}

// ListTemplates returns templates optionally filtered by owner or search query.
func (s *TemplateService) ListTemplates(ctx context.Context, filters TemplateFilters) ([]*openapi.ModelsTemplateResponse, error) {
	params := &sqldb.ListTemplatesParams{}

	if filters.OwnerID != nil && *filters.OwnerID != "" {
		if id, err := parseUUID(*filters.OwnerID); err == nil {
			params.Column1 = id
		}
	}

	if filters.Query != nil && *filters.Query != "" {
		params.Column2 = *filters.Query
	}

	rows, err := s.queries.ListTemplates(ctx, params)
	if err != nil {
		return nil, err
	}

	responses := make([]*openapi.ModelsTemplateResponse, 0, len(rows))
	for _, tpl := range rows {
		response, err := s.composeTemplateResponse(ctx, tpl.ID, tpl.Name, tpl.OwnerID, tpl.UpdatedAt, tpl.IsUsed)
		if err != nil {
			return nil, err
		}
		responses = append(responses, response)
	}

	return responses, nil
}

// GetTemplate returns a template by ID along with is_used flag.
func (s *TemplateService) GetTemplate(ctx context.Context, id string) (*openapi.ModelsTemplateResponse, error) {
	pgID, err := parseUUID(id)
	if err != nil {
		return nil, ErrInvalidTemplateID
	}

	template, err := s.queries.GetTemplateByID(ctx, pgID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrTemplateNotFound
		}
		return nil, err
	}
	return s.composeTemplateResponse(ctx, template.ID, template.Name, template.OwnerID, template.UpdatedAt, template.IsUsed)
}

// CreateTemplate creates a template with its fields.
func (s *TemplateService) CreateTemplate(ctx context.Context, ownerID string, name string, fields []sqldb.Field) (*openapi.ModelsTemplateResponse, error) {
	pgOwner, err := parseUUID(ownerID)
	if err != nil {
		return nil, ErrInvalidAccountID
	}

	template, err := s.queries.CreateTemplate(ctx, &sqldb.CreateTemplateParams{
		Name:    name,
		OwnerID: pgOwner,
	})
	if err != nil {
		return nil, err
	}

	for idx, field := range fields {
		if idx > math.MaxInt32-1 {
			return nil, errors.New("too many fields")
		}
		order := field.Order
		if order == 0 {
			if idx >= math.MaxInt32 {
				return nil, errors.New("too many fields")
			}
			order = int32(idx + 1) //nolint:gosec // idx is bounded above
		}
		_, err = s.queries.CreateField(ctx, &sqldb.CreateFieldParams{
			TemplateID: template.ID,
			Label:      field.Label,
			Order:      order,
			IsRequired: field.IsRequired,
		})
		if err != nil {
			return nil, err
		}
	}

	return s.composeTemplateResponse(ctx, template.ID, template.Name, template.OwnerID, template.UpdatedAt, false)
}

// UpdateTemplate updates template name only (fields handled separately).
func (s *TemplateService) UpdateTemplate(ctx context.Context, id, name string) (*openapi.ModelsTemplateResponse, error) {
	templateID, err := parseUUID(id)
	if err != nil {
		return nil, ErrInvalidTemplateID
	}

	template, err := s.queries.UpdateTemplate(ctx, &sqldb.UpdateTemplateParams{
		ID:   templateID,
		Name: name,
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrTemplateNotFound
		}
		return nil, err
	}

	isUsed, err := s.queries.CheckTemplateInUse(ctx, templateID)
	if err != nil {
		return nil, err
	}

	return s.composeTemplateResponse(ctx, template.ID, template.Name, template.OwnerID, template.UpdatedAt, isUsed)
}

// DeleteTemplate removes template if not in use.
func (s *TemplateService) DeleteTemplate(ctx context.Context, id string) error {
	templateID, err := parseUUID(id)
	if err != nil {
		return ErrInvalidTemplateID
	}

	inUse, err := s.queries.CheckTemplateInUse(ctx, templateID)
	if err != nil {
		return err
	}
	if inUse {
		return ErrTemplateInUse
	}

	if err := s.queries.DeleteTemplate(ctx, templateID); err != nil {
		return err
	}
	return nil
}

func (s *TemplateService) composeTemplateResponse(ctx context.Context, templateID pgtype.UUID, name string, ownerID pgtype.UUID, updatedAt pgtype.Timestamptz, isUsed bool) (*openapi.ModelsTemplateResponse, error) {
	owner, err := s.queries.GetAccountByID(ctx, ownerID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrAccountNotFound
		}
		return nil, err
	}

	fieldRows, err := s.queries.ListFieldsByTemplate(ctx, templateID)
	if err != nil {
		return nil, err
	}

	fields := make([]openapi.ModelsField, 0, len(fieldRows))
	for _, field := range fieldRows {
		fields = append(fields, openapi.ModelsField{
			Id:         uuidToString(field.ID),
			Label:      field.Label,
			Order:      field.Order,
			IsRequired: field.IsRequired,
		})
	}

	response := &openapi.ModelsTemplateResponse{
		Id:      uuidToString(templateID),
		Name:    name,
		OwnerId: uuidToString(ownerID),
		Owner: openapi.ModelsAccountSummary{
			Id:        uuidToString(owner.ID),
			FirstName: owner.FirstName,
			LastName:  owner.LastName,
			Thumbnail: textToPointer(owner.Thumbnail),
		},
		Fields:    fields,
		UpdatedAt: updatedAt.Time,
		IsUsed:    isUsed,
	}

	return response, nil
}
