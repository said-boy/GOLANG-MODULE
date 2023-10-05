package service

import (
	"testing"
	"github.com/said-boy/GOLANG-MODULE/entity"
	"github.com/said-boy/GOLANG-MODULE/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_GET_NotFound(t *testing.T){

	categoryRepository.Mock.On("FindById", "1").Return(nil)

	category, err := categoryService.Get("1")
	assert.Nil(t, category)
	assert.NotNil(t, err)

}

func TestCategoryService_GET_success(t *testing.T) {

	category := entity.Category{
		Id: "2",
        Name: "test",
	}
	categoryRepository.Mock.On("FindById", "2").Return(category)
	result, err := categoryService.Get("2")
	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Equal(t, category.Id, result.Id)
	assert.Equal(t, category.Name, result.Name)

}