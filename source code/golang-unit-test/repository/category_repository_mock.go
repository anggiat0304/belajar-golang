package repository

import (
	"golang-unit-test/entity"

	"github.com/stretchr/testify/mock"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

func (c CategoryRepositoryMock) FindById(id string) *entity.Category {
	arguments := c.Mock.Called(id)
	if arguments.Get(0) == nil {
		return nil
	}
	category := arguments.Get(0).(entity.Category)
	return &category
}
