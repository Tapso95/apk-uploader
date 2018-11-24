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
	user:=c.connected()
	if user != nil {
		fmt.Println("user1 %s",user.NomUtilisateur)
	}
	return c.Render(user)
}
func (c Application) Hello(username string) revel.Result{
	// c.Validation.Required(username).Message("Le mail est obligatoire!")
	c.Validation.MaxSize(username,3).Message("Username inferieur a 3")
    // c.Validation.Required(password).Message("Le mot de passe est obligatoire!")
    // c.Validation.MinSize(password, 6).Message("Le mot de passe doit contenir au moins 6 caractères!")
    if c.Validation.HasErrors() {
    	c.Validation.Keep()
    	c.FlashParams()
    	return c.Redirect(Application.Login)
    }
	return c.Render(username)
}

func (c Application) connected() *models.Utilisateur {
	if c.ViewArgs["utilisateur"] != nil {
		return c.ViewArgs["utilisateur"].(*models.Utilisateur)
	}
	if email, ok := c.Session["utilisateur"]; ok {
		fmt.Println("user0 %s",email)
		return c.getUser(email)
	}
	return nil
}


func (c Application) Login() revel.Result{
	return c.Render()
}

func (c Application) Logout() revel.Result {
	delete(c.Session, "utilisateur")
	return c.Redirect(routes.Application.Index())
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



func (c Application) PostLogin(email, password string, remember bool) revel.Result {
	fmt.Println(email)
	fmt.Println(password)
	fmt.Println(remember)
	c.Validation.Required(email).Message("Missing username")
    c.Validation.Required(password).Message("Missing password")
    if c.Validation.HasErrors() {
        c.Validation.Keep()
        c.FlashParams()
    }
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
			c.Flash.Error("Echec de connexion")
		}
	}
	c.Flash.Out["email"] = email
	c.Flash.Error("Echec de connexion")
	return c.Redirect(routes.Application.Login())
}


// func (c Application) Hello(username string, password string) revel.Result{
// 	c.Validation.Required(username).Message("Le mail est obligatoire!")
//     c.Validation.Required(password).Message("Le mot de passe est obligatoire!")
//     c.Validation.MinSize(password, 6).Message("Le mot de passe doit contenir au moins 6 caractères!")
//     if c.Validation.HasErrors() {
//     	c.Validation.Keep()
//     	c.FlashParams()
//     }
// 	return c.Render(username)
// }