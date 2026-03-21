package service

import (
	"golang-unit-test/entity"
	"golang-unit-test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_Get(t *testing.T) {

	categoryRepository.Mock.On("FindById", "1").Return(nil)
	category, err := categoryService.Get("1")
	assert.NotNil(t, err)
	assert.Nil(t, category)
}
func TestCategoryService_GetFound(t *testing.T) {
	categoryMockData := entity.Category{
		Id:   "1",
		Name: "Elektronik",
	}

	categoryRepository.Mock.On("FindById", "1").Return(categoryMockData)
	category, err := categoryService.Get("1")
	assert.NotNil(t, category)
	assert.Equal(t, category.Id, category.Id)
	assert.Equal(t, category.Name, category.Name)
	assert.Nil(t, err)
}
