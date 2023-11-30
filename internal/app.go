package internal

import (
	"fmt"
	"log"
	"sync"

	"github.com/joydebnath/bebolt/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var db *gorm.DB
var app *App
var dbOnce sync.Once
var appOnce sync.Once

type App struct {
	DB     *gorm.DB
	Logger *zap.SugaredLogger
	Env    *config.Env
}

func NewApp() *App {
	env := config.NewEnv()
	appOnce.Do(func() {
		app = &App{
			DB: getDatabaseConnection(
				env.DBUsername,
				env.DBPassword,
				env.DBHost,
				env.DBPort,
				env.DBDatabase,
			),
			Logger: config.NewLogger(),
			Env:    env,
		}
	})
	return app
}

func getDatabaseConnection(username, password, host, port, database string) *gorm.DB {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, database)
	log.Println("connecting to: " + dns)

	// dbOnce.Do(func() {
	// 	connection, err := gorm.Open(mysql.New(mysql.Config{
	// 		DSN:                       dns,   // data source name
	// 		DefaultStringSize:         256,   // default size for string fields
	// 		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
	// 		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
	// 		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
	// 		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	// 	}), &gorm.Config{})
	// 	if err != nil {
	// 		panic("Unable to set db connection.")
	// 	}
	// 	db = connection
	// })

	return db
}
