package user

import (
	"github.com/bowoBp/myDate/internal/dto"
	"github.com/bowoBp/myDate/pkg/mapper"
	middleware2 "github.com/bowoBp/myDate/pkg/middleware"
	time2 "github.com/bowoBp/myDate/pkg/time"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type (
	RequestHandler struct {
		mapper mapper.MapperUtility
		clock  time2.Clock
		ctrl   ControllerInterface
		enigma middleware2.EnigmaUtility
	}
)

func (rh RequestHandler) Register(context *gin.Context) {
	var payload = RegisterPayload{}
	if err := rh.enigma.BindAndValidate(context, &payload); len(err) > 0 {
		context.JSON(http.StatusBadRequest, dto.DefaultInvalidInputFormResponse(err))
		return
	}

	res, err := rh.ctrl.AddUser(context.Request.Context(), payload)
	if err != nil {
		log.Println(err)
		switch {
		case rh.mapper.CompareSliceOfErr(registerErrs, err):
			context.JSON(http.StatusBadRequest, dto.DefaultErrorInvalidDataWithMessage(err.Error()))
			return
		default:
			context.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage(err.Error()))
			return
		}
	}
	context.JSON(http.StatusOK, res)
}

func (rh RequestHandler) Verify(context *gin.Context) {
	var payload = VerifyOtpPayload{}
	if err := rh.enigma.BindAndValidate(context, &payload); len(err) > 0 {
		context.JSON(http.StatusBadRequest, dto.DefaultInvalidInputFormResponse(err))
		return
	}

	res, err := rh.ctrl.VerifyOtp(context.Request.Context(), payload)
	if err != nil {
		log.Println(err)
		switch {
		case rh.mapper.CompareSliceOfErr(verifyOtpErrs, err):
			context.JSON(http.StatusBadRequest, dto.DefaultErrorResponseWithMessage(err.Error()))
			return
		default:
			context.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage(err.Error()))
			return
		}
	}
	context.JSON(http.StatusOK, res)
}
