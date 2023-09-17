package shared

import (
	"github.com/google/uuid"
)

type UUID string

func NewUUID() UUID {
	u, err := uuid.NewUUID()
	if err != nil {
		panic(err.Error())
	}
	return UUID(u.String())
}
