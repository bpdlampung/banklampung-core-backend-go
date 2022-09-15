package sentry

import (
	"github.com/bpdlampung/banklampung-core-backend-go/logs"
	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
)

type Sentry struct {
	logger logs.Collections
}

func InitConnection(DSN string) Sentry {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: DSN,
	})

	if err != nil {
		panic(err)
	}

	return Sentry{}
}

func (s Sentry) RegisterToRouter(ginRouter *gin.Engine) {
	ginRouter.Use(sentrygin.New(sentrygin.Options{
		Repanic: true,
	}))
}
