package service

import (
	"errors"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

// parseUUID converts a string to pgtype.UUID.
func parseUUID(value string) (pgtype.UUID, error) {
	var result pgtype.UUID
	value = strings.TrimSpace(value)
	if value == "" {
		return result, errors.New("empty uuid")
	}

	id, err := uuid.Parse(value)
	if err != nil {
		return result, err
	}

	copy(result.Bytes[:], id[:])
	result.Valid = true
	return result, nil
}
