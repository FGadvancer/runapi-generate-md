package db

import (
	"github.com/ruapi-generate-md/pkg/db/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type DataBase struct {
	conn  *gorm.DB
	dbDir string
	*model.Page
	*model.Catalog
	*model.Item
	*model.RunapiGlobalParam
}

func NewDataBase(dbDir string) *DataBase {
	return &DataBase{dbDir: dbDir}
}

func (d *DataBase) Init() {
	db, err := gorm.Open(sqlite.Open(d.dbDir), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetConnMaxLifetime(time.Hour * 1)
	sqlDB.SetMaxOpenConns(3)
	sqlDB.SetMaxIdleConns(2)
	sqlDB.SetConnMaxIdleTime(time.Minute * 10)
	d.conn = db
	d.Page = model.NewPage(db)
	d.Item = model.NewItem(db)
	d.Catalog = model.NewCatalog(db)
	d.RunapiGlobalParam = model.NewRunapiGlobalParam(db)
}
