package entities

type PaginationOptions struct {
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Sort  string `json:"order"`
}

func (p *PaginationOptions) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *PaginationOptions) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}

func (p *PaginationOptions) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *PaginationOptions) GetSort() string {
	if p.Sort == "" {
		p.Sort = "Id desc"
	}
	return p.Sort
}
