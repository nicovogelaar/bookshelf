package app

type pagination struct {
	page         int `form:"page"`
	itemsPerPage int `form:"limit"`
}

func newPagination(page int, limit int) pagination {
	return pagination{page, limit}
}

func (p *pagination) limit() int {
	return p.itemsPerPage
}

func (p *pagination) offset() int {
	return p.itemsPerPage * (p.page - 1)
}
