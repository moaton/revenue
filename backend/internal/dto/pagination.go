package dto



const MaxPaginationSize = 50

type Pagination struct {
	Page uint64 `json:"page"`
	Size uint64 `json:"size"`
}

func (p *Pagination) Validate(){
	if p.Size > MaxPaginationSize{
		p.Size = MaxPaginationSize
	}
	if p.Size <= 0 {
		p.Size = 10
	}
	if p.Page <=0 {
		p.Page = 1
	}
}