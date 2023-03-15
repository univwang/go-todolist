package model

//将定义的模型映射到数据库

func migration() {
	DB.Set(`gorm:table_options`, `charset=utf8mb4`).
		AutoMigrate(&User{})
}
