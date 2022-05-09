package routes

import (
	"server-service/repository"
	"server-service/service"
	"server-service/controller"

	"github.com/gin-gonic/gin"
)

var (
	MembershipRepository = repository.NewMembershipRepository(DB)
	MembershipService = service.NewMembershipService(MembershipRepository)
	MembershipController = controller.NewMembershipController(MembershipService)
)

func MembershipRoute(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.GET("/memberships/all", MembershipController.GetAllMembership)
		v1.GET("memberships/:id", MembershipController.GetMembershipById)
		v1.POST("/memberships", MembershipController.CreateMembership)
		v1.PUT("/memberships/:id", MembershipController.UpdateMembershipById)
		v1.DELETE("/memberships/:id", MembershipController.DeleteMembershipById)
	}
}
