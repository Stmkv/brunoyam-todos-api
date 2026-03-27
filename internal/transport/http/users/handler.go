package users

import (
	"errors"
	"net/http"
	domain "todos-api/internal/domain/users"
	usecase "todos-api/internal/usecase/users"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	uc usecase.UseCase
}

func New(uc usecase.UseCase) *Handler {
	return &Handler{uc: uc}
}

// Create godoc
// @Summary Create user
// @Description create new user
// @Tags users
// @Accept json
// @Produce json
// @Param input body createUserRequest true "user"
// @Success 201 {object} userResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users [post]
func (h *Handler) Create(c *gin.Context) {
	var req createUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	user, err := h.uc.Create(c.Request.Context(), req.Name, req.Email, req.Password)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, toResponse(user))
}

// GetAll godoc
// @Summary Get all users
// @Description Get all users
// @Tags users
// @Produce json
// @Success 200 {array} userResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /users [get]
func (h *Handler) GetAll(c *gin.Context) {
	users, err := h.uc.GetAll(c.Request.Context())
	if err != nil {
		handleError(c, err)
		return
	}

	resp := make([]userResponse, 0, len(users))
	for _, t := range users {
		resp = append(resp, toResponse(t))
	}

	c.JSON(http.StatusOK, resp)
}

// GetByID godoc
// @Summary Get user by ID
// @Description Get user by ID
// @Tags users
// @Produce json
// @Param id path string true "Task ID"
// @Success 200 {object} userResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /users/{id} [get]
func (h *Handler) GetByID(c *gin.Context) {
	uid := c.Param("id")

	user, err := h.uc.GetByID(c.Request.Context(), uid)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, toResponse(user))
}

// Update godoc
// @Summary Update user
// @Description Update user
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User UID"
// @Param input body updateUserRequest true "user"
// @Success 200 {object} userResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /users/{id} [put]
func (h *Handler) Update(c *gin.Context) {
	uid := c.Param("id")

	var req updateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	user, err := h.uc.Update(
		c.Request.Context(),
		uid,
		req.Name,
		req.Email,
	)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, toResponse(user))
}

// Delete godoc
// @Summary Delete user
// @Description Delete user
// @Tags users
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} userResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /users/{id} [delete]
func (h *Handler) Delete(c *gin.Context) {
	uid := c.Param("id")

	if err := h.uc.Delete(c.Request.Context(), uid); err != nil {
		handleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func toResponse(t *domain.User) userResponse {
	return userResponse{
		UID:   t.UID,
		Name:  t.Name,
		Email: t.Email,
	}
}

func handleError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, domain.ErrUserNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})

	case errors.Is(err, domain.ErrUserAlreadyExists):
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})

	case errors.Is(err, domain.ErrEmptyEmail),
		errors.Is(err, domain.ErrEmptyUID),
		errors.Is(err, domain.ErrEmptyName),
		errors.Is(err, domain.ErrEmptyPassword):
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
	}
}
