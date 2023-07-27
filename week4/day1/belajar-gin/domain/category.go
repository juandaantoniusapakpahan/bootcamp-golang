package domain

type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type AddCategory struct {
	Name string `json:"name"`
}

func NewCategory(category AddCategory) Category {
	return Category{
		Name: category.Name,
	}
}
