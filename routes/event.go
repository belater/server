package routes

import (
	"server-service/controller"
	"server-service/repository"
	"server-service/service"

	"github.com/gin-gonic/gin"
)

var (
	EventRepository = repository.NewEventRepository(DB)
	EventService    = service.NewEventService(EventRepository)
	EventController = controller.NewEventController(EventService)
)

func EventRoute(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.GET("/events/all", EventController.GetAllEvent)
		v1.GET("events/:id", EventController.GetEventById)
		v1.POST("/events", EventController.CreateEvent)
		v1.PUT("/events/:id", EventController.UpdateEventById)
		v1.DELETE("/events/:id", EventController.DeleteEventById)
	}
}
