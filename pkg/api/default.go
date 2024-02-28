package api

import (
	"fmt"
	"github.com/bowoBp/myDate/pkg/db"
	"github.com/gin-gonic/gin"
	"log"
)

func Default() *Api {
	server := gin.Default()
	_, err := db.Default()
	if err != nil {
		log.Println(err)
		panic(fmt.Sprintf("panic at db connection: %s", err.Error()))
	}
	fmt.Println("database connected: 5432")
	//var routers = []Router{
	//	employee.NewRoute(sqlConn),
	//}
	return &Api{
		server:  server,
		routers: nil,
	}
}
