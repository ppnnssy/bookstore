package controller

import (
	"bookstore0612/dao"
	"html/template"
	"net/http"
)

//Login 处理用户登录的函数
func Login(w http.ResponseWriter, r *http.Request) {
	//从浏览器提交的表单中获取用户名和密码
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	//调用userdao中验证用户名和密码的方法
	user, _ := dao.CheckUserNameAndPassword(username, password)
	if user.ID > 0 {
		//用户名和密码正确
		t := template.Must(template.ParseFiles("views/pages/user/login_success.html"))
		t.Execute(w, "")
	} else {
		//用户名或密码不正确
		t := template.Must(template.ParseFiles("views/pages/user/login.html"))
		t.Execute(w, "用户名或密码不正确！")
	}
}

//Regist 处理用户的函注册数
func Regist(w http.ResponseWriter, r *http.Request) {
	//获取用户名和密码
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")
	//调用userdao中验证用户名和密码的方法
	user, _ := dao.CheckUserName(username)
	if user.ID > 0 {
		//用户名已存在
		t := template.Must(template.ParseFiles("views/pages/user/regist.html"))
		t.Execute(w, "用户名已存在！")
	} else {
		//用户名可用，将用户信息保存到数据库中
		dao.SaveUser(username, password, email)
		//用户名和密码正确
		t := template.Must(template.ParseFiles("views/pages/user/regist_success.html"))
		t.Execute(w, "")
	}
}

//检查用户名是否可用
func CheckUserName(w http.ResponseWriter,r *http.Request) {
	//获取用户名
	username:=r.PostFormValue("username")
	user,_:=dao.CheckUserName(username)
	if user.ID>0{
		//用户名已存在
		w.Write([]byte("用户名已存在"))
	}else {
		//用户名不存在
		w.Write([]byte("<font style='color:green'>用户名可用</font>"))
	}
}

