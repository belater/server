package routes

import (
	"server-service/repository"
	"server-service/service"
	"server-service/controller"

	"github.com/gin-gonic/gin"
)

var (
	MembershipPaymentRepository = repository.NewMembershipPaymentRepository(DB)
	MembershipPaymentService = service.NewMembershipPaymentService(MembershipPaymentRepository)
	MembershipPaymentController = controller.NewMembershipPaymentController(MembershipPaymentService)
)

func MembershipPaymentRoute(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.GET("/membershipPayments/all", MembershipPaymentController.GetAllMembershipPayment)
		v1.GET("membershipPayments/:id", MembershipPaymentController.GetMembershipPaymentById)
		v1.POST("/membershipPayments", MembershipPaymentController.CreateMembershipPayment)
		v1.PUT("/membershipPayments/:id", MembershipPaymentController.UpdateMembershipPaymentById)
		v1.DELETE("/membershipPayments/:id", MembershipPaymentController.DeleteMembershipPaymentById)
	}
}
