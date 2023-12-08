package entity


import (

"gorm.io/driver/sqlite"

"gorm.io/gorm"

)


var db *gorm.DB


func DB() *gorm.DB {

return db

}


func SetupDatabase() {

database, err := gorm.Open(sqlite.Open("reviewsystem.db"), &gorm.Config{})

if err != nil {

panic("failed to connect database")

}



// Migrate the schema
database.AutoMigrate(&User{}, &Movie{}, &Rating{}, &Genre{}, &Review{})


db = database

user1 := User{
	Username: "noteja",
}
db.Model(&User{}).Create(&user1)

movie1 := Movie{
	Title: "Opphenheimer",
	Duration: "180min",
	Description: "เรื่องราวของ Oppenheimer ชายผู้มีปัญหาในตัวเองมากมาย แต่ก็ถูกมองข้ามไปด้วยความปราดเปรื่องของตัวเขา เมื่อเขาถูกขอความช่วยเหลือให้หาหนทางยุติสงครามโลกครั้งที่สอง เขาก็ชี้ไปที่ความหวังเดียวเท่านั้น คือ อาวุธปรมาณูที่มีพลังทำลายล้างรุนแรงจนสามารถยับยั้งไม่ให้ทุกฝ่ายต่อสู้กันต่อไปได้อีก",
	Director: "Christopher Nolan",
	Cast: "Cillian Murphy, Emily Blunt, Matt Damon",
	Poster: "https://i.imgur.com/VIQdyDD.jpg",
}
db.Model(&Movie{}).Create(&movie1)

rating1 := Rating{
	RatingValue: 1,
}
db.Model(&Rating{}).Create(&rating1)

rating2 := Rating{
	RatingValue: 2,
}
db.Model(&Rating{}).Create(&rating2)

rating3 := Rating{
	RatingValue: 3,
}
db.Model(&Rating{}).Create(&rating3)

rating4 := Rating{
	RatingValue: 4,
}
db.Model(&Rating{}).Create(&rating4)

rating5 := Rating{
	RatingValue: 5,
}
db.Model(&Rating{}).Create(&rating5)

genre1 := Genre{
	Name:"แง่บวก",

}
db.Model(&Genre{}).Create(&genre1)

genre2 := Genre{
	Name:"แง่ลบ",

}
db.Model(&Genre{}).Create(&genre2)
}



