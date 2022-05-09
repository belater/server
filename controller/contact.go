package controller

import (
	"errors"
	"server-service/service"
	"server-service/entity"
	"server-service/utils"

	"github.com/gin-gonic/gin"
)

type contactController struct {
	contactService service.ContactService
}

func NewContactController(contactService service.ContactService) *contactController {
	return &contactController{
		contactService: contactService,
	}
}

func (cs *contactController) GetAllContact(c *gin.Context) {
	datas, err := cs.contactService.GetAllContact()
	if err != nil {
		c.JSON(500, utils.ErrorMessages(utils.ErrorInternalServer, err))
		return
	}

	c.JSON(200, datas)
}

func (cs *contactController) GetContactById(c *gin.Context) {
	param := c.Param("id")

	if param == "" {
		c.JSON(400, utils.ErrorMessages(utils.ErrorBadRequest, errors.New("parameter not valid")))
		return
	}

	data, err := cs.contactService.GetContactById(param)
	if err != nil {
		c.JSON(500, utils.ErrorMessages(utils.ErrorInternalServer, err))
		return
	}

	c.JSON(200, data)
}

func (cs *contactController) CreateContact(c *gin.Context) {
	var input entity.ContactInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, utils.ErrorMessages(utils.ErrorBadRequest, err))
		return
	}

	err := cs.contactService.CreateContact(input)
	if err != nil {
		c.JSON(500, utils.ErrorMessages(utils.ErrorInternalServer, err))
		return
	}

	c.JSON(201, gin.H{
		"message": "success create contact",
	})
}

func (cs *contactController) UpdateContactById(c *gin.Context) {
	var input entity.ContactInput

	param := c.Param("id")
	if param == "" {
		c.JSON(400, utils.ErrorMessages(utils.ErrorBadRequest, errors.New("parameter not valid")))
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, utils.ErrorMessages(utils.ErrorBadRequest, err))
		return
	}

	err := cs.contactService.UpdateContactById(param, input)
	if err != nil {
		c.JSON(500, utils.ErrorMessages(utils.ErrorInternalServer, err))
		return
	}

	c.JSON(200, gin.H{
		"message": "success update contact id: " + param,
	})
}

func (cs *contactController) DeleteContactById(c *gin.Context) {
	param := c.Param("id")
	if param == "" {
		c.JSON(400, utils.ErrorMessages(utils.ErrorBadRequest, errors.New("parameter not valid")))
		return
	}

	err := cs.contactService.DeleteContactById(param)
	if err != nil {
		c.JSON(500, utils.ErrorMessages(utils.ErrorInternalServer, err))
		return
	}

	c.JSON(200, gin.H{
		"message": "success delete contact id: " + param,
	})
}