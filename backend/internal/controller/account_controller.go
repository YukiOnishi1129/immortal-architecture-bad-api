package controller

import (
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"

	"immortal-architecture-bad-api/backend/internal/db/sqlc"
	"immortal-architecture-bad-api/backend/internal/service"
)

// AccountController exposes account-related HTTP endpoints.
type AccountController struct {
	service *service.AccountService
}

// NewAccountController creates a new controller instance.
func NewAccountController(service *service.AccountService) *AccountController {
	return &AccountController{service: service}
}

// Register wires account routes.
func (h *AccountController) Register(router *echo.Group) {
	router.POST("/auth", h.createOrGetAccount)
	router.GET("/me", h.getCurrentAccount)
	router.GET("/:accountId", h.getAccountByID)
}

type createOrGetAccountRequest struct {
	Email             string  `json:"email"`
	Name              string  `json:"name"`
	Provider          string  `json:"provider"`
	ProviderAccountID string  `json:"providerAccountId"`
	Thumbnail         *string `json:"thumbnail"`
}

func (req *createOrGetAccountRequest) validate() string {
	switch {
	case strings.TrimSpace(req.Email) == "":
		return "email is required"
	case strings.TrimSpace(req.Name) == "":
		return "name is required"
	case strings.TrimSpace(req.Provider) == "":
		return "provider is required"
	case strings.TrimSpace(req.ProviderAccountID) == "":
		return "providerAccountId is required"
	default:
		return ""
	}
}

func (h *AccountController) createOrGetAccount(c echo.Context) error {
	var payload createOrGetAccountRequest
	if err := c.Bind(&payload); err != nil {
		return respondError(c, http.StatusBadRequest, "invalid payload")
	}

	if msg := payload.validate(); msg != "" {
		return respondError(c, http.StatusBadRequest, msg)
	}

	account, err := h.service.CreateOrGetAccount(
		c.Request().Context(),
		service.CreateOrGetAccountInput{
			Email:             payload.Email,
			FullName:          payload.Name,
			Provider:          payload.Provider,
			ProviderAccountID: payload.ProviderAccountID,
			Thumbnail:         payload.Thumbnail,
		},
	)
	if err != nil {
		return respondError(c, http.StatusInternalServerError, "failed to upsert account")
	}

	return c.JSON(http.StatusOK, newAccountResponse(account))
}

func (h *AccountController) getAccountByID(c echo.Context) error {
	accountID := c.Param("accountId")
	if strings.TrimSpace(accountID) == "" {
		return respondError(c, http.StatusBadRequest, "account id is required")
	}

	account, err := h.service.GetAccountByID(c.Request().Context(), accountID)
	if err != nil {
		switch err {
		case service.ErrAccountNotFound:
			return respondError(c, http.StatusNotFound, "account not found")
		case service.ErrInvalidAccountID:
			return respondError(c, http.StatusBadRequest, "invalid account id")
		default:
			return respondError(c, http.StatusInternalServerError, "failed to fetch account")
		}
	}

	return c.JSON(http.StatusOK, newAccountResponse(account))
}

func (h *AccountController) getCurrentAccount(c echo.Context) error {
	accountID := c.Request().Header.Get("X-Account-ID")
	if accountID == "" {
		accountID = c.QueryParam("accountId")
	}

	if accountID == "" {
		return respondError(c, http.StatusUnauthorized, "missing X-Account-ID header")
	}

	account, err := h.service.GetAccountByID(c.Request().Context(), accountID)
	if err != nil {
		if err == service.ErrAccountNotFound {
			return respondError(c, http.StatusNotFound, "account not found")
		}
		if err == service.ErrInvalidAccountID {
			return respondError(c, http.StatusBadRequest, "invalid account id")
		}
		return respondError(c, http.StatusInternalServerError, "failed to fetch account")
	}

	return c.JSON(http.StatusOK, newAccountResponse(account))
}

type accountResponse struct {
	ID          string  `json:"id"`
	Email       string  `json:"email"`
	FirstName   string  `json:"firstName"`
	LastName    string  `json:"lastName"`
	FullName    string  `json:"fullName"`
	Thumbnail   *string `json:"thumbnail,omitempty"`
	LastLoginAt *string `json:"lastLoginAt,omitempty"`
	CreatedAt   *string `json:"createdAt,omitempty"`
	UpdatedAt   *string `json:"updatedAt,omitempty"`
}

func newAccountResponse(account *sqldb.Account) accountResponse {
	resp := accountResponse{
		ID:        uuidFromPg(account.ID),
		Email:     account.Email,
		FirstName: account.FirstName,
		LastName:  account.LastName,
		FullName:  strings.TrimSpace(account.FirstName + " " + account.LastName),
	}

	if thumb := textToString(account.Thumbnail); thumb != nil && *thumb != "" {
		resp.Thumbnail = thumb
	}

	if ts := timeToRFC3339(account.LastLoginAt); ts != nil {
		resp.LastLoginAt = ts
	}
	if ts := timeToRFC3339(account.CreatedAt); ts != nil {
		resp.CreatedAt = ts
	}
	if ts := timeToRFC3339(account.UpdatedAt); ts != nil {
		resp.UpdatedAt = ts
	}

	return resp
}

func uuidFromPg(id pgtype.UUID) string {
	if !id.Valid {
		return ""
	}
	u, err := uuid.FromBytes(id.Bytes[:])
	if err != nil {
		return ""
	}
	return u.String()
}

func textToString(val pgtype.Text) *string {
	if !val.Valid {
		return nil
	}
	s := val.String
	return &s
}

func timeToRFC3339(ts pgtype.Timestamptz) *string {
	if !ts.Valid {
		return nil
	}
	formatted := ts.Time.UTC().Format(time.RFC3339)
	return &formatted
}

func respondError(c echo.Context, status int, message string) error {
	return c.JSON(status, map[string]string{
		"code":    http.StatusText(status),
		"message": message,
	})
}
