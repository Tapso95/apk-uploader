package controllers

import (
	"apk-uploader/app/models"
	"apk-uploader/app/routes"
	"fmt"
	// "time"
	// "os"
	// "io"
	// "log"
	// "reflect"
	// "path"
	// "io/ioutil"
	// "mime"
	// "mime/multipart"
	// "strings"
	"github.com/revel/revel"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Application
}

// func Index() {

// }


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

func (u User) preloadTypeCat() (categorie *models.Categorie)  {
	categorie = &models.Categorie{}
	type_cat := DB.Find(&categorie).Joins("JOIN type_applications ON type_applications.categorie_id = categories.id_categorie")
	// type_cat := DB.Joins("JOIN emails ON emails.user_id = users.id AND emails.email = ?", "jinzhu@example.org").Joins("JOIN credit_cards ON credit_cards.user_id = users.id").Where("credit_cards.number = ?", "411111111111").Find(&user)
	fmt.Println("256",type_cat)
	if type_cat == nil {
		panic(type_cat)
	}
	return
}

func (u User) getAllUser() revel.Result{
	user:=u.connected()
	var utilisateurs []*models.Utilisateur
	if user != nil {
		// if(user.ProfilId){

		// }
		id := user.IdUtilisateur
		err := DB.Where("").Find(&utilisateurs)
		// err := DB.Joins("JOIN type_applications ON type_applications.categorie_id = categories.id_categorie").Find(&categories)
		if err != nil {
			
		return u.Render(user,id , utilisateurs)
		}
		return u.Redirect(routes.App.ListApp())
	}
	return u.Redirect(routes.Application.Login())
	
}
// func (u User) addComment(comment *models.EvaluationApplication, idUser int, idApplication int) (bool) {
// 	err := DB.Create(&comment)
// 	if err == nil {
// 		// go revel.WARN.Printf("Impossible de sauvegarder le compte: %v erreur %v", utilisateur.EmailUtilisateur, err)
// 		return false
// 	}
// 	return true
// }
