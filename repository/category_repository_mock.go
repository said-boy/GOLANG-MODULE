package repository

import (
	"github.com/said-boy/GOLANG-MODULE/entity"
	"github.com/stretchr/testify/mock"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

func (mock *CategoryRepositoryMock) FindById(id string) *entity.Category {
    args := mock.Mock.Called(id)
	if args.Get(0) == nil {
		return nil
	} else {
		result := args.Get(0).(entity.Category)
		return &result
	}
}