package categories

type Category struct {
	Id   string
	Name string
}

type Categories []Category

var CustomCategories = Categories{
	{"20", "Star Wars"},
	{"22", "Harry Potter"},
	{"21", "Disney"},
}
