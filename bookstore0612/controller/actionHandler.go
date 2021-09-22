package controller

import (
	"bookstore0612/dao"
	"fmt"
	"html/template"
	"net/http"
)

// IndexHandler 去首页
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	pageNo:=r.FormValue("pageNo")
	if pageNo==""{
		pageNo="1"
	}
	page,_:=dao.GetPageBooks(pageNo)
	//解析模板
	for k,v:=range page.Books{
		fmt.Printf("第%d本书是：%v\n",k+1,v)
	}
	t := template.Must(template.ParseFiles("views/index.html"))
	//执行
	t.Execute(w, page)
}
