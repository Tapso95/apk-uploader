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

func (c Application) Register() revel.Result {
	return c.Render()
}
func (c Application) Login() revel.Result{
	return c.Render()
}
func (c Application) getUser(email string) (utilisateur *models.Utilisateur){
	 utilisateur = &models.Utilisateur{}
	fmt.Println("get user",email)
	user:=DB.Where("email_utilisateur=?",email).Find(utilisateur)
	fmt.Println(user)
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
	utilisateur := c.getUser(email)
	if utilisateur != nil {
		bcryptPassword, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
		if bcryptPassword != nil {
			
		fmt.Println( string(bcryptPassword))
		}else{
			fmt.Printf("bcryptPassword")
		}
		utilisateur.PasswordUtilisateur ,_ = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		err := bcrypt.CompareHashAndPassword(utilisateur.PasswordUtilisateur, []byte(password))
		if err != nil {
			// username
			c.Session["utilisateur"] = email
			if remember {
				c.Session.SetDefaultExpiration()
			} else {
				c.Session.SetNoExpiration()
			}
			c.Flash.Success("Welcome, " + email)
			return c.Redirect(routes.Application.Index())
		}
	}
	// fmt.Println("get ",utilisateur.EmailUtilisateur,"+",utilisateur.PasswordUtilisateur,"+",utilisateur.NomUtilisateur,"+",utilisateur.PrenomUtilisateur)
	c.Flash.Out["email"] = email
	c.Flash.Error("Login failed")
	return c.Redirect(routes.Application.Login())
	// return c.Render(utilisateur)
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