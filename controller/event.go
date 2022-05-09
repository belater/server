package controller

import (
	"errors"
	"server-service/service"
	"server-service/entity"
	"server-service/utils"

	"github.com/gin-gonic/gin"
)

type eventController struct {
	eventService service.EventService
}

func NewEventController(eventService service.EventService) *eventController {
	return &eventController{
		eventService: eventService,
	}
}

func (cs *eventController) GetAllEvent(c *gin.Context) {
	datas, err := cs.eventService.GetAllEvent()
	if err != nil {
		c.JSON(500, utils.ErrorMessages(utils.ErrorInternalServer, err))
		return
	}

	c.JSON(200, datas)
}

func (cs *eventController) GetEventById(c *gin.Context) {
	param := c.Param("id")

	if param == "" {
		c.JSON(400, utils.ErrorMessages(utils.ErrorBadRequest, errors.New("parameter not valid")))
		return
	}

	data, err := cs.eventService.GetEventById(param)
	if err != nil {
		c.JSON(500, utils.ErrorMessages(utils.ErrorInternalServer, err))
		return
	}

	c.JSON(200, data)
}

func (cs *eventController) CreateEvent(c *gin.Context) {
	var input entity.EventInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, utils.ErrorMessages(utils.ErrorBadRequest, err))
		return
	}

	err := cs.eventService.CreateEvent(input)
	if err != nil {
		c.JSON(500, utils.ErrorMessages(utils.ErrorInternalServer, err))
		return
	}

	c.JSON(201, gin.H{
		"message": "success create event",
	})
}

func (cs *eventController) UpdateEventById(c *gin.Context) {
	var input entity.EventInput

	param := c.Param("id")
	if param == "" {
		c.JSON(400, utils.ErrorMessages(utils.ErrorBadRequest, errors.New("parameter not valid")))
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, utils.ErrorMessages(utils.ErrorBadRequest, err))
		return
	}

	err := cs.eventService.UpdateEventById(param, input)
	if err != nil {
		c.JSON(500, utils.ErrorMessages(utils.ErrorInternalServer, err))
		return
	}

	c.JSON(200, gin.H{
		"message": "success update event id: " + param,
	})
}

func (cs *eventController) DeleteEventById(c *gin.Context) {
	param := c.Param("id")
	if param == "" {
		c.JSON(400, utils.ErrorMessages(utils.ErrorBadRequest, errors.New("parameter not valid")))
		return
	}

	err := cs.eventService.DeleteEventById(param)
	if err != nil {
		c.JSON(500, utils.ErrorMessages(utils.ErrorInternalServer, err))
		return
	}

	c.JSON(200, gin.H{
		"message": "success delete event id: " + param,
	})
}