package auth

import (
	"errors"
	"net/http"

	domain "todos-api/internal/domain/auth"
	usecase "todos-api/internal/usecase/auth"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	uc usecase.UseCase
}

func New(uc usecase.UseCase) *Handler {
	return &Handler{uc: uc}
}

// Login godoc
// @Summary Login user
// @Description authenticate user and return tokens
// @Tags auth
// @Accept json
// @Produce json
// @Param input body loginRequest true "credentials"
// @Success 200 {object} loginResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /auth/login [post]
func (h *Handler) Login(c *gin.Context) {
	var req loginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	access, refresh, err := h.uc.Login(
		c.Request.Context(),
		req.Email,
		req.Password,
	)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, loginResponse{
		AccessToken:  access,
		RefreshToken: refresh,
	})
}

func handleError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, domain.ErrInvalidCredentials):
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})

	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
	}
}
