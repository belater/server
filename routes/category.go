package routes

import (
	"server-service/repository"
	"server-service/service"
	"server-service/controller"

	"github.com/gin-gonic/gin"
)

var (
	CategoryRepository = repository.NewCategoryRepository(DB)
	CategoryService = service.NewCategoryService(CategoryRepository)
	CategoryController = controller.NewCategoryController(CategoryService)
)

func CategoryRoute(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.GET("/categories/all", CategoryController.GetAllCategory)
		v1.GET("categories/:id", CategoryController.GetCategoryById)
		v1.POST("/categories", CategoryController.CreateCategory)
		v1.PUT("/categories/:id", CategoryController.UpdateCategoryById)
		v1.DELETE("/categories/:id", CategoryController.DeleteCategoryById)
	}
}
