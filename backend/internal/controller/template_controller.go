// Package controller exposes HTTP handlers for Echo.
package controller

import (
	"errors"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"

	sqldb "immortal-architecture-bad-api/backend/internal/db/sqlc"
	"immortal-architecture-bad-api/backend/internal/service"
)

// TemplateController handles template endpoints.
type TemplateController struct {
	service *service.TemplateService
}

// NewTemplateController creates controller.
func NewTemplateController(service *service.TemplateService) *TemplateController {
	return &TemplateController{service: service}
}

// Register routes under /api/templates.
func (h *TemplateController) Register(router *echo.Group) {
	router.GET("", h.listTemplates)
	router.POST("", h.createTemplate)
	router.GET("/:templateId", h.getTemplate)
	router.PUT("/:templateId", h.updateTemplate)
	router.DELETE("/:templateId", h.deleteTemplate)
}

type templateFieldPayload struct {
	Label      string `json:"label"`
	IsRequired bool   `json:"isRequired"`
	Order      int32  `json:"order"`
}

type templatePayload struct {
	Name    string                 `json:"name"`
	OwnerID string                 `json:"ownerId"`
	Fields  []templateFieldPayload `json:"fields"`
}

func (p *templatePayload) validate() string {
	if strings.TrimSpace(p.Name) == "" {
		return "name is required"
	}
	if strings.TrimSpace(p.OwnerID) == "" {
		return "ownerId is required"
	}
	return ""
}

func (h *TemplateController) listTemplates(c echo.Context) error {
	owner := c.QueryParam("ownerId")
	query := c.QueryParam("q")

	filters := service.TemplateFilters{}
	if owner != "" {
		filters.OwnerID = &owner
	}
	if query != "" {
		filters.Query = &query
	}

	templates, err := h.service.ListTemplates(c.Request().Context(), filters)
	if err != nil {
		return respondError(c, http.StatusInternalServerError, "failed to list templates")
	}

	return c.JSON(http.StatusOK, templates)
}

func (h *TemplateController) getTemplate(c echo.Context) error {
	templateID := c.Param("templateId")
	template, err := h.service.GetTemplate(c.Request().Context(), templateID)
	if err != nil {
		if errors.Is(err, service.ErrTemplateNotFound) {
			return respondError(c, http.StatusNotFound, "template not found")
		}
		if errors.Is(err, service.ErrInvalidTemplateID) {
			return respondError(c, http.StatusBadRequest, "invalid template id")
		}
		return respondError(c, http.StatusInternalServerError, "failed to fetch template")
	}
	return c.JSON(http.StatusOK, template)
}

func (h *TemplateController) createTemplate(c echo.Context) error {
	var payload templatePayload
	if err := c.Bind(&payload); err != nil {
		return respondError(c, http.StatusBadRequest, "invalid payload")
	}
	if msg := payload.validate(); msg != "" {
		return respondError(c, http.StatusBadRequest, msg)
	}

	fields := make([]sqldb.Field, len(payload.Fields))
	for i, f := range payload.Fields {
		fields[i] = sqldb.Field{
			Label:      f.Label,
			Order:      f.Order,
			IsRequired: f.IsRequired,
		}
	}

	template, err := h.service.CreateTemplate(c.Request().Context(), payload.OwnerID, payload.Name, fields)
	if err != nil {
		return respondError(c, http.StatusInternalServerError, "failed to create template")
	}
	return c.JSON(http.StatusCreated, template)
}

func (h *TemplateController) updateTemplate(c echo.Context) error {
	templateID := c.Param("templateId")
	var payload struct {
		Name string `json:"name"`
	}
	if err := c.Bind(&payload); err != nil {
		return respondError(c, http.StatusBadRequest, "invalid payload")
	}
	if strings.TrimSpace(payload.Name) == "" {
		return respondError(c, http.StatusBadRequest, "name is required")
	}

	template, err := h.service.UpdateTemplate(c.Request().Context(), templateID, payload.Name)
	if err != nil {
		if errors.Is(err, service.ErrTemplateNotFound) {
			return respondError(c, http.StatusNotFound, "template not found")
		}
		if errors.Is(err, service.ErrInvalidTemplateID) {
			return respondError(c, http.StatusBadRequest, "invalid template id")
		}
		return respondError(c, http.StatusInternalServerError, "failed to update template")
	}

	return c.JSON(http.StatusOK, template)
}

func (h *TemplateController) deleteTemplate(c echo.Context) error {
	templateID := c.Param("templateId")
	err := h.service.DeleteTemplate(c.Request().Context(), templateID)
	if err != nil {
		if errors.Is(err, service.ErrTemplateNotFound) {
			return respondError(c, http.StatusNotFound, "template not found")
		}
		if errors.Is(err, service.ErrTemplateInUse) {
			return respondError(c, http.StatusConflict, "template in use")
		}
		if errors.Is(err, service.ErrInvalidTemplateID) {
			return respondError(c, http.StatusBadRequest, "invalid template id")
		}
		return respondError(c, http.StatusInternalServerError, "failed to delete template")
	}
	return c.NoContent(http.StatusNoContent)
}
