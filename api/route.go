package api

import (
	"CC-Nicepay/api/controller"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	payment := r.Group("/payment")
	{
		payment.POST("/registration", controller.RegistrationController)
		payment.POST("/inquiry", controller.StatusInquiry)
		payment.POST("/callback", controller.CallBackPaymentNicePay)
		payment.POST("/payment", controller.PaymentCCNicePay)
	}
}
