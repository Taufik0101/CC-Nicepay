package controller

import (
	"CC-Nicepay/api/dto"
	"CC-Nicepay/api/service"
	"CC-Nicepay/api/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegistrationController(ctx *gin.Context) {
	var DTOCreateTransaction dto.CreateTransactionNicePay
	errCreate := ctx.ShouldBind(&DTOCreateTransaction)

	if errCreate != nil {
		response := utils.BuildErrorResponse("Failed to parsing", errCreate.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
	} else {
		if DTOCreateTransaction.Amount < 10000 {
			resp := utils.BuildResponse(false, "Minimal Nominal Harus 10000", nil)
			ctx.JSON(http.StatusBadRequest, resp)
			return
		}

		res, _ := service.RegistrationCCNicePay(DTOCreateTransaction)
		if res == nil {
			resp := utils.BuildResponse(false, "Something Went Wrong", nil)
			ctx.JSON(http.StatusBadRequest, resp)
			return
		}

		resp := utils.BuildResponse(true, "Registration Credit Card Success", res)
		ctx.JSON(http.StatusOK, resp)
	}
}

func StatusInquiry(ctx *gin.Context) {
	var DTOStatusInquiry dto.RequestStatusInquiry
	errCreate := ctx.ShouldBind(&DTOStatusInquiry)

	if errCreate != nil {
		response := utils.BuildErrorResponse("Failed to parsing", errCreate.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
	} else {
		res, _ := service.CheckStatusInquiry(DTOStatusInquiry)
		resp := utils.BuildResponse(true, "Registration Credit Card Success", res)
		ctx.JSON(http.StatusOK, resp)
	}
}

func CallBackPaymentNicePay(ctx *gin.Context) {
	var dataForCallback map[string]interface{}
	err := json.NewDecoder(ctx.Request.Body).Decode(&dataForCallback)
	if err != nil {
		return
	}

	fmt.Println(dataForCallback)
}

func PaymentCCNicePay(ctx *gin.Context) {
	var DTORequestPayment dto.RequestPaymentCCNicePay
	errCreate := ctx.ShouldBind(&DTORequestPayment)

	if errCreate != nil {
		response := utils.BuildErrorResponse("Failed to parsing", errCreate.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
	} else {
		res, _ := service.PaymentCCNicePay(DTORequestPayment)
		if !res {
			response := utils.BuildResponse(false, "Failed To Simulate Payment", res)
			ctx.JSON(http.StatusBadRequest, response)
			return
		}
		resp := utils.BuildResponse(true, "Success To Simulate Payment", res)
		ctx.JSON(http.StatusOK, resp)
	}
}
