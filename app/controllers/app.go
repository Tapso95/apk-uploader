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
		// fmt.Println("user0 %s",email)
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
	user:=DB.Where("email_utilisateur=?",email).Find(utilisateur)
	if user == nil {
		fmt.Printf("user not found")
		c.Log.Error("Failed to find user")
		panic(user)
	}
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
	fmt.Println("users:",utilisateur)
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

func (c Application) getCategorieApp() int{
	var categories []*models.Categorie
	cat:=DB.Find(&categories).Joins("JOIN type_applications ON type_applications.categorie_id=categories.id_categorie")
	if cat == nil {
		panic(cat)
	}
	for _, categorie := range categories {
		fmt.Println("-+",cat)
		fmt.Println("-+",categorie)
		}
	return 0
}

func (c Application) LoadTypeApp() revel.Result{
	idCat := c.Params.Route.Get("id")
	fmt.Println("id:",idCat)
	var typeApps []*models.TypeApplication
	if user := c.connected(); user != nil {
		err := DB.Where("categorie_id=?",idCat).Find(&typeApps)
		if err!= nil {
			// erreur := true;
			fmt.Println("-+",typeApps)
			// return c.RenderJSON(erreur, typeApps,idCat)
			return c.RenderJSON(typeApps)
		}
		erreur := false
		c.Log.Fatal("Erreur lors du chargement des types d'application de l'application", "error", err)
		return c.RenderJSON(erreur)
	}
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