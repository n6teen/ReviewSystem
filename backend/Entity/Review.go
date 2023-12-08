package entity

import(
	"time"
	"gorm.io/gorm"
)

type User struct{
	gorm.Model
	Username string

	Review	[]Review	`gorm:foreignkey:UserID`
}

type Movie struct{
	gorm.Model
	Title string
	Duration string
	Description string
	Director string
	Cast string
	Poster string

	Review	[]Review	`gorm:foreignkey:MovieID`
}

type Rating struct{
	gorm.Model
	RatingValue int

	Review	[]Review	`gorm:foreignkey:RatingID`
}

type Genre struct{
	gorm.Model
	Name string

	Review	[]Review	`gorm:foreignkey:GenreID`
}

type Review struct{
	gorm.Model
	ReviewText string
	DateTime time.Time

	UserID *uint
	User User `gorm:"references:id"`

	MovieID *uint
	Movie Movie `gorm:"references:id"`

	RatingID *uint
	Rating Rating `gorm:"references:id"`

	GenreID *uint
 	Genre Genre `gorm:"references:id"`

}

