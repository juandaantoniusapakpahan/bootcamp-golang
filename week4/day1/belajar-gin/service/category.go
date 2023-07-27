package service

import (
	"belajar-gin/domain"
	"belajar-gin/repository/repositoryinterface"
	"belajar-gin/service/serviceinterface"
)

type CategoryServiceImplement struct {
	CategoryRepository repositoryinterface.CategoryInterface
}

func NewCategoryServicelement(repositoy repositoryinterface.CategoryInterface) serviceinterface.CategoryServiceInterface {
	return &CategoryServiceImplement{
		CategoryRepository: repositoy,
	}
}

func (cs *CategoryServiceImplement) Create(category domain.AddCategory) domain.Category {
	newCategory := domain.NewCategory(category)
	cs.CategoryRepository.Add(&newCategory)
	return newCategory
}

func (cs *CategoryServiceImplement) FindByName(categoryname string) *domain.Category {
	category := cs.CategoryRepository.FindByName(categoryname)

	return category
}

func (cs *CategoryServiceImplement) FindById(categoryId int) *domain.Category {
	category := cs.CategoryRepository.FindById(categoryId)
	return category
}
