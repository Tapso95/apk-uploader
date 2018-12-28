package models


type Evaluation struct {
	IdEvaluation 		int `gorm:"primary_key";"AUTO_INCREMENT"`
	UtilisateurId		int
	ApplicationId		int
	evaluation 			int
	commentaire 		string
	date_evaluation		string
}
