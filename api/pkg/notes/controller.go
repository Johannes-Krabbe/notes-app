package notes

import (
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

type handler struct {
    DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
    h := &handler{
        DB: db,
    }

    routes := r.Group("/notes")
    routes.POST("/", h.AddNote)
    routes.GET("/", h.GetNotes)
    routes.GET("/:id", h.GetNote)
    routes.PUT("/:id", h.UpdateNote)
    routes.DELETE("/:id", h.DeleteNote)
}