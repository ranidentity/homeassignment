package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Path    string
	Handler http.HandlerFunc
	Method  string
}

func RegisterRoutes(container *Container) *gin.Engine {
	r := gin.Default()
	r.GET("/accounts/:account_id", container.AccountHandler.GetAccountInfo)
	r.POST("/accounts", container.AccountHandler.GetAccountInfo)
	r.POST("/accounts", container.TransactionHandler.SubmitTransaction)
	return r

}
