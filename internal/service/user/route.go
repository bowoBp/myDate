package user

import (
	"github.com/bowoBp/myDate/internal/adapter/repository"
	global_usage "github.com/bowoBp/myDate/internal/service/global-usage"
	"github.com/bowoBp/myDate/pkg/environment"
	"github.com/bowoBp/myDate/pkg/mailing"
	"github.com/bowoBp/myDate/pkg/maker"
	"github.com/bowoBp/myDate/pkg/mapper"
	"github.com/bowoBp/myDate/pkg/middleware"
	time2 "github.com/bowoBp/myDate/pkg/time"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type (
	Router struct {
		rq *RequestHandler
	}
)

func NewRoute(
	db *gorm.DB,
	maker maker.Generator,
	mailing mailing.SendInBlueInterface,
	enigma middleware.EnigmaUtility,
	clock time2.Clock,
	env environment.Environment,
	mapper mapper.MapperUtility,
) *Router {
	return &Router{
		rq: &RequestHandler{
			mapper: mapper,
			clock:  clock,
			ctrl: &Controller{
				uc: &UseCase{
					userRepo:    repository.NewUserRepo(db),
					maker:       maker,
					env:         env,
					clock:       clock,
					globalUsage: global_usage.NewUsecase(),
					smtp:        mailing,
				},
				mapper: mapper,
			},
			enigma: enigma,
		},
	}
}

func (r Router) Route(router *gin.RouterGroup) {
	account := router.Group("/account")

	account.POST(""+
		"/register",
		r.rq.Register,
	)

	account.POST(
		"/verify-otp",
		r.rq.Verify,
	)
}
