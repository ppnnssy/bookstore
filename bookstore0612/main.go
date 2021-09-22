package main

import (
	"bookstore0612/controller"
	"net/http"
)


func main() {
	//设置处理静态资源，如css和js文件
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static"))))
	//直接去html页面
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages"))))
	http.HandleFunc("/main", controller.IndexHandler)

	//去登录
	http.HandleFunc("/login", controller.Login)
	//去注册
	http.HandleFunc("/regist", controller.Regist)
	//通过Ajax检验用户的用户名是否可用
	http.HandleFunc("/checkUserName", controller.CheckUserName)
	//获取带分页的图书
	http.HandleFunc("/getPageBooks", controller.GetPageBooks)

	//删除图书
	http.HandleFunc("/deleteBook",controller.DeleteBook)
	//去更新图书的页面
	http.HandleFunc("/toUpdateBookPage",controller.ToUpdateBook)
	//添加或更新图书
	http.HandleFunc("/updateBook",controller.UpdateOrAddBook)



	http.ListenAndServe(":8999", nil)
}
