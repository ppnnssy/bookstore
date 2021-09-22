package dao

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("测试bookdao中的方法")
	m.Run()
}

func TestBook(t *testing.T) {
	fmt.Println("测试bookdao中的函数")
	t.Run("测试更新图书", testGetPageBooks)
}

func testGetPageBooks(t *testing.T) {
	page,err:=GetPageBooks("1")
	if err!=nil{
		//fmt.Println(err)
		return
	}
	fmt.Println("当前页是：",page.PageNum)
	fmt.Println(page.TotalPageNum)
	fmt.Println(page.PageSize)
	fmt.Println(page.TotalRecord)

	for _,v:=range page.Books{
		fmt.Println("图书的信息是：",v)
	}
}
