package dao

import (
	"fmt"
	"log"
	"path"

	"durian/dao/sqlite"
	"github.com/jinzhu/gorm"
)

var Object *gorm.DB

// DatabaseOptions holds information about how the database instance should be initialized
type DatabaseOptions struct {
	Connection string
	LogMode    bool
}

// NewDb initializes a new database instance.
func NewDb(options DatabaseOptions, dbPath string) *gorm.DB {
	var err error
	dbDefaultPath := path.Join(dbPath, "metadata.db")
	Object, err = sqlite.NewSQLiteDatabase(dbDefaultPath, true)
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %s\n", err))
	}

	err = migrateSchema(Object)
	if err != nil {
		log.Fatalf("failed to migrate database: %s", err)
	}

	return Object
}

var allModels = []interface{}{
	&Picture{}, &PictureArtist{}, &PictureCharacter{}, &PictureTag{}, &PictureMetaData{},
	&PictureCopyRight{}, &AirTask{},
}

func migrateSchema(db *gorm.DB) error {
	return db.AutoMigrate(allModels...).Error
}
