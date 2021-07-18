package model

//分页信息
type Page struct {
	Books []*Book
	PageNo int64//当前页
	PageSize int64//每页大小
	PageTotal int64//总页数
	RecordTotal int64//总记录数
	MinPrice string
	MaxPrice string
}

//跳转边界判断
func (p *Page) IsHasPrev() bool {
	return p.PageNo>1
}
func (p *Page) IsHasNext() bool {
	return p.PageNo < p.PageTotal
}

//获取上一页、下一页
func (p *Page)GetPrevPageNo() int64 {
	if p.IsHasPrev() {
		return p.PageNo-1
	}else{
		return 1
	}
}
func (p *Page)GetNextPageNo() int64 {
	if p.IsHasNext() {
		return p.PageNo+1
	}else{
		return p.PageTotal
	}
}