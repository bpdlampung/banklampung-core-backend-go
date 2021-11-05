package routers

import (
	"fmt"
	"github.com/bpdlampung/banklampung-core-backend-go/enums"
	"github.com/bpdlampung/banklampung-core-backend-go/errors"
	"github.com/bpdlampung/banklampung-core-backend-go/responses"
	"github.com/gin-gonic/gin"
	"os"
)

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

	router.Use(CatchError())

	return router
}

func CatchError() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				responses.Error(c, errors.InternalServerError(fmt.Sprintf("Something when wrong, Panic System :: %v", err)))
			}
		}()
		c.Next()
	}
}