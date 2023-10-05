package repository

import "github.com/said-boy/GOLANG-MODULE/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}