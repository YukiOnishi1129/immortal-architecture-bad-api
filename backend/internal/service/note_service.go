package service

import (
	"context"
	"errors"
	"strings"

	"github.com/jackc/pgx/v5"

	sqldb "immortal-architecture-bad-api/backend/internal/db/sqlc"
)

var (
	// ErrNoteNotFound indicates note missing.
	ErrNoteNotFound = errors.New("note not found")
	// ErrInvalidNoteID invalid UUID.
	ErrInvalidNoteID = errors.New("invalid note id")
)

// NoteService bundles note logic.
type NoteService struct {
	queries *sqldb.Queries
}

// NewNoteService creates service.
func NewNoteService(q *sqldb.Queries) *NoteService {
	return &NoteService{queries: q}
}

// NoteFilters for listing notes.
type NoteFilters struct {
	Status     *string
	TemplateID *string
	OwnerID    *string
	Query      *string
}

// ListNotes returns note list with filters.
func (s *NoteService) ListNotes(ctx context.Context, filters NoteFilters) ([]*sqldb.ListNotesRow, error) {
	params := &sqldb.ListNotesParams{}

	if filters.Status != nil && *filters.Status != "" {
		params.Column1 = *filters.Status
	}
	if filters.TemplateID != nil && *filters.TemplateID != "" {
		if id, err := parseUUID(*filters.TemplateID); err == nil {
			params.Column2 = id
		}
	}
	if filters.OwnerID != nil && *filters.OwnerID != "" {
		if id, err := parseUUID(*filters.OwnerID); err == nil {
			params.Column3 = id
		}
	}
	if filters.Query != nil && *filters.Query != "" {
		params.Column4 = *filters.Query
	}

	return s.queries.ListNotes(ctx, params)
}

// GetNote returns a note by ID.
func (s *NoteService) GetNote(ctx context.Context, id string) (*sqldb.GetNoteByIDRow, []*sqldb.ListSectionsByNoteRow, error) {
	noteID, err := parseUUID(id)
	if err != nil {
		return nil, nil, ErrInvalidNoteID
	}

	note, err := s.queries.GetNoteByID(ctx, noteID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil, ErrNoteNotFound
		}
		return nil, nil, err
	}

	sections, err := s.queries.ListSectionsByNote(ctx, noteID)
	if err != nil {
		return note, nil, err
	}

	return note, sections, nil
}

// CreateNote creates note and sections.
func (s *NoteService) CreateNote(ctx context.Context, ownerID, templateID, title string, sections map[string]string) (*sqldb.Note, error) {
	ownerUUID, err := parseUUID(ownerID)
	if err != nil {
		return nil, ErrInvalidAccountID
	}
	templateUUID, err := parseUUID(templateID)
	if err != nil {
		return nil, ErrInvalidTemplateID
	}

	note, err := s.queries.CreateNote(ctx, &sqldb.CreateNoteParams{
		Title:      strings.TrimSpace(title),
		TemplateID: templateUUID,
		OwnerID:    ownerUUID,
		Status:     "Draft",
	})
	if err != nil {
		return nil, err
	}

	for fieldID, content := range sections {
		fid, parseErr := parseUUID(fieldID)
		if parseErr != nil {
			return note, parseErr
		}
		if _, err = s.queries.CreateSection(ctx, &sqldb.CreateSectionParams{
			NoteID:  note.ID,
			FieldID: fid,
			Content: content,
		}); err != nil {
			return note, err
		}
	}

	return note, nil
}

// UpdateNote updates title and sections.
func (s *NoteService) UpdateNote(ctx context.Context, id, title string, sections map[string]string) (*sqldb.Note, error) {
	noteID, err := parseUUID(id)
	if err != nil {
		return nil, ErrInvalidNoteID
	}

	note, err := s.queries.UpdateNote(ctx, &sqldb.UpdateNoteParams{
		ID:    noteID,
		Title: strings.TrimSpace(title),
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNoteNotFound
		}
		return nil, err
	}

	for sectionID, content := range sections {
		secUUID, parseErr := parseUUID(sectionID)
		if parseErr != nil {
			return note, parseErr
		}
		if _, err := s.queries.UpdateSectionContent(ctx, &sqldb.UpdateSectionContentParams{
			ID:      secUUID,
			Content: content,
		}); err != nil {
			return note, err
		}
	}

	return note, nil
}

// ChangeStatus toggles note status.
func (s *NoteService) ChangeStatus(ctx context.Context, id, status string) (*sqldb.Note, error) {
	noteID, err := parseUUID(id)
	if err != nil {
		return nil, ErrInvalidNoteID
	}

	if status != "Draft" && status != "Publish" {
		return nil, errors.New("invalid status")
	}

	note, err := s.queries.UpdateNoteStatus(ctx, &sqldb.UpdateNoteStatusParams{
		ID:     noteID,
		Status: status,
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNoteNotFound
		}
		return nil, err
	}
	return note, nil
}

// DeleteNote removes note and sections.
func (s *NoteService) DeleteNote(ctx context.Context, id string) error {
	noteID, err := parseUUID(id)
	if err != nil {
		return ErrInvalidNoteID
	}

	if err := s.queries.DeleteNote(ctx, noteID); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrNoteNotFound
		}
		return err
	}
	return nil
}
