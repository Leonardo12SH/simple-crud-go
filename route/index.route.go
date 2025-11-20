package route

import (
	"simple-crud-go/app/controllers/note_controller"

	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {
	route := app

	// Note Routes
	route.GET("/note", note_controller.GetAllNote)
	route.GET("/note/paginate", note_controller.GetNotePaginate)
	route.GET("/note/:id", note_controller.GetById)
	route.POST("/note", note_controller.Store)
	route.PATCH("/note/:id", note_controller.UpdateById)
	route.DELETE("/note/:id", note_controller.DeleteById)
}
