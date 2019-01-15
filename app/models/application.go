package models

import (
	"time"
	// "github.com/jinzhu/gorm"
)

type Application struct {
	IdApplication         		int `gorm:"primary_key";"AUTO_INCREMENT"`
	UtilisateurId				int
	AdminId						int
	TypeApplicationId			int
	NomApplication   			string
	EmplacementApplication      string
	ImageApplication 			string
	DateCreationApplication     time.Time
	DescriptionApplication		string
	TailleApplication 			int64
	VersionApplication			string
	StatutApplication			int
	// Admin 						Admin `gorm:"foreignkey:AdminId;association_foreignkey:IdAdmin"`
	TypeApplication 			TypeApplication `gorm:"foreignkey:TypeApplicationId;association_foreignkey:IdTypeApplication"`
}
