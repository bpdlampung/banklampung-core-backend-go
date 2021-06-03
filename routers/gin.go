package routers

import (
	"github.com/bpdlampung/banklampung-core-backend-go/enums"
	"github.com/bpdlampung/banklampung-core-backend-go/errors"
	"github.com/bpdlampung/banklampung-core-backend-go/responses"
	"github.com/gin-gonic/gin"
	"os"
)

//var ginRouter *gin.Engine

func GetRouter() *gin.Engine {
	if enums.Environment(os.Getenv("ENVIRONMENT")) == enums.Production {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	router.NoRoute(func(c *gin.Context) {
		responses.Error(c, errors.NotFound("Page Not Found"))
	})

	router.NoMethod(func(c *gin.Context) {
		responses.Error(c, errors.MethodNotAllowed("Method Not Allowed"))
	})

	return router
}
