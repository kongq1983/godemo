package mycom_entity

import "fmt"

func init()  {
	fmt.Printf("City init is called. \n")
}

/**
  (name string,age int8,cname string) 是返回值
 */
func City()(name string)  {
	return name;
}