package users

import "github.com/gin-gonic/gin"

func RegisterPublicRoutes(r *gin.RouterGroup, h *Handler) {
	r.POST("", h.Create)
	//r.POST("/login", h.Login)
}

func RegisterPrivateRoutes(r *gin.RouterGroup, h *Handler) {
	r.GET("", h.GetAll)
	r.GET("/:id", h.GetByID)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
}
