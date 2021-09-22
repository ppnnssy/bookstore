package model
//对页面上显示的书单做分页处理
type Page struct {
	Books []*Book //存放每一页中的图书
	PageNo int64 //当前页码
	PageSize int64 //每页显示的图书数量
	TotalPageNo int64 //通过计算得到，总页数
	TotalRecord int64 //图书的总数量，通过查询数据库得到
}

//是否有上一页
func(p *Page)IsHasPrev()bool{
	return p.PageNo>1
}
//是否有下一页
func(p *Page)IsHasNext()bool{
	return p.PageNo<p.TotalPageNo
}

//获取上一页
func (p *Page)GetPrevPageNo()int64  {
	if p.IsHasPrev(){
		return p.PageNo-1
	}else {
		return 1
	}
}

//获取下一页
func(p *Page) GetNextPageNo() int64 {
	if p.IsHasNext(){
		return p.PageNo+1
	}else {
		return p.TotalPageNo
	}
}