package repository

import (
	interfaceRepo "github.com/url-shortner/pkg/repository/interface"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) interfaceRepo.Repository {
	return &Repository{DB: db}
}

func (r *Repository) Save(url string,short string)error{
	if err:=r.DB.Exec(`INSERT INTO urls(original,shorten) VALUES(?,?)`,url,short).Error;err!=nil{
		return err
	}
	return nil
} 

func (r *Repository) IsShortExist(short string)(bool,error){
	var count int
	if err:=r.DB.Raw(`SELECT COUNT(*) FROM urls WHERE shorten=?`,short).Scan(&count).Error;err!=nil{
		return false,err
	}
	return count>0,nil
}

func (r *Repository) GetURL(short string)(string,error){
	var original string
	if err:=r.DB.Raw(`SELECT original FROM urls WHERE shorten=?`,short).Scan(&original).Error;err!=nil{
		return "",err
	}
	return original,nil
}

func (r *Repository) UpdateCount(short string)error{
	if err:=r.DB.Exec(`UPDATE urls SET counts=counts+1 WHERE shorten=?`,short).Error;err!=nil{
		return err
	}
	return nil
}

func (r *Repository) GetCount(short string)(int,error){
	var counts int
	if err:=r.DB.Raw(`SELECT counts FROM urls WHERE shorten=?`,short).Scan(&counts).Error;err!=nil{
		return 0,err
	}
	return counts,nil
}