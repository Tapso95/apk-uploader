package controllers

import (
	"github.com/revel/revel"
	"apk-uploader/app/routes"
	"apk-uploader/app/models"
	"fmt"
	"os"
	"path"
	"time"
	"io"
	"log"
	"reflect"
	// "mime"
	"mime/multipart"
	"strings"
)

type App struct {
	Application
}

func (c App) ListApp() revel.Result {
	user:= c.connected()
	fmt.Println("path:",path.Base(""))
	path:="./uploads"
	if et, err := os.Stat(path); os.IsNotExist(err) {
    os.Mkdir(path, 755)
	fmt.Println("erreur:",et)
	}else{
		fmt.Println("use")
	}
	var applications []*models.Application
	if user != nil {
		id := user.IdUtilisateur
		apps:= DB.Where("utilisateur_id=?",id).Find(&applications)
		// apps := DB.Select("nom_application,nom_type_app,nom_categorie,statut_application,date_creation_application,code_application,version_application,description_application").Joins("JOIN type_application ON type_application.id_type_app=applications.type_app_id").Joins("JOIN categories ON categorie.id_categorie=type_applications.type_app_id").Joins("JOIN utilisateurs ON utilisateurs.utilisateur_id=applications.utilisateur_id AND utilisateurs.utilisateur_id=?", id)
		if apps != nil {
			
		fmt.Println("User:",id)
		fmt.Println("Bon:",applications)
		return c.Render(user,id , applications)
		}
		return c.Render(user)
	}
	return c.Redirect(routes.Application.Login())
}
func (c App) ViewApp() revel.Result {
	idApp := c.Params.Route.Get("id")
	app := models.Application{}
	var apps1 []*models.TypeApplication
	fmt.Println("app:",app)
	// var app []*models.Application
	fmt.Println("1 call")
	cat:=c.getCategorieApp()
	fmt.Println("cat",cat)
	if user := c.connected(); user != nil {
		fmt.Println("2 call")
		erreur:=DB.Select("type_applications.id_type_app,categories.id_categorie,type_applications.categorie_id,type_applications.nom_type_app").Find(&apps1).Joins("JOIN categories ON categories.id_categorie=type_applications.categorie_id")
		err := DB.Joins("JOIN type_applications ON type_applications.id_type_app=applications.type_application_id").Where("applications.id_application=?",idApp).Find(&app)
		// err := DB.Table("applications").Joins("JOIN type_applications ON type_applications.id_type_app=applications.type_application_id").Where("applications.id_application=?",idApp).Scan(&app)
		fmt.Println("app1",apps1)
		for _, app1 := range apps1 {
			fmt.Println("-+",app1)
		}
		if err!= nil {
			// text:="Bonjour tout le monde"
			// name:="file.text"
			// localfile, err := os.Create(path.Join("http://localhost:9000/public/uploads/", name))
			// if err != nil {
			// 	panic(err)
			// }
			// _, err = localfile.Write([]byte (text))
			// if err != nil {
			// 	panic(err)
			// }
			fmt.Println("id:",idApp,erreur)
			fmt.Println("app:",app)
			return c.Render(app,idApp)
		}
		c.Log.Fatal("Erreur lors du chargement des information de l'application", "error", err)
	}
	return c.Redirect(routes.Application.Login())
}
func (c App) NewApp() revel.Result {
	user:=c.connected()
	var categories []*models.Categorie
	if user != nil {
		id := user.IdUtilisateur
		err := DB.Find(&categories)
		// err := DB.Joins("JOIN type_applications ON type_applications.categorie_id = categories.id_categorie").Find(&categories)
		if err != nil {
			
		fmt.Println("User:",id)
		for _, categorie := range categories {
		
		fmt.Println("cat_typ:",categorie.NomCategorieApp)
		}
		return c.Render(user,id , categories)
		}
		return c.Redirect(routes.App.ListApp())
	}
	return c.Redirect(routes.Application.Login())
}

