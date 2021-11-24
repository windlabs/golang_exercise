package main

import (
	"demo/base/sql/common"
	"fmt"
)
type goUser struct {
	Id int
	Name string
}
func main(){

	db,err := common.GetDb()
	if err != nil {
		fmt.Println("db failed.")
	}

	list := []goUser{}
	err = db.Table("user").Find(&list).Error
	if err != nil {
		fmt.Println("select table failed.")
	}

	fmt.Println(list)

}