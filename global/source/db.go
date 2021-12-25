package source

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB
var err error

func init() {
	dsn := "blog:Tenderness0912!@tcp(localhost:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
	DB, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名，启用该选项，此时，`Article` 的表名应该是 `it_article`
		},
	})
	if DB == nil {
		fmt.Println("数据库连接异常！")
	}
	fmt.Println("创建连接")
}
