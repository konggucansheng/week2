package main

import (
	"fmt"
	"week2/dao"
)

func main() {
	customer, err := dao.GetResult()
	if err != nil {
		fmt.Printf("err : %+v", err)
		return
	}
	fmt.Println("customer : ", customer)

}
