package service

import (
	"errors"
	"github.com/said-boy/GOLANG-MODULE/entity"
	"github.com/said-boy/GOLANG-MODULE/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return nil, errors.New("Not Found")
	} else {
			return category, nil
	}
}