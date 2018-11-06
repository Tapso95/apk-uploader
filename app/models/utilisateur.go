package models
import (
	// "github.com/revel/revel"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	// "regexp"
)


type Utilisateur struct {
	IdUtilisateur   		int
	NomUtilisateur			string
	PrenomUtilisateur 		string
	EmailUtilisateur       	string
	PasswordUtilisateur    	[]byte
	StatutUtilisateur 		int
}

func (u *Utilisateur) String() string {
	return fmt.Sprintf("Utilisateur(%s %s)", u.NomUtilisateur,u.PrenomUtilisateur)
}
type Password struct {
	Pass        string
	PassConfirm string
}

// var utilisateurRegex = regexp.MustCompile("^\\w*$")

// func (utilisateur *Utilisateur) Validate(v *revel.Validation) {
// 	v.Check(utilisateur.NomUtilisateur,
// 		revel.Required{},
// 		revel.MaxSize{40},
// 		revel.MinSize{2},
// 	)
// 	v.Check(utilisateur.PrenomUtilisateur,
// 		revel.Required{},
// 		revel.MaxSize{50},
// 	)
// 	v.Check(utilisateur.EmailUtilisateur,
// 		revel.Required{},
// 	)
// 	v.Email(utilisateur.EmailUtilisateur)
// }

// func (utilisateur *Utilisateur) ValidatePassword(v *revel.Validation, password Password){
// 	v.Check(password.Pass,
// 		revel.MinSize{8},
// 	)
// 	v.Check(password.PassConfirm,
// 		revel.MinSize{8},
// 	)
// 	v.Required(password.Pass == password.PassConfirm).Message("Les mot de passe sont différents.")
// }
