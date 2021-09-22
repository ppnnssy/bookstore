package dao

import (
	"bookstore0612/model"
	"bookstore0612/utils"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"math"
	"strconv"
)

//获取数据库中所有的图书
func GetBooks() ([]*model.Book, error) {
	sqlStr := "select ID, Title ,Author ,Price,Sales,Stock  from books"
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		fmt.Println("读取书本信息出错", err)
		return nil, err
	}

	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock)
		books = append(books, book)
	}
	return books, nil
}

//添加图书
func AddBook(b *model.Book) error {
	sqlStr := "insert into books(title,author,price,sales,stock,img_path) value(?,?,?,?,?,?)"
	_, err := utils.Db.Exec(sqlStr, b.Title, b.Author, b.Price, b.Sales, b.Stock, b.Img_path)
	if err != nil {
		return err
	}
	return nil
}

//删除图书（根据id）
func DeleteBook(bookId string) error {
	sqlStr := "delete from books where id=?"
	_, err := utils.Db.Exec(sqlStr, bookId)
	if err != nil {
		return err
	}
	return nil
}

//根据id查询一个图书
func GetBookById(bookID string) *model.Book {
	sqlStr := "select id,title,author,price,sales,stock,img_path from books where id=?"
	row := utils.Db.QueryRow(sqlStr, bookID)
	book := &model.Book{}
	row.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.Img_path)
	return book
}

//UpdateBook 根据图书id更新图书信息
func UpdateBook(book *model.Book) error {
	sqlStr := "update books set title=?,author=?,price=?,sales=?,stock=? where id=?"
	_, err := utils.Db.Exec(sqlStr, book.Title, book.Author, book.Price, book.Sales, book.Stock, book.ID)
	if err != nil {
		return err
	}
	return nil
}

//获取带分页的图书信息
func GetPageBooks(pageNo string) (*model.Page, error) {
	//先把pageNo转成int64
	intPageNo,err:=strconv.ParseInt(pageNo,10,64)
	if err!=nil {
		fmt.Println("pageNo类型转换失败：",err)
		return nil,err
	}

	sqlStr1 := "select count(*) from books" //count(*)表示查询结果的个数
	//设置总记录数
	var totalRecord int64
	row:=utils.Db.QueryRow(sqlStr1)
	row.Scan(&totalRecord)
	//设置每页只显示4条记录
	var pageSize int64
	pageSize=4
	//获取总页数
	var totalPageNum int64
	fPagenum:=float64(totalRecord)/float64(pageSize)
	totalPageNum=int64(math.Ceil(fPagenum))

	//获取当前页中的所有图书
	sqlStr2 := "select ID, Title ,Author ,Price,Sales,Stock,img_path  from books limit ?,?"
	rows,err:=utils.Db.Query(sqlStr2,(intPageNo-1)*pageSize,pageSize)
	if err!=nil{
		fmt.Println("数据查询失败：",err)
		return nil,err
	}
	var books []*model.Book
	for rows.Next(){
		book:=&model.Book{}
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock,&book.Img_path)
		books=append(books,book)
	}


	//创建page
	page:=&model.Page{
		Books: books,
		PageNo: intPageNo,
		PageSize: pageSize,
		TotalPageNo: totalPageNum,
		TotalRecord: totalRecord,
	}
	return page,nil
}
