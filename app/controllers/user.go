package controllers

import (
	"apk-uploader/app/models"
	"apk-uploader/app/routes"
	"fmt"

	"github.com/revel/revel"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Application
}

// func Index() {

// }

func (u User) DeleteApp(id int) bool {
	return true
}
func (u User) ListApp() revel.Result {
	user:= u.connected()
	var applications []*models.Application
	if user != nil {
		id := user.IdUtilisateur
		apps:= DB.Where("utilisateur_id=?",id).Find(&applications)
		// apps := DB.Select("nom_application,nom_type_app,nom_categorie,statut_application,date_creation_application,code_application,version_application,description_application").Joins("JOIN type_application ON type_application.id_type_app=applications.type_app_id").Joins("JOIN categories ON categorie.id_categorie=type_applications.type_app_id").Joins("JOIN utilisateurs ON utilisateurs.utilisateur_id=applications.utilisateur_id AND utilisateurs.utilisateur_id=?", id)
		if apps != nil {
			
		fmt.Println("User:",id)
		fmt.Println("Bon:",applications)
		return u.Render(user,id , applications)
		}
		return u.Render(user)
	}
	return u.Redirect(routes.Application.Login())
}
func (u User) ViewApp() revel.Result {
	idApp := u.Params.Route.Get("id")
	var app *models.Application
	fmt.Println("1 call")
	if user := u.connected(); user != nil {
		fmt.Println("2 call")
		err := DB.Where("id_application=?",idApp).Find(&app)
		fmt.Println("inconnu")
		if err!= nil {
			fmt.Println("id:",idApp)
			return u.Render(app,idApp)
		}
		u.Log.Fatal("Erreur lors du chargement des information de l'application", "error", err)
	}
	return u.Redirect(routes.Application.Login())
}
func (u User) NewApp() revel.Result {
	user:=u.connected()
	if user != nil {
		categories := u.getCategorieApp()
		typeApplication := u.getTypeApp() 
		fmt.Println("++",categories.IdCategorie)
		// fmt.Println("++" ,user.IdUtilisateur)
		return u.Render(user, categories, typeApplication)
	}
	return u.Redirect(routes.Application.Login())
}

func (u User) SaveApp(app *models.Application) bool {
	err := DB.Create(&app)
	if err == nil {
		go revel.WARN.Printf("Impossible de sauvegarder le compte: %v erreur %v", err)
		return false
	}
	return true
}

func (u User) ViewUser(id int) revel.Result {
	fmt.Println(id)
	// user := u.connected()
	// if user != nil {
	// 	fmt.Println("user %s", user.NomUtilisateur)
	// }
	return u.Render()
}

func (u User) Register() revel.Result {
	return u.Render()
}

func (u User) SaveUser(utilisateur *models.Utilisateur, password models.Password) revel.Result {
	fmt.Println(utilisateur)
	u.Validation.Required(utilisateur.NomUtilisateur).Message("Missing username")
	u.Validation.Required(utilisateur.PasswordUtilisateur).Message("Missing password")
	u.Validation.Required(utilisateur.EmailUtilisateur).Message("Missing email")
	if u.Validation.HasErrors() {
		u.Validation.Keep()
		u.FlashParams()
	}
	if exists := u.getUser(utilisateur.EmailUtilisateur); exists.EmailUtilisateur == utilisateur.EmailUtilisateur {
		msg := fmt.Sprint("Cet utilisateur existe déjà, veuillez chager l'adresse email ou contacter l'administrateur")
		fmt.Println("Cet utilisateur existe déjà, veuillez chager l'adresse email ou contacter l'administrateur")
		u.Validation.Required(utilisateur.EmailUtilisateur != exists.EmailUtilisateur).
			Message(msg)
	}
	fmt.Println("suite")
	// utilisateur.Validate(c.Validation)
	// *models.Utilisateur.ValidatePassword(u.Validation, password)
	// if u.Validation.HasErrors() {
	// 	u.Validation.Keep()
	// 	u.FlashParams()
	// 	u.Flash.Error("veuillez corriger l'erreur signalée")
	// 	return u.Redirect(routes.User.Register())
	// }
	err := u.Save(utilisateur, password)
	if err {
		fmt.Println("User saved")
	} else {
		fmt.Println("Something wrong")
	}
	u.Session["utilisateur"] = utilisateur.EmailUtilisateur
	u.Flash.Success("Welcome, " + utilisateur.NomUtilisateur)
	return u.Redirect(routes.Application.Index())
}

func (u User) Save(utilisateur *models.Utilisateur, p models.Password) bool {
	// Calculate the new password hash or load the existing one so we don't clobber it on save.
	if p.Pass != "" {
		utilisateur.PasswordUtilisateur, _ = bcrypt.GenerateFromPassword([]byte(p.Pass), bcrypt.DefaultCost)
	}
	err := DB.Create(&utilisateur)
	if err == nil {
		go revel.WARN.Printf("Impossible de sauvegarder le compte: %v erreur %v", utilisateur.EmailUtilisateur, err)
		return false
	}
	return true
}

func (u User) Setting() revel.Result {
	return u.Render()
}

func (u User) SaveSetting() {

}

// func (u User) addComment(comment *models.EvaluationApplication, idUser int, idApplication int) (bool) {
// 	err := DB.Create(&comment)
// 	if err == nil {
// 		// go revel.WARN.Printf("Impossible de sauvegarder le compte: %v erreur %v", utilisateur.EmailUtilisateur, err)
// 		return false
// 	}
// 	return true
// }
