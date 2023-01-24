package notes

import (
	"net/http"

	"github.com/Johannes-Krabbe/notes-app/api/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) DeleteNote(c *gin.Context) {
    id := c.Param("id")

    var note models.Note

    if result := h.DB.First(&note, id); result.Error != nil {
        c.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

    h.DB.Delete(&note)

    c.Status(http.StatusOK)
}