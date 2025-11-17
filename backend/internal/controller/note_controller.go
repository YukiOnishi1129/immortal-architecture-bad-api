// Package controller exposes HTTP handlers for Echo.
package controller

import (
	"errors"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"

	"immortal-architecture-bad-api/backend/internal/service"
)

// NoteController handles note APIs.
type NoteController struct {
	service *service.NoteService
}

// NewNoteController creates controller.
func NewNoteController(service *service.NoteService) *NoteController {
	return &NoteController{service: service}
}

// Register routes for /api/notes.
func (h *NoteController) Register(router *echo.Group) {
	router.GET("", h.listNotes)
	router.POST("", h.createNote)
	router.GET("/:noteId", h.getNote)
	router.PUT("/:noteId", h.updateNote)
	router.POST("/:noteId/publish", h.publishNote)
	router.POST("/:noteId/unpublish", h.unpublishNote)
	router.DELETE("/:noteId", h.deleteNote)
}

func (h *NoteController) listNotes(c echo.Context) error {
	status := c.QueryParam("status")
	templateID := c.QueryParam("templateId")
	ownerID := c.QueryParam("ownerId")
	q := c.QueryParam("q")

	filters := service.NoteFilters{}
	if status != "" {
		filters.Status = &status
	}
	if templateID != "" {
		filters.TemplateID = &templateID
	}
	if ownerID != "" {
		filters.OwnerID = &ownerID
	}
	if q != "" {
		filters.Query = &q
	}

	notes, err := h.service.ListNotes(c.Request().Context(), filters)
	if err != nil {
		return respondError(c, http.StatusInternalServerError, "failed to list notes")
	}
	return c.JSON(http.StatusOK, notes)
}

func (h *NoteController) getNote(c echo.Context) error {
	noteID := c.Param("noteId")
	note, sections, err := h.service.GetNote(c.Request().Context(), noteID)
	if err != nil {
		if errors.Is(err, service.ErrNoteNotFound) {
			return respondError(c, http.StatusNotFound, "note not found")
		}
		if errors.Is(err, service.ErrInvalidNoteID) {
			return respondError(c, http.StatusBadRequest, "invalid note id")
		}
		return respondError(c, http.StatusInternalServerError, "failed to fetch note")
	}
	return c.JSON(http.StatusOK, map[string]any{
		"note":     note,
		"sections": sections,
	})
}

func (h *NoteController) createNote(c echo.Context) error {
	var payload struct {
		Title      string            `json:"title"`
		TemplateID string            `json:"templateId"`
		OwnerID    string            `json:"ownerId"`
		Sections   map[string]string `json:"sections"`
	}
	if err := c.Bind(&payload); err != nil {
		return respondError(c, http.StatusBadRequest, "invalid payload")
	}
	if strings.TrimSpace(payload.Title) == "" {
		return respondError(c, http.StatusBadRequest, "title is required")
	}
	if strings.TrimSpace(payload.TemplateID) == "" || strings.TrimSpace(payload.OwnerID) == "" {
		return respondError(c, http.StatusBadRequest, "templateId and ownerId are required")
	}

	note, err := h.service.CreateNote(c.Request().Context(), payload.OwnerID, payload.TemplateID, payload.Title, payload.Sections)
	if err != nil {
		return respondError(c, http.StatusInternalServerError, "failed to create note")
	}
	return c.JSON(http.StatusCreated, note)
}

func (h *NoteController) updateNote(c echo.Context) error {
	noteID := c.Param("noteId")

	var payload struct {
		Title    string            `json:"title"`
		Sections map[string]string `json:"sections"`
	}
	if err := c.Bind(&payload); err != nil {
		return respondError(c, http.StatusBadRequest, "invalid payload")
	}
	if strings.TrimSpace(payload.Title) == "" {
		return respondError(c, http.StatusBadRequest, "title is required")
	}

	note, err := h.service.UpdateNote(c.Request().Context(), noteID, payload.Title, payload.Sections)
	if err != nil {
		if errors.Is(err, service.ErrNoteNotFound) {
			return respondError(c, http.StatusNotFound, "note not found")
		}
		if errors.Is(err, service.ErrInvalidNoteID) {
			return respondError(c, http.StatusBadRequest, "invalid note id")
		}
		return respondError(c, http.StatusInternalServerError, "failed to update note")
	}

	return c.JSON(http.StatusOK, note)
}

func (h *NoteController) publishNote(c echo.Context) error {
	return h.changeStatus(c, "Publish")
}

func (h *NoteController) unpublishNote(c echo.Context) error {
	return h.changeStatus(c, "Draft")
}

func (h *NoteController) changeStatus(c echo.Context, status string) error {
	noteID := c.Param("noteId")
	note, err := h.service.ChangeStatus(c.Request().Context(), noteID, status)
	if err != nil {
		if errors.Is(err, service.ErrNoteNotFound) {
			return respondError(c, http.StatusNotFound, "note not found")
		}
		if errors.Is(err, service.ErrInvalidNoteID) {
			return respondError(c, http.StatusBadRequest, "invalid note id")
		}
		return respondError(c, http.StatusInternalServerError, "failed to change status")
	}
	return c.JSON(http.StatusOK, note)
}

func (h *NoteController) deleteNote(c echo.Context) error {
	noteID := c.Param("noteId")
	if err := h.service.DeleteNote(c.Request().Context(), noteID); err != nil {
		if errors.Is(err, service.ErrNoteNotFound) {
			return respondError(c, http.StatusNotFound, "note not found")
		}
		if errors.Is(err, service.ErrInvalidNoteID) {
			return respondError(c, http.StatusBadRequest, "invalid note id")
		}
		return respondError(c, http.StatusInternalServerError, "failed to delete note")
	}
	return c.NoContent(http.StatusNoContent)
}
