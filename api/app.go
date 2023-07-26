package app

import (
	"fmt"
	"log"

	"github.com/gorilla/mux"
	model "github.com/hashi7412/restfulapi-with-gorm-and-gorillamux/api/app"
	"github.com/jinzhu/gorm"
	"honnef.co/go/tools/config"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (a *App) Initialize(config *config.Config) {
	dbURI := fmt.Sprintf(
		"%s:%s@/%s?charset=%s&parseTime=True",
		config.DB.Username,
		config.DB.Password,
		config.DB.Name,
		config.DB.Charset,
	)

	db, err := gorm.Open(config.DB.Dialect, dbURI)
	if err != nil {
		log.Fatal("Could not connect database")
	}

	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}
