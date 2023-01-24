package notes

import (
	"net/http"

	"github.com/Johannes-Krabbe/notes-app/api/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetNotes(c *gin.Context) {
    var notes []models.Note

    if result := h.DB.Find(&notes); result.Error != nil {
        c.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

    c.JSON(http.StatusOK, &notes)
}