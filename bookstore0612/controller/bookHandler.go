package controller

import (
	"bookstore0612/dao"
	"bookstore0612/model"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// GetPageBooks 获取带分页的图书
func GetPageBooks(w http.ResponseWriter,r *http.Request)  {
	pageNo:=r.FormValue("pageNo")
	if pageNo==""{
		pageNo="1"
	}
	page,_:=dao.GetPageBooks(pageNo)
	//解析模板文件
	t:=template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	t.Execute(w,page)
}

// AddBook 添加图书。浏览器中发过来需要添加的图书的信息，接收信息并添加到数据库中
//func AddBook(w http.ResponseWriter,r *http.Request)  {
//	//获取图书信息
//	fmt.Println("添加一本图书")
//	title:=r.PostFormValue("title")
//	author:=r.PostFormValue("author")
//	price:=r.PostFormValue("price")
//	sales:=r.PostFormValue("sales")
//	stock:=r.PostFormValue("stock")
//	fPrice,_:= strconv.ParseFloat(price,64)
//	intSales,_:=strconv.ParseInt(sales,10,0)
//	intStock,_:=strconv.ParseInt(stock,10,0)
//	book:=&model.Book{
//		Title: title,
//		Author: author,
//		Price:fPrice,
//		Sales: int(intSales),
//		Stock: int(intStock),
//		Img_path: "static/img/default.jpg",
//	}
//	dao.AddBook(book)
//	fmt.Println("添加图书成功：",book)
//	/*
//		我这里数据库字段设置的有问题，价格小数点设置成0位，暂时不想改
//	*/
//
//	//调用getbook处理器函数，再查询一次数据库，返回到图书界面
//	GetBooks(w,r)
//}

// DeleteBook 删除图书
func DeleteBook(w http.ResponseWriter,r *http.Request)  {
	bookID:=r.FormValue("bookId") //从网页表单中传过来的
	err:=dao.DeleteBook(bookID)
	if err!=nil {
		fmt.Println("删除图书出错")
	}
	//重新查询所有图书
	GetPageBooks(w,r)
}

// ToUpdateBook 去更新或添加图书页面
func ToUpdateBook(w http.ResponseWriter,r *http.Request)  {
	bookID:=r.FormValue("bookId")
	book:=dao.GetBookById(bookID)

	//通过前端传过来的图书form中有没有id来判断是更新还是添加图书。对应的book_edit.html也要进行相应判断
	if book.ID>0{//传过来的有id，证明要更新图书
		//解析模板
		t:=template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
		//执行
		t.Execute(w,book)
	}else {//没有id，要添加图书
		//解析模板
		t:=template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
		//执行
		t.Execute(w,"")
	}




}

//更新图书
func UpdateOrAddBook(w http.ResponseWriter,r *http.Request){
	id:=r.PostFormValue("bookId")//<input type="hidden"name="bookId" value="{{.ID}}"/>前端隐式发送bookid
	title:=r.PostFormValue("title")
	author:=r.PostFormValue("author")
	price:=r.PostFormValue("price")
	sales:=r.PostFormValue("sales")
	stock:=r.PostFormValue("stock")

	intID,_:=strconv.ParseInt(id,10,0)
	fPrice,_:= strconv.ParseFloat(price,64)
	intSales,_:=strconv.ParseInt(sales,10,0)
	intStock,_:=strconv.ParseInt(stock,10,0)
	book:=&model.Book{
		ID: int(intID),
		Title: title,
		Author: author,
		Price:fPrice,
		Sales: int(intSales),
		Stock: int(intStock),
		Img_path: "static/img/default.jpg",
	}
	if book.ID>0 {//前端传来了ID说明是更新图书
		err:=dao.UpdateBook(book)
		if err!=nil {
			fmt.Println("更新图书失败：",err)
			return
		}
	}else {//id为0说明前端没有传来id，是添加图书
		dao.AddBook(book)
		fmt.Println("添加图书成功：",book)
	}

	GetPageBooks(w,r)
}