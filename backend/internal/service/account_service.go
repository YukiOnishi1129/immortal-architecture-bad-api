package service

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	"immortal-architecture-bad-api/backend/internal/db/sqlc"
)

var (
	// ErrAccountNotFound indicates the account could not be located.
	ErrAccountNotFound = errors.New("account not found")
	// ErrInvalidAccountID indicates the provided ID is not a UUID.
	ErrInvalidAccountID = errors.New("invalid account id")
)

// AccountService bundles all account-related operations.
type AccountService struct {
	queries *sqldb.Queries
}

// NewAccountService creates a new service instance.
func NewAccountService(queries *sqldb.Queries) *AccountService {
	return &AccountService{queries: queries}
}

// CreateOrGetAccountInput represents the payload required to fetch or create an account.
type CreateOrGetAccountInput struct {
	Email             string
	FullName          string
	Provider          string
	ProviderAccountID string
	Thumbnail         *string
}

// CreateOrGetAccount upserts an account based on OAuth payload.
func (s *AccountService) CreateOrGetAccount(ctx context.Context, input CreateOrGetAccountInput) (*sqldb.Account, error) {
	firstName, lastName := splitName(input.FullName)

	params := &sqldb.UpsertAccountParams{
		Email:             strings.TrimSpace(input.Email),
		FirstName:         firstName,
		LastName:          lastName,
		Provider:          strings.TrimSpace(input.Provider),
		ProviderAccountID: strings.TrimSpace(input.ProviderAccountID),
		LastLoginAt: pgtype.Timestamptz{
			Time:  time.Now().UTC(),
			Valid: true,
		},
	}

	if input.Thumbnail != nil && strings.TrimSpace(*input.Thumbnail) != "" {
		params.Thumbnail = pgtype.Text{String: strings.TrimSpace(*input.Thumbnail), Valid: true}
	}

	return s.queries.UpsertAccount(ctx, params)
}

// GetAccountByID fetches an account using its UUID string.
func (s *AccountService) GetAccountByID(ctx context.Context, id string) (*sqldb.Account, error) {
	pgID, err := parseUUID(id)
	if err != nil {
		return nil, ErrInvalidAccountID
	}

	account, err := s.queries.GetAccountByID(ctx, pgID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrAccountNotFound
		}
		return nil, err
	}

	return account, nil
}

func parseUUID(value string) (pgtype.UUID, error) {
	var result pgtype.UUID
	value = strings.TrimSpace(value)
	if value == "" {
		return result, ErrInvalidAccountID
	}

	u, err := uuid.Parse(value)
	if err != nil {
		return result, err
	}

	copy(result.Bytes[:], u[:])
	result.Valid = true
	return result, nil
}

func splitName(full string) (string, string) {
	full = strings.TrimSpace(full)
	if full == "" {
		return "Unknown", ""
	}

	parts := strings.Fields(full)
	if len(parts) == 1 {
		return parts[0], ""
	}

	return parts[0], strings.Join(parts[1:], " ")
}
