package sharedmodel

type PagingDTO struct {
	Page  int   `json:"page" form:"page"`
	Limit int   `json:"limit" form:"limit"`
	Total int64 `json:"total"`
}

// Process normalizes the paging parameters.
func (p *PagingDTO) Process() {
	if p.Page <= 0 {
		p.Page = 1
	}
	
	if p.Limit <= 0 || p.Limit > 1000 {
		p.Limit = 10
	}
}
