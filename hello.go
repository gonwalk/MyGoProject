package main

import (
	"fmt"
	"strconv"
)

func main(){
	fmt.Println("hello world!")
	strNum := "228809"
	intNum , err := strconv.ParseInt(strNum, 10, 64)
	if err != nil{
		fmt.Println("convert error:", err)
	}
	fmt.Println("convert success:", intNum)
}
