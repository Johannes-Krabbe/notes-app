package notes

import (
	"net/http"

	"github.com/Johannes-Krabbe/notes-app/api/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type UpdateNoteRequestBody struct {
    Title       string `json:"title"`
    Content string `json:"content"`
}

func (h handler) UpdateNote(c *gin.Context) {
    id := c.Param("id")
    body := UpdateNoteRequestBody{}

    // getting request's body
    if err := c.BindJSON(&body); err != nil {
        c.AbortWithError(http.StatusBadRequest, err)
        return
    }

    var note models.Note

    if result := h.DB.First(&note, id); result.Error != nil {
        c.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

    note.Title = body.Title
    note.Content = body.Content

    h.DB.Save(&note)

    c.JSON(http.StatusOK, &note)
}