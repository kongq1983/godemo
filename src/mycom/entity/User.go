package mycom_entity

/**
  注意 get  set 不知道能不能用
*/
type User struct {

	/** 首字母大写，允许其他包访问，可以直接使用User.Name="admin"  也可以通过setName、getName */
	Name string

	Province string

	address string

	/** 不允许外面的包访问，本包可以使用setAge  getAge*/
	age int
}
