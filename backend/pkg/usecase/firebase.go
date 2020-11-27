package usecase

import (
	"link-note/backend/pkg/model"
)

type UserInteractor struct {
	UserRepository UserRepository
}

type UserRepository interface {
	InsertData(data model.Content)
	DeleteData(uid, id string) error
	GetData(uid string) []model.Content
	AuthJWT(jwt string) error
}

func (fb *UserInteractor) InsertData(data model.Content) {
	fb.UserRepository.InsertData(data)
	return
}

func (fb *UserInteractor) DeleteData(uid, id string) error {
	err := fb.UserRepository.DeleteData(uid, id)
	return err
}

func (fb *UserInteractor) GetData(uid string) []model.Content {
	data := fb.UserRepository.GetData(uid)
	return data
}

func (fb *UserInteractor) AuthJWT(jwt string) error {
	err := fb.UserRepository.AuthJWT(jwt)
	return err
}
