package controllers

import (
	"apk-uploader/app/models"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
)

var DB *gorm.DB

func InitDB() {
	dbInfo, _ := revel.Config.String("db.info")
	db, err := gorm.Open("mysql", dbInfo)
	if err != nil {

		log.Panicf("Failed gorm.Open: %v\n", err)
	} else {
		fmt.Println("Connected")
	}

	db.DB()
	db.AutoMigrate(&models.Application{})
	db.AutoMigrate(&models.Utilisateur{})
	db.AutoMigrate(&models.Admin{})
	db.AutoMigrate(&models.Evaluation{})
	db.AutoMigrate(&models.Categorie{})
	db.AutoMigrate(&models.Profil{})
	db.AutoMigrate(&models.TypeApplication{})
	db.AutoMigrate(&models.DetailTelechargement{})
	// Dbm := rgorp.Db.Map
	//     setColumnSizes := func(t *gorp.TableMap, colSizes map[string]int) {
	//         for col, size := range colSizes {
	//             t.ColMap(col).MaxSize = size
	//         }
	//     }

	//     t := Dbm.AddTable(models.User{}).SetKeys(true, "UserId")
	//     t.ColMap("Password").Transient = true
	//     setColumnSizes(t, map[string]int{
	//         "Username": 20,
	//         "Name":     100,
	//     })

	//     t = Dbm.AddTable(models.Hotel{}).SetKeys(true, "HotelId")
	//     setColumnSizes(t, map[string]int{
	//         "Name":    50,
	//         "Address": 100,
	//         "City":    40,
	//         "State":   6,
	//         "Zip":     6,
	//         "Country": 40,
	//     })

	//     t = Dbm.AddTable(models.Booking{}).SetKeys(true, "BookingId")
	//     t.ColMap("User").Transient = true
	//     t.ColMap("Hotel").Transient = true
	//     t.ColMap("CheckInDate").Transient = true
	//     t.ColMap("CheckOutDate").Transient = true
	//     setColumnSizes(t, map[string]int{
	//         "NameOnCard": 50,
	//     })
	DB = db
}
