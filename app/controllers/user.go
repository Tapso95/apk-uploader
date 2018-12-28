package controllers

import (
	"apk-uploader/app/models"
	"apk-uploader/app/routes"
	"fmt"
	"time"
	"os"
	"io"
	"log"
	"reflect"
	"path"
	"io/ioutil"
	// "mime"
	"mime/multipart"
	"strings"
	"github.com/revel/revel"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Application
}

// func Index() {

// }

func (u User) DeleteApp(id int) bool {
	return true
}
func (u User) ListApp() revel.Result {
	user:= u.connected()
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
		return u.Render(user,id , applications)
		}
		return u.Render(user)
	}
	return u.Redirect(routes.Application.Login())
}
func (u User) ViewApp() revel.Result {
	idApp := u.Params.Route.Get("id")
	app := models.Application{}
	// var application []*models.Application
	fmt.Println("1 call")
	if user := u.connected(); user != nil {
		fmt.Println("2 call")
		err := DB.Where("id_application=?",idApp).Find(&app)
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
			files, err := ioutil.ReadDir(".")
			if err != nil {
				log.Fatal(err)
			}

			for _, file := range files {
				fmt.Println(file.Name())
			}
			dir, err:= os.Getwd()
			fmt.Println("dir",dir)
			fmt.Println("dir:",revel.Filters)
			fmt.Println("id:",idApp)
			fmt.Println("app:",app)
			return u.Render(app,idApp)
		}
		u.Log.Fatal("Erreur lors du chargement des information de l'application", "error", err)
	}
	return u.Redirect(routes.Application.Login())
}
func (u User) NewApp() revel.Result {
	user:=u.connected()
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
		return u.Render(user,id , categories)
		}
		return u.Redirect(routes.User.ListApp())
	}
	return u.Redirect(routes.Application.Login())
}

func (u User) SaveApp1(app *models.Application) (bool) {
	err := DB.Create(&app)
	if err == nil {
		go revel.WARN.Printf("Impossible de sauvegarder le compte: %v erreur %v", err)
		return false
	}
	return true
}

func (u User) SaveApp(application *models.Application) revel.Result {
	application.DateCreationApplication = time.Now()
	application.UtilisateurId = u.connected().IdUtilisateur
	appFile:=u.Params.Files["application.EmplacementApplication"]
	 fmt.Println("type :",reflect.TypeOf(appFile))
	 fmt.Println("user",reflect.TypeOf(appFile[0]))
	 path,err := u.saveFileToDisk(application.NomApplication,appFile)
	 fmt.Println("err:",err,path)
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

 	//The function will took type []*multipart.FileHeader as parameter

	// err := DB.Create(&application)
	// if err == nil {
	// 	panic(err)
	// }
	return u.Redirect(routes.User.ListApp())
}

func (u User) saveFileToDisk(filename string,files []*multipart.FileHeader) (string, error) {
	var chemin string
	for i,_ := range files{
		//Verification de l'existance d'un fichier
		file,err := files[i].Open()
		fmt.Println("header",files[i].Filename)
		fmt.Println("params:",files)
		fmt.Println("cool",files[i].Header.Get("Content-Type"))
		fmt.Println("length",files[i].Size)
		defer file.Close()
		if err != nil {
			log.Print(err)
		}
		//Verification du type de mime
		contentType:=files[i].Header.Get("Content-Type")
		switch (contentType) {
			case "image/jpeg":
				fmt.Println("image jpeg")
				chemin="public/uploads/images"
				break
			case "image/png":
				fmt.Println("image png")
				chemin="public/uploads/images"
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
		//Limitation des acces
		defer os.Chmod(dst_path,(os.FileMode)(0644))
		if err != nil {
			return chemin, err
		}
		//Copie du fichier uploader vers la destination
		if _, err := io.Copy(dst, file); err != nil {
			return chemin, err
		}
	}
	return chemin, nil
}

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

// func (u User) addComment(comment *models.EvaluationApplication, idUser int, idApplication int) (bool) {
// 	err := DB.Create(&comment)
// 	if err == nil {
// 		// go revel.WARN.Printf("Impossible de sauvegarder le compte: %v erreur %v", utilisateur.EmailUtilisateur, err)
// 		return false
// 	}
// 	return true
// }
