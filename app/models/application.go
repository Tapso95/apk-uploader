package models

import "time"

type Application struct {
	IdApplication         		int `gorm:"primary_key";"AUTO_INCREMENT"`
	UtilisateurId				int
	AdminId						int
	TypeApplicationId			int
	NomApplication   			string
	CodeApplication 			string
	EmplacementApplication      string
	ImageApplication 			string
	DateCreationApplication     time.Time
	DescriptionApplication		string
	TailleApplication 			int
	VersionApplication			int
	StatutApplication			int
}
