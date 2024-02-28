package api

import (
	"fmt"
	"github.com/bowoBp/myDate/internal/service/user"
	"github.com/bowoBp/myDate/pkg/db"
	"github.com/bowoBp/myDate/pkg/environment"
	mailing2 "github.com/bowoBp/myDate/pkg/mailing"
	maker2 "github.com/bowoBp/myDate/pkg/maker"
	mapper2 "github.com/bowoBp/myDate/pkg/mapper"
	"github.com/bowoBp/myDate/pkg/middleware"
	time2 "github.com/bowoBp/myDate/pkg/time"
	"github.com/gin-gonic/gin"
	"log"
)

func Default() *Api {
	server := gin.Default()
	sqlConn, err := db.Default()
	if err != nil {
		log.Println(err)
		panic(fmt.Sprintf("panic at db connection: %s", err.Error()))
	}
	fmt.Println("database connected: 5432")

	maker := maker2.DefaultMaker()
	mailing := mailing2.NewConfig()
	enigma := middleware.NewEnigma()
	clock := time2.Default()
	env := environment.NewEnvironment()
	mapper := mapper2.Default()

	var routers = []Router{
		user.NewRoute(sqlConn, maker, mailing, enigma, clock, env, mapper),
	}
	return &Api{
		server:  server,
		routers: routers,
	}
}
