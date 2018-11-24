package controllers

import (
	"golang.org/x/crypto/bcrypt"
	"github.com/revel/revel"
	"apk-uploader/app/routes"
	"apk-uploader/app/models"
	"fmt"
)

type User struct {
	Application
}

// func Index() {
	
// }



func (u User) DeleteApp(id int) (bool){
	return true
}
func (u User) NewApp() revel.Result{
	u.Render()
}
func (u User) SaveApp(app *models.Application) (bool){
	err:=u.Save(app)
	if err{
		fmt.Println("Application saved")
	}else{
		fmt.Println("Something wrong")
	}
}

func (u User) ViewUser(id int) revel.Result{
	user:=u.connected()
	if user != nil {
		fmt.Println("user %s",user.NomUtilisateur)
	}
	return u.Render(user)
}

func (u User) Register() revel.Result {
	return u.Render()
}

func (u User) getUser(email string) (utilisateur *models.Utilisateur){
	 utilisateur = &models.Utilisateur{}
	// fmt.Println("get user",email)
	user:=DB.Where("email_utilisateur=?",email).Find(utilisateur)
	// fmt.Println(user)
	// err:= DB.Select(utilisateur,`SELECT * FROM utilisateurs WHERE email_utilisateur=?`,email)
	if user == nil {
		fmt.Printf("user not found")
		u.Log.Error("Failed to find user")
		panic(user)
	}
	fmt.Printf("user %s",utilisateur)
	return 
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
	if exists :=u.getUser(utilisateur.EmailUtilisateur); exists.EmailUtilisateur == utilisateur.EmailUtilisateur {
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
	err:=u.Save(utilisateur,password)
	if err{
		fmt.Println("User saved")
	}else{
		fmt.Println("Something wrong")
	}
	u.Session["utilisateur"] = utilisateur.EmailUtilisateur
	u.Flash.Success("Welcome, " + utilisateur.NomUtilisateur)
	return u.Redirect(routes.Application.Index())
}

func (u User) Save(utilisateur *models.Utilisateur, p models.Password) (bool) {
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

func (u User) Setting() revel.Result{
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