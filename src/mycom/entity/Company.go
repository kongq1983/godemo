package mycom_entity

import "fmt"

func init()  {
	fmt.Printf("Company init is called. \n")
}

/**
  (name string,age int8,cname string) 是返回值
 */
func Company()(name string)  {
	return name;
}