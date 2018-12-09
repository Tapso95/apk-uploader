package models

type Application struct {
	IdApplication         		int
	UtilisateurId				int
	AdminId						int
	TypeApplicationId			int
	NomApplication   			string
	CodeApplication 			string
	EmplacementApplication      string
	ImageApplication 			string
	DateCreationApplication     string
	DescriptionApplication		string
	VersionApplication			int
	StatutApplication			int
}
