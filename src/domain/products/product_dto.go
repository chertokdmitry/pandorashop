package products

type Product struct {
	Id			int
	Url			string
	Name		string
	Price		int
	CategoryId	int
	Stock		int
}

type Products []Product
