package helpers

import uuid "github.com/satori/go.uuid"

func UUID() string {
	u4 := uuid.NewV4()
	return u4.String()
}
