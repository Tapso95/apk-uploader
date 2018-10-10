package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Hello(username string, password string) revel.Result{
	c.Validation.Required(username).Message("Le mail est obligatoire!")
    c.Validation.Required(password).Message("Le mot de passe est obligatoire!")
    c.Validation.MinSize(password, 6).Message("Le mot de passe doit contenir au moins 6 caractères!")
    if c.Validation.HasErrors() {
    	c.Validation.Keep()
    	c.FlashParams()
    }
	return c.Render(username)
}