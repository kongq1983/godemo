package mycom_entity

import "fmt"

func init()  {
	fmt.Printf("Employee init is called. \n")
}

/**
  (name string,age int8,cname string) 是返回值
 */
func Empoyee()(name string,age int8,cname string)  {
	return name,age,cname;
}