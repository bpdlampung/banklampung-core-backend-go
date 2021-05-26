package middlewares

import (
	"banklampung-core/encryption"
	"banklampung-core/entities"
	"banklampung-core/errors"
	"banklampung-core/helpers"
	"github.com/gin-gonic/gin"
	"os"
)

func GetUserRequest(context *gin.Context) (*entities.User, error) {
	key, ok := os.LookupEnv("UA_KEY")

	if !ok {
		return nil, errors.NotFound("ua key must be required")
	}

	ua, ok := context.GetQuery("ua")

	if !ok {
		return nil, errors.NotFound("user request must be required")
	}

	decrypted, err := encryption.StringToTripleDesECBDecrypt(ua, key)

	if err != nil {
		return nil, err
	}

	userReq := entities.User{}
	err = helpers.JsonStringToStruct(*decrypted, &userReq)

	if err != nil {
		return nil, err
	}

	return &userReq, err
}
