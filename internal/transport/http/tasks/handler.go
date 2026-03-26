package tasks

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	domain "todos-api/internal/domain/tasks"
	usecase "todos-api/internal/usecase/tasks"
)

type Handler struct {
	uc usecase.UseCase
}

func New(uc usecase.UseCase) *Handler {
	return &Handler{uc: uc}
}

func (h *Handler) Create(c *gin.Context) {
	var req createTaskRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	task, err := h.uc.Create(c.Request.Context(), req.Title, req.Description)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, toResponse(task))
}

func (h *Handler) GetAll(c *gin.Context) {
	tasks, err := h.uc.GetAll(c.Request.Context())
	if err != nil {
		handleError(c, err)
		return
	}

	resp := make([]taskResponse, 0, len(tasks))
	for _, t := range tasks {
		resp = append(resp, toResponse(t))
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetByID(c *gin.Context) {
	id := c.Param("id")

	task, err := h.uc.GetByID(c.Request.Context(), id)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, toResponse(task))
}

func (h *Handler) Update(c *gin.Context) {
	id := c.Param("id")

	var req updateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	task, err := h.uc.Update(
		c.Request.Context(),
		id,
		req.Title,
		req.Description,
		domain.Status(req.Status),
	)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, toResponse(task))
}

func (h *Handler) Delete(c *gin.Context) {
	id := c.Param("id")

	if err := h.uc.Delete(c.Request.Context(), id); err != nil {
		handleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func toResponse(t *domain.Task) taskResponse {
	return taskResponse{
		ID:          string(t.TID),
		Title:       t.Title,
		Description: t.Description,
		Status:      int(t.Status),
	}
}

func handleError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, domain.ErrTaskNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})

	case errors.Is(err, domain.ErrTaskAlreadyExists):
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})

	case errors.Is(err, domain.ErrEmptyTitle),
		errors.Is(err, domain.ErrEmptyTID):
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
	}
}
