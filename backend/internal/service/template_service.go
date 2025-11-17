package service

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"

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
func (s *TemplateService) ListTemplates(ctx context.Context, filters TemplateFilters) ([]*sqldb.ListTemplatesRow, error) {
	params := &sqldb.ListTemplatesParams{}

	if filters.OwnerID != nil && *filters.OwnerID != "" {
		if id, err := parseUUID(*filters.OwnerID); err == nil {
			params.Column1 = id
		}
	}

	if filters.Query != nil && *filters.Query != "" {
		params.Column2 = *filters.Query
	}

	return s.queries.ListTemplates(ctx, params)
}

// GetTemplate returns a template by ID along with is_used flag.
func (s *TemplateService) GetTemplate(ctx context.Context, id string) (*sqldb.GetTemplateByIDRow, error) {
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
	return template, nil
}

// CreateTemplate creates a template with its fields.
func (s *TemplateService) CreateTemplate(ctx context.Context, ownerID string, name string, fields []sqldb.Field) (*sqldb.Template, error) {
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
		if idx > 1<<30 {
			return template, errors.New("too many fields")
		}
		order := int32(idx + 1) //nolint:gosec // idx は実クエリで自然数範囲
		_, err = s.queries.CreateField(ctx, &sqldb.CreateFieldParams{
			TemplateID: template.ID,
			Label:      field.Label,
			Order:      order,
			IsRequired: field.IsRequired,
		})
		if err != nil {
			return template, err
		}
	}

	return template, nil
}

// UpdateTemplate updates template name only (fields handled separately).
func (s *TemplateService) UpdateTemplate(ctx context.Context, id, name string) (*sqldb.Template, error) {
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

	return template, nil
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
