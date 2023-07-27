package repository

import (
	"belajar-gin/domain"
	"belajar-gin/repository/repositoryinterface"
)

type CategoryImplement struct {
	Categories []*domain.Category
}

func NewCategory() repositoryinterface.CategoryInterface {
	return &CategoryImplement{}
}

func (ci *CategoryImplement) Add(category *domain.Category) {
	id := len(ci.Categories) + 1
	category.Id = id
	ci.Categories = append(ci.Categories, category)
}

func (ci *CategoryImplement) FindByName(categoryname string) *domain.Category {
	var category domain.Category
	for _, v := range ci.Categories {
		if v.Name == categoryname {
			category = *v
		}
	}
	if category == (domain.Category{}) {
		panic("category not found")
	}
	return &category
}

func (ci *CategoryImplement) FindById(categoryId int) *domain.Category {
	var category domain.Category
	for _, v := range ci.Categories {
		if v.Id == categoryId {
			category = *v
		}
	}
	if category == (domain.Category{}) {
		panic("category not found")
	}
	return &category
}
