package usecase

import "github.com/google/uuid"

func Uuid4() string {
	uuidObj, _ := uuid.NewRandom()
	return uuidObj.String()
}