func (c App) SaveApp(application *models.Application) revel.Result {
	application.DateCreationApplication = time.Now()
	application.UtilisateurId = c.connected().IdUtilisateur
	appFile:=c.Params.Files["application"]
	application.TailleApplication = appFile[0].Size
	fmt.Println("type :",reflect.TypeOf(appFile))
	fmt.Println("user",reflect.TypeOf(appFile[0]))
	app_detail,erreur := c.saveFileToDisk(application.NomApplication,appFile)
	if erreur!=nil {
		fmt.Println(erreur)
	}
	application.EmplacementApplication= app_detail[0]

	imgFile := c.Params.Files["image"]
	image_detail, erreur :=c.saveFileToDisk(application.NomApplication, imgFile)
	if erreur!= nil {
		fmt.Println(erreur)
	}
	application.ImageApplication=image_detail[0]

	// filename := c.Params.Files["file"][0].Filename
	// contenttype := mime.TypeByExtension(path.Ext(filename))
	// if contenttype == "" {
	// 	// Try to figure out the content type from the data
	// 	contenttype = http.DetectContentType(file)
	// }
	// filename, err := SaveFileToDb(c, contenttype, filename)
	// if err != nil {
	// 	c.Flash.Error(err.Error())
	// 	return c.Redirect(App.Index)
	// }
	// if err := WriteFileToDisk(filename, file); err != nil {
	// 	c.Flash.Error(err.Error())
	// 	return c.Redirect(App.Index)
	// }
	// c.Flash.Success("Success!")
	// c.Flash.Out["FileName"] = filename


	err:= DB.Create(&application)
	if err == nil {
		panic(err)
	}
	return c.Redirect(routes.App.ListApp())
}

 	//The function will took types string and []*multipart.FileHeader as parameter
func (c App) saveFileToDisk(filename string,files []*multipart.FileHeader) ([5]string, error) {
	var filedata [5]string
	var chemin string
	for i,_ := range files{
		//Verification de l'existance d'un fichier
		file,err := files[i].Open()
		fmt.Println("header",files[i].Filename)
		fmt.Println("params:",files)
		fmt.Println("cool",files[i].Header.Get("Content-Type"))
		fmt.Println("length %6.2f",files[i].Size)

		defer file.Close()
		if err != nil {
			log.Print(err)
		}
		//Verification du type de mime
		contentType:=files[i].Header.Get("Content-Type")
		switch (contentType) {
			case "image/jpeg":
				fmt.Println("image jpeg")
				chemin="public/uploads/images/"
				break
			case "image/png":
				fmt.Println("image png")
				chemin="public/uploads/images/"
				break
			case "application/vnd.android.package-archive" :
				fmt.Println("apk")
				chemin="public/uploads/applications/"
		}
		filerename := strings.Replace(filename," ","-",-1)+path.Ext(files[i].Filename)
		fmt.Println("filename:",path.Ext(files[i].Filename))
		fmt.Println("filerename:",filerename)
		// path:=
		//Creation de la destination du fichier, droit d'ecriture obligatoire
		dst_path := chemin + strings.ToLower(filerename)
		dst,err := os.Create(dst_path)
		defer dst.Close()
		fmt.Println("dst_path",dst_path)
		filedata[0]=dst_path
		filedata[1]=contentType
		// filedata[2]=strings.toString(files[i].Size)
		//Limitation des acces
		defer os.Chmod(dst_path,(os.FileMode)(0644))
		if err != nil {
			return filedata, err
		}
		//Copie du fichier uploader vers la destination
		if _, err := io.Copy(dst, file); err != nil {
			return filedata, err
		}
	}
	return filedata, nil
}

// func (c App) WriteFileToDisk(filename string, file []byte) error {
// 	localfile, err := os.Create(path.Join("uploads", filename))
// 	if err != nil {
// 		return err
// 	}
// 	_, err = localfile.Write(file)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

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

func (c App) DeleteApp(id int) bool {
	return true
}


func (c App) SaveApp1(app *models.Application) (bool) {
	err := DB.Create(&app)
	if err == nil {
		go revel.WARN.Printf("Impossible de sauvegarder le compte: %v erreur %v", err)
		return false
	}
	return true
}

