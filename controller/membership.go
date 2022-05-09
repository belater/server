package controller

import (
	"errors"
	"server-service/entity"
	"server-service/service"
	"server-service/utils"

	"github.com/gin-gonic/gin"
)

type membershipController struct {
	membershipService service.MembershipService
}

func NewMembershipController(membershipService service.MembershipService) *membershipController {
	return &membershipController{
		membershipService: membershipService,
	}
}

func (cs *membershipController) GetAllMembership(c *gin.Context) {
	datas, err := cs.membershipService.GetAllMembership()
	if err != nil {
		c.JSON(500, utils.ErrorMessages(utils.ErrorInternalServer, err))
		return
	}

	c.JSON(200, datas)
}

func (cs *membershipController) GetMembershipById(c *gin.Context) {
	param := c.Param("id")

	if param == "" {
		c.JSON(400, utils.ErrorMessages(utils.ErrorBadRequest, errors.New("parameter not valid")))
		return
	}

	data, err := cs.membershipService.GetMembershipById(param)
	if err != nil {
		c.JSON(500, utils.ErrorMessages(utils.ErrorInternalServer, err))
		return
	}

	c.JSON(200, data)
}

func (cs *membershipController) CreateMembership(c *gin.Context) {
	var input entity.MembershipInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, utils.ErrorMessages(utils.ErrorBadRequest, err))
		return
	}

	err := cs.membershipService.CreateMembership(input)
	if err != nil {
		c.JSON(500, utils.ErrorMessages(utils.ErrorInternalServer, err))
		return
	}

	c.JSON(201, gin.H{
		"message": "success create membership",
	})
}

func (cs *membershipController) UpdateMembershipById(c *gin.Context) {
	var input entity.MembershipInput

	param := c.Param("id")
	if param == "" {
		c.JSON(400, utils.ErrorMessages(utils.ErrorBadRequest, errors.New("parameter not valid")))
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, utils.ErrorMessages(utils.ErrorBadRequest, err))
		return
	}

	err := cs.membershipService.UpdateMembershipById(param, input)
	if err != nil {
		c.JSON(500, utils.ErrorMessages(utils.ErrorInternalServer, err))
		return
	}

	c.JSON(200, gin.H{
		"message": "success update membership id: " + param,
	})
}

func (cs *membershipController) DeleteMembershipById(c *gin.Context) {
	param := c.Param("id")
	if param == "" {
		c.JSON(400, utils.ErrorMessages(utils.ErrorBadRequest, errors.New("parameter not valid")))
		return
	}

	err := cs.membershipService.DeleteMembershipById(param)
	if err != nil {
		c.JSON(500, utils.ErrorMessages(utils.ErrorInternalServer, err))
		return
	}

	c.JSON(200, gin.H{
		"message": "success delete membership id: " + param,
	})
}
