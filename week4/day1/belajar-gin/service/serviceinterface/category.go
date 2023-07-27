package serviceinterface

import "belajar-gin/domain"

type CategoryServiceInterface interface {
	Create(category domain.AddCategory) domain.Category
	FindByName(categoryname string) *domain.Category
	FindById(categoryId int) *domain.Category
}
