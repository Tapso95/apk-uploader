package controllers

import (
	"github.com/revel/revel"
	// "apk-uploader/app/routes"
	"apk-uploader/app/models"
	// "fmt"
	"os"
	"path"
)

type App struct {
	Application
}

func (c App) WriteFileToDisk(filename string, file []byte) error {
	localfile, err := os.Create(path.Join("uploads", filename))
	if err != nil {
		return err
	}
	_, err = localfile.Write(file)
	if err != nil {
		return err
	}
	return nil
}

// func SaveFileToDb(c File, contenttype string, filename string) (string, error) {
// 	var existing db.File
// 	if err := conn.Where("file_name = ?", filename).First(&existing).Error; err == nil {
// 		ext := path.Ext(filename)
// 		newfile := strings.TrimSuffix(filename, ext) + "_" + ext
// 		if len(filename) > 64 {
// 			return "", errors.New("Too many files with same name") // I'm lazy
// 		}
// 		return SaveFileToDb(c, contenttype, newfile)
// 	}
// 	file := db.File {
// 		FileName: filename,
// 		ContentType: contenttype,
// 	}
// 	user := c.CurrentUser()
// 	if user == nil {
// 		if err := conn.Create(&file).Error; err != nil {
// 			return "", err
// 		}
// 	} else {
// 		if err := conn.Model(&user).Association("Files").Append(&file).Error; err != nil {
// 			return "", err
// 		}
// 		return filename, nil
// 	}
// 	return filename, nil
// }

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