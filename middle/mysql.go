package middle

import "github.com/jinzhu/gorm"

func InitMysql() {
	db, err := gorm.Open("mysql", "root:*****@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
