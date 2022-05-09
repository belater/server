package routes

import (
	"server-service/controller"
	"server-service/repository"
	"server-service/service"

	"github.com/gin-gonic/gin"
)

var (
	UserRepository = repository.NewUserRepository(DB)
	UserService    = service.NewUserService(UserRepository)
	UserController = controller.NewUserController(UserService)
)

func UserRoute(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.GET("/users/all", UserController.GetAllUser) //for admin
		v1.GET("users/:user_id", UserController.GetUserById)

		// verify email, atau nomer wa
		v1.POST("/users/register", UserController.RegisterUser)
		v1.POST("/users/login", UserController.LoginUser)

		// update surname, fullname (not email, and not phone number)
		v1.PUT("/users/:id", UserController.UpdateUserById)

		// soft delete, active false
		v1.DELETE("/users/inactive/:user_id", UserController.DeleteUserById)

		// route forgot password
		// set email / nomer wa
		v1.POST("/users/forgot-password")

		// route reset password
		v1.POST("/users/reset-password")
	}
}
