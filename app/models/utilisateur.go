package models
import (
	"fmt"
	"github.com/revel/revel"
	"regexp"
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
var utilisateurRegex = regexp.MustCompile("^\\w*$")

func (utilisateur *Utilisateur) Validate(v *revel.Validation) {
	v.Check(utilisateur.NomUtilisateur,
		revel.Required{},
		revel.MaxSize{40},
		revel.MinSize{2},
		revel.Match{utilisateurRegex},
	)

	ValidatePassword(v, utilisateur.PasswordUtilisateur).
		Key("utilisateur.PasswordUtilisateur")

	v.Check(utilisateur.PrenomUtilisateur,
		revel.Required{},
		revel.MaxSize{50},
	)
	v.Check(utilisateur.EmailUtilisateur,
		revel.Required{},
		revel.MaxSize{60},
	)
}

func ValidatePassword(v *revel.Validation, password []byte) *revel.ValidationResult {
	return v.Check(password,
		revel.Required{},
		revel.MaxSize{25},
		revel.MinSize{6},
	)
}