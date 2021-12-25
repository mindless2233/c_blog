package entity

type User struct {
	//  `gorm:"column:beast_id"`
	id       int    `gorm:"column:id"`
	Userid   string `gorm:"column:user_id"`
	Username string `gorm:"column:user_name" json:"username"`
	Password string `gorm:"column:user_pwd"  json:"password"`
	age      int    `gorm:"column:user_age"`
}

func (User) TableName() string {
	//实现TableName接口，以达到结构体和表对应，如果不实现该接口，并未设置全局表名禁用复数，gorm会自动扩展表名为articles（结构体+s）
	return "tb_user_info"
}
