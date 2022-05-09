package controller

import (
	"errors"
	"server-service/service"
	"server-service/entity"
	"server-service/utils"

	"github.com/gin-gonic/gin"
)

type categoryController struct {
	categoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) *categoryController {
	return &categoryController{
		categoryService: categoryService,
	}
}

func (cs *categoryController) GetAllCategory(c *gin.Context) {
	datas, err := cs.categoryService.GetAllCategory()
	if err != nil {
		c.JSON(500, utils.ErrorMessages(utils.ErrorInternalServer, err))
		return
	}

	c.JSON(200, datas)
}

func (cs *categoryController) GetCategoryById(c *gin.Context) {
	param := c.Param("id")

	if param == "" {
		c.JSON(400, utils.ErrorMessages(utils.ErrorBadRequest, errors.New("parameter not valid")))
		return
	}

	data, err := cs.categoryService.GetCategoryById(param)
	if err != nil {
		c.JSON(500, utils.ErrorMessages(utils.ErrorInternalServer, err))
		return
	}

	c.JSON(200, data)
}

func (cs *categoryController) CreateCategory(c *gin.Context) {
	var input entity.CategoryInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, utils.ErrorMessages(utils.ErrorBadRequest, err))
		return
	}

	err := cs.categoryService.CreateCategory(input)
	if err != nil {
		c.JSON(500, utils.ErrorMessages(utils.ErrorInternalServer, err))
		return
	}

	c.JSON(201, gin.H{
		"message": "success create category",
	})
}

func (cs *categoryController) UpdateCategoryById(c *gin.Context) {
	var input entity.CategoryInput

	param := c.Param("id")
	if param == "" {
		c.JSON(400, utils.ErrorMessages(utils.ErrorBadRequest, errors.New("parameter not valid")))
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, utils.ErrorMessages(utils.ErrorBadRequest, err))
		return
	}

	err := cs.categoryService.UpdateCategoryById(param, input)
	if err != nil {
		c.JSON(500, utils.ErrorMessages(utils.ErrorInternalServer, err))
		return
	}

	c.JSON(200, gin.H{
		"message": "success update category id: " + param,
	})
}

func (cs *categoryController) DeleteCategoryById(c *gin.Context) {
	param := c.Param("id")
	if param == "" {
		c.JSON(400, utils.ErrorMessages(utils.ErrorBadRequest, errors.New("parameter not valid")))
		return
	}

	err := cs.categoryService.DeleteCategoryById(param)
	if err != nil {
		c.JSON(500, utils.ErrorMessages(utils.ErrorInternalServer, err))
		return
	}

	c.JSON(200, gin.H{
		"message": "success delete category id: " + param,
	})
}