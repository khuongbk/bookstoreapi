package main
import "fmt"
type Data struct {
	data []Customer
}
type Customer struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Total     int    `json:"total"`
	Email     string `json:"email"`
	Book_list string `json:"book_list"`
}

func main(){
	var data Data
	var cus Customer
	slice1 :=make([]Customer,0)

	cus.Book_list= "khuong"
	cus.Email="ggg"
	cus.Total=1333
	cus.Name="cong"
	for i:=1 ;i<10;i++{
		cus.Id=i
slice1=append(slice1,cus)
	}

//   slice :=make([]string,0)

data.data=slice1
fmt.Println(slice1)
fmt.Println(data)
}