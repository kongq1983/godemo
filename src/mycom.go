package main

import "fmt"
import . "mycom/entity"

func main() {

	var user User
	user.Name = "admin"
	user.Province = "zhejiang"

	fmt.Printf("name=%s,province=%s \n", user.Name, user.Province)

}
