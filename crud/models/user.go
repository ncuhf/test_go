// user
package user

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	Id       int    //`gorm:"column:user_id";"AUTO_INCREMENT"`
	Name     string `gorm:"column:user_name"`
	Number   string `gorm:"column:user_number";UNIQUE`
	PassWord string `gorm:"xolumn:user_password"`
}

func Connection() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@/testgorm?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("连接数据库失败")
	}
	//defer db.Close()

	db.AutoMigrate(&User{})

	return db
}
