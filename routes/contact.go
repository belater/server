package routes

import (
	"server-service/repository"
	"server-service/service"
	"server-service/controller"

	"github.com/gin-gonic/gin"
)

var (
	ContactRepository = repository.NewContactRepository(DB)
	ContactService = service.NewContactService(ContactRepository)
	ContactController = controller.NewContactController(ContactService)
)

func ContactRoute(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.GET("/contacts/all", ContactController.GetAllContact)
		v1.GET("contacts/:id", ContactController.GetContactById)
		v1.POST("/contacts", ContactController.CreateContact)
		v1.PUT("/contacts/:id", ContactController.UpdateContactById)
		v1.DELETE("/contacts/:id", ContactController.DeleteContactById)
	}
}
