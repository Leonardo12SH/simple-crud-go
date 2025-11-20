package note_controller

import (
	"net/http"
	"simple-crud-go/app/models"
	"simple-crud-go/app/requests"
	"simple-crud-go/app/responses"
	"simple-crud-go/database"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllNote(ctx *gin.Context) {

	notes := new([]models.Note)

	err := database.DB.Table("notes").Find(&notes).Error

	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"error": "internal server error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"data": notes,
	})
}

func GetById(ctx *gin.Context) {

	id := ctx.Param("id")
	note := new(responses.NoteResponse)
	errDB := database.DB.Table("notes").Where("id = ?", id).Find(&note).Error

	if errDB != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "do not found note with id " + id,
		})
		return
	}

	if note.ID == nil {
		ctx.JSON(500, gin.H{
			"message": "internal server error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "data transmitted.",
		"data":    note,
	})

}

func Store(ctx *gin.Context) {

	noteReq := new(requests.NoteRequest)

	if errReq := ctx.ShouldBind(&noteReq); errReq != nil {
		ctx.JSON(400, gin.H{
			"message": errReq.Error(),
		})
		return
	}

	note := new(models.Note)
	note.Judul = &noteReq.Judul
	note.Notes = &noteReq.Notes

	errDB := database.DB.Table("notes").Create(&note).Error

	if errDB != nil {
		ctx.JSON(500, gin.H{
			"message": "cannot create data.",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "data created successfully.",
		"data":    note,
	})
}

func UpdateById(ctx *gin.Context) {

	id := ctx.Param("id")
	note := new(models.Note)
	noteReq := new(requests.NoteRequest)

	if errReq := ctx.ShouldBind(&noteReq); errReq != nil {
		ctx.JSON(400, gin.H{
			"message": errReq.Error(),
		})
		return
	}

	errDB := database.DB.Table("notes").Where("id = ?", id).Find(&note)

	if errDB == nil {
		ctx.JSON(500, gin.H{
			"message": "internal server error.",
		})
		return
	}

	if note.ID == nil {
		ctx.JSON(404, gin.H{
			"message": "do not found note with id " + id,
		})
		return
	}

	note.Judul = &noteReq.Judul
	note.Notes = &noteReq.Notes

	errUpdate := database.DB.Table("notes").Where("id = ?", id).Updates(&note).Error

	if errUpdate != nil {
		ctx.JSON(500, gin.H{
			"message": "cannot update data.",
		})
		return
	}

	noteResponse := responses.NoteResponse{
		ID:    note.ID,
		Judul: note.Judul,
		Notes: note.Notes,
	}

	ctx.JSON(200, gin.H{
		"message": "data updated successfully.",
		"data":    noteResponse,
	})

}

func DeleteById(ctx *gin.Context) {

	id := ctx.Param("id")
	note := new(models.Note)
	errDB := database.DB.Table("notes").Where("id = ?", id).Find(&note).Error

	if errDB != nil {
		ctx.JSON(500, gin.H{
			"message": "internal server error",
		})
		return
	}

	if note.ID == nil {
		ctx.JSON(404, gin.H{
			"message": "do not found note with id " + id,
		})
		return
	}

	errDelete := database.DB.Table("notes").Where("id = ?", id).Delete(&note).Error

	if errDelete != nil {
		ctx.JSON(500, gin.H{
			"message": "cannot delete note with id " + id,
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "note with id " + id + " deleted successfully.",
	})
}

func GetNotePaginate(ctx *gin.Context) {

	page := ctx.Query("page")
	perPage := ctx.Query("perPage")

	pageInt, _ := strconv.Atoi(page)
	if pageInt < 1 {
		pageInt = 1
	}

	perPageInt, _ := strconv.Atoi(perPage)
	if perPageInt < 1 {
		perPageInt = 10
	}

	notes := new([]models.Note)

	offset := (pageInt - 1) * perPageInt
	err := database.DB.Table("notes").Limit(perPageInt).Offset(offset).Find(&notes).Error

	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"error": "internal server error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"data":     notes,
		"page":     pageInt,
		"per_page": perPageInt,
	})
}
