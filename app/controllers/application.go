package controllers

import (
	"github.com/revel/revel"
	// "apk-uploader/app/routes"
	"apk-uploader/app/models"
	// "fmt"
)

type App struct {
	Application
}

func (c App) GetApplicationListById() revel.Result{

	return c.Render()
}
func (c App) SaveApplication(application *models.Application) (bool){
	err := DB.Create(&application)
	if err == nil {
		// go revel.WARN.Printf("Impossible de sauvegarder le compte: %v erreur %v", utilisateur.EmailUtilisateur, err)
		return false
	}
	return true
}

func (c App) AddApp(id int) {
	
}

func (c App) GetApplication(id int) {
	
}