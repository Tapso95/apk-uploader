package controllers

import (
	"golang.org/x/crypto/bcrypt"
	"github.com/revel/revel"
	"apk-uploader/app/routes"
	"apk-uploader/app/models"
	"fmt"
)

type Application struct {
	*revel.Controller
}

func (c Application) Index() revel.Result {
	return c.Render()
}
func (c Application) connected() *models.Utilisateur {
	if c.ViewArgs["utilisateur"] != nil {
		return c.ViewArgs["utilisateur"].(*models.Utilisateur)
	}
	if email, ok := c.Session["utilisateur"]; ok {
		return c.getUser(email)
	}
	return nil
}

func (c Application) Register() revel.Result {
	return c.Render()
}
func (c Application) Login() revel.Result{
	return c.Render()
}
func (c Application) getUser(email string) (utilisateur *models.Utilisateur){
	 utilisateur = &models.Utilisateur{}
	// fmt.Println("get user",email)
	user:=DB.Where("email_utilisateur=?",email).Find(utilisateur)
	// fmt.Println(user)
	// err:= DB.Select(utilisateur,`SELECT * FROM utilisateurs WHERE email_utilisateur=?`,email)
	if user == nil {
		fmt.Printf("user not found")
		c.Log.Error("Failed to find user")
		panic(user)
	}
	fmt.Printf("user %s",utilisateur)
	return 
}

func (c Application) SaveUser(utilisateur *models.Utilisateur, password models.Password) revel.Result {
		fmt.Println(utilisateur)
	if exists :=c.getUser(utilisateur.EmailUtilisateur); exists.EmailUtilisateur == utilisateur.EmailUtilisateur {
		msg := fmt.Sprint("Cet utilisateur existe déjà, veuillez chager l'adresse email ou conatcter l'administrateur")
		fmt.Println("Cet utilisateur existe déjà, veuillez chager l'adresse email ou conatcter l'administrateur")
		c.Validation.Required(utilisateur.EmailUtilisateur != exists.EmailUtilisateur).
			Message(msg)
	}
	fmt.Println("suite")
	// utilisateur.Validate(c.Validation)
	// *models.Utilisateur.ValidatePassword(c.Validation, password)
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		c.Flash.Error("veuillez corriger l'erreur signalée")
		return c.Redirect(routes.Application.Register())
	}
	err:=c.Save(utilisateur,password)
	if err{
		fmt.Println("User saved")
	}else{
		fmt.Println("Something wrong")
	}
	c.Session["utilisateur"] = utilisateur.EmailUtilisateur
	c.Flash.Success("Welcome, " + utilisateur.NomUtilisateur)
	return c.Redirect(routes.Application.Index())
}

func (c Application) PostLogin(email, password string, remember bool) revel.Result {
	fmt.Println(email)
	fmt.Println(password)
	fmt.Println(remember)
	utilisateur := c.getUser(email)
	if utilisateur != nil {
		err := bcrypt.CompareHashAndPassword(utilisateur.PasswordUtilisateur,[]byte(password))
		if err == nil {
			fmt.Println("cool")
			c.Session["utilisateur"] = email
			if remember {
				c.Session.SetDefaultExpiration()
			} else {
				c.Session.SetNoExpiration()
			}
			c.Flash.Success("Welcome, " + email)
			return c.Redirect(routes.Application.Index())
		}else{
			fmt.Println("desolé")
		}
	}
	c.Flash.Out["email"] = email
	c.Flash.Error("Echec de connexion")
	return c.Redirect(routes.Application.Login())
}

func (c Application) Save(utilisateur *models.Utilisateur, p models.Password) (bool) {
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

func (c Application) Hello(username string, password string) revel.Result{
	c.Validation.Required(username).Message("Le mail est obligatoire!")
    c.Validation.Required(password).Message("Le mot de passe est obligatoire!")
    c.Validation.MinSize(password, 6).Message("Le mot de passe doit contenir au moins 6 caractères!")
    if c.Validation.HasErrors() {
    	c.Validation.Keep()
    	c.FlashParams()
    }
	return c.Render(username)
}