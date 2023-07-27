package repositoryinterface

import "belajar-gin/domain"

type CategoryInterface interface {
	Add(category *domain.Category)
	FindByName(categoryname string) *domain.Category
	FindById(categoryId int) *domain.Category
}
