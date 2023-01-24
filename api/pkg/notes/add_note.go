package notes

import (
	"net/http"

	"github.com/Johannes-Krabbe/notes-app/api/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type AddNoteRequestBody struct {
    Title       string `json:"title"`
    Content     string `json:"content"`
}

func (h handler) AddNote(c *gin.Context) {
    body := AddNoteRequestBody{}

    // getting request's body
    if err := c.BindJSON(&body); err != nil {
        c.AbortWithError(http.StatusBadRequest, err)
        return
    }

    var note models.Note

    note.Title = body.Title
    note.Content = body.Content

    if result := h.DB.Create(&note); result.Error != nil {
        c.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

    c.JSON(http.StatusCreated, &note)
}