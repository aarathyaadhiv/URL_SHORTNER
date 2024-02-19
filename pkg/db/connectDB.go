package db

import (
	"fmt"

	"github.com/url-shortner/pkg/config"
	"github.com/url-shortner/pkg/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)




func ConnectDB(c config.Config)(*gorm.DB,error){
	psql:=fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s",c.DB_HOST,c.DB_USER,c.DB_NAME,c.DB_PORT,c.DB_PASSWORD)
	db,err:=gorm.Open(postgres.Open(psql),&gorm.Config{SkipDefaultTransaction: true})

	db.AutoMigrate(&domain.Url{})
	return db,err
}