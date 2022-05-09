package controller

import (
	"errors"
	"server-service/entity"
	"server-service/service"
	"server-service/utils"

	"github.com/gin-gonic/gin"
)

type membershipPaymentController struct {
	membershipPaymentService service.MembershipPaymentService
}

func NewMembershipPaymentController(membershipPaymentService service.MembershipPaymentService) *membershipPaymentController {
	return &membershipPaymentController{
		membershipPaymentService: membershipPaymentService,
	}
}

func (cs *membershipPaymentController) GetAllMembershipPayment(c *gin.Context) {
	datas, err := cs.membershipPaymentService.GetAllMembershipPayment()
	if err != nil {
		c.JSON(500, utils.ErrorMessages(utils.ErrorInternalServer, err))
		return
	}

	c.JSON(200, datas)
}

func (cs *membershipPaymentController) GetMembershipPaymentById(c *gin.Context) {
	param := c.Param("id")

	if param == "" {
		c.JSON(400, utils.ErrorMessages(utils.ErrorBadRequest, errors.New("parameter not valid")))
		return
	}

	data, err := cs.membershipPaymentService.GetMembershipPaymentById(param)
	if err != nil {
		c.JSON(500, utils.ErrorMessages(utils.ErrorInternalServer, err))
		return
	}

	c.JSON(200, data)
}

func (cs *membershipPaymentController) CreateMembershipPayment(c *gin.Context) {
	var input entity.MembershipPaymentInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, utils.ErrorMessages(utils.ErrorBadRequest, err))
		return
	}

	err := cs.membershipPaymentService.CreateMembershipPayment(input)
	if err != nil {
		c.JSON(500, utils.ErrorMessages(utils.ErrorInternalServer, err))
		return
	}

	c.JSON(201, gin.H{
		"message": "success create membershipPayment",
	})
}

func (cs *membershipPaymentController) UpdateMembershipPaymentById(c *gin.Context) {
	var input entity.MembershipPaymentInput

	param := c.Param("id")
	if param == "" {
		c.JSON(400, utils.ErrorMessages(utils.ErrorBadRequest, errors.New("parameter not valid")))
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, utils.ErrorMessages(utils.ErrorBadRequest, err))
		return
	}

	err := cs.membershipPaymentService.UpdateMembershipPaymentById(param, input)
	if err != nil {
		c.JSON(500, utils.ErrorMessages(utils.ErrorInternalServer, err))
		return
	}

	c.JSON(200, gin.H{
		"message": "success update membershipPayment id: " + param,
	})
}

func (cs *membershipPaymentController) DeleteMembershipPaymentById(c *gin.Context) {
	param := c.Param("id")
	if param == "" {
		c.JSON(400, utils.ErrorMessages(utils.ErrorBadRequest, errors.New("parameter not valid")))
		return
	}

	err := cs.membershipPaymentService.DeleteMembershipPaymentById(param)
	if err != nil {
		c.JSON(500, utils.ErrorMessages(utils.ErrorInternalServer, err))
		return
	}

	c.JSON(200, gin.H{
		"message": "success delete membershipPayment id: " + param,
	})
}
