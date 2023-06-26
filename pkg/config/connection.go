package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var (
	db *gorm.DB
)

func Connect() {
	dbURL := "postgres://root:x42WtpTNdUWyiI12aDdhtUceIaAp63ru@dpg-cibn5tiip7vnjjnvq5pg-a.frankfurt-postgres.render.com/library_z93u"
	d, err := gorm.Open("postgres", dbURL)
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
